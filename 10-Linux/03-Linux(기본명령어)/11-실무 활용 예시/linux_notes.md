##  🔥11. 실무 활용 예제

### ⭐ 1. 로그 분석 ⭐

```bash
# 에러 로그 실시간 모니터링
$ tail -f /var/log/apache2/error.log | grep "ERROR"

# 특정 날짜의 로그 추출
$ grep "2025-08-06" /var/log/syslog

# 가장 많이 접속한 IP 찾기
$ awk '{print $1}' /var/log/apache2/access.log | sort | uniq -c | sort -nr | head -10

# 로그 파일 크기별 정렬
$ du -h /var/log/*.log | sort -hr
```

### ⭐ 2. 시스템 관리 ⭐

```bash
# 디스크 사용량이 많은 디렉토리 찾기
$ du -h / 2>/dev/null | sort -hr | head -20

# 메모리를 많이 사용하는 프로세스
$ ps aux --sort=-%mem | head -10

# 오래된 파일 찾기 및 삭제
$ find /tmp -type f -mtime +7 -delete

# 대용량 파일 찾기
$ find / -size +100M -type f 2>/dev/null
```

### ⭐ 3. 백업 스크립트 ⭐

```bash
# 일일 백업 스크립트
"   #!/bin/bash   " 
DATE=$(date +%Y%m%d)
tar -czf /backup/website_$DATE.tar.gz /var/www/html/
find /backup -name "website_*.tar.gz" -mtime +7 -delete
```