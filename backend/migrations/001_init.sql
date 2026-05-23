CREATE DATABASE IF NOT EXISTS blog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE blog;

-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
    `id`          BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT,
    `username`    VARCHAR(64)      NOT NULL,
    `password`    VARCHAR(255)     NOT NULL COMMENT 'bcrypt hash',
    `nickname`    VARCHAR(64)      NOT NULL DEFAULT '',
    `avatar`      VARCHAR(255)     NOT NULL DEFAULT '',
    `created_at`  DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 分类表
CREATE TABLE IF NOT EXISTS `categories` (
    `id`          BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(64)      NOT NULL,
    `slug`        VARCHAR(64)      NOT NULL,
    `sort_order`  INT              NOT NULL DEFAULT 0,
    `created_at`  DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 标签表
CREATE TABLE IF NOT EXISTS `tags` (
    `id`          BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(64)      NOT NULL,
    `slug`        VARCHAR(64)      NOT NULL,
    `created_at`  DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 文章表
CREATE TABLE IF NOT EXISTS `articles` (
    `id`            BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT,
    `title`         VARCHAR(255)     NOT NULL,
    `slug`          VARCHAR(255)     NOT NULL,
    `content`       MEDIUMTEXT       NOT NULL,
    `summary`       VARCHAR(512)     NOT NULL DEFAULT '',
    `cover_image`   VARCHAR(255)     NOT NULL DEFAULT '',
    `is_published`  TINYINT(1)       NOT NULL DEFAULT 0,
    `pinned`        TINYINT(1)       NOT NULL DEFAULT 0,
    `view_count`    INT UNSIGNED     NOT NULL DEFAULT 0,
    `category_id`   BIGINT UNSIGNED  NOT NULL DEFAULT 0,
    `author_id`     BIGINT UNSIGNED  NOT NULL DEFAULT 1,
    `published_at`  DATETIME         DEFAULT NULL,
    `created_at`    DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_slug` (`slug`),
    KEY `idx_category` (`category_id`),
    KEY `idx_author` (`author_id`),
    KEY `idx_published_at` (`is_published`, `published_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 文章-标签关联表
CREATE TABLE IF NOT EXISTS `article_tags` (
    `article_id`  BIGINT UNSIGNED NOT NULL,
    `tag_id`      BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`article_id`, `tag_id`),
    KEY `idx_tag` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 评论表
CREATE TABLE IF NOT EXISTS `comments` (
    `id`              BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT,
    `article_id`      BIGINT UNSIGNED  NOT NULL,
    `parent_id`       BIGINT UNSIGNED  NOT NULL DEFAULT 0,
    `author_name`     VARCHAR(64)      NOT NULL,
    `author_email`    VARCHAR(128)     NOT NULL DEFAULT '',
    `author_website`  VARCHAR(255)     NOT NULL DEFAULT '',
    `content`         TEXT             NOT NULL,
    `is_approved`     TINYINT(1)       NOT NULL DEFAULT 0,
    `is_admin`        TINYINT(1)       NOT NULL DEFAULT 0,
    `ip`              VARCHAR(45)      NOT NULL DEFAULT '',
    `user_agent`      VARCHAR(255)     NOT NULL DEFAULT '',
    `created_at`      DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_article` (`article_id`, `is_approved`),
    KEY `idx_parent` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 默认管理员 admin / admin123
INSERT INTO `users` (`username`, `password`, `nickname`) VALUES ('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '博主');
