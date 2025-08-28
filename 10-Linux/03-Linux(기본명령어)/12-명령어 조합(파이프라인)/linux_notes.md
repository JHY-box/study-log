##  🔥12. 명령어 조합 (파이프라인)

### ⭐ 1. 파이프( | ) 활용 ⭐

```bash
# 프로세스 검색 및 종료
$ ps aux | grep apache | grep -v grep | awk '{print $2}' | xargs kill

# 디렉토리별 파일 개수
$ find /home -type f | cut -d/ -f1-3 | sort | uniq -c

# 로그에서 IP별 접속 횟수
$ cat access.log | awk '{print $1}' | sort | uniq -c | sort -nr

# 시스템 리소스 요약
$ df -h | grep -v tmpfs | tail -n +2 | awk '{sum+=$3} END {print "Total used: " sum "KB"}'
```

### ⭐ 2. 리다이렉션 활용 ⭐

```bash
# 출력을 파일로 저장
$ ls -la > file_list.txt
$ ps aux >> process_list.txt # 추가 모드

# 에러 출력 처리
$ find / -name "*.conf" 2>/dev/null > config_files.txt
$ command > output.txt 2>&1 # 표준출력과 에러를 모두 파일로

# 입력 리다이렉션
$ sort < unsorted_list.txt
$ mysql -u user -p database < backup.sql
```
