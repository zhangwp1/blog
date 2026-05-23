import dayjs from 'dayjs';

export function formatDate(date: string | null | undefined): string {
  if (!date) return '';
  return dayjs(date).format('YYYY年MM月DD日');
}

export function formatDateTime(date: string | null | undefined): string {
  if (!date) return '';
  return dayjs(date).format('YYYY-MM-DD HH:mm');
}

export function formatDateShort(date: string | null | undefined): string {
  if (!date) return '';
  return dayjs(date).format('YYYY-MM-DD');
}
