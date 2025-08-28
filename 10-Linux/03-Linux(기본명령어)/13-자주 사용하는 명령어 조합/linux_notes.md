##  🔥13. 자주 사용하는 명령어 조합


```bash
# 원라이너 시스템 체크
$ echo "=== 시스템 정보 ===" && uname -a && echo "=== 메모리 ===" && free -h && echo "=== 디스
크 ===" && df -h && echo "=== 부하 ===" && uptime

# 서비스 상태 확인
$ systemctl status apache2 mysql nginx

# 포트 사용 현황
$ netstat -tuln | grep :80
$ ss -tuln | grep :3306
```
