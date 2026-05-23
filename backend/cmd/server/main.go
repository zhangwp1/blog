package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"blog/config"
	"blog/internal/model"
	"blog/internal/router"

	"golang.org/x/crypto/bcrypt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func seedAdmin(db *gorm.DB) {
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count == 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		db.Create(&model.User{
			Username: "admin",
			Password: string(hash),
			Nickname: "博主",
		})
		log.Println("Default admin user created: admin / admin123")
	}
}

func main() {
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := gorm.Open(sqlite.Open(config.AppConfig.Database.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	sqlDB.SetMaxIdleConns(config.AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.AppConfig.Database.MaxOpenConns)

	if err := db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Tag{},
		&model.Article{},
		&model.Comment{},
	); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	seedAdmin(db)

	r := router.SetupRouter(db)

	go func() {
		addr := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
		log.Printf("Server starting on %s", addr)
		if err := r.Run(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
