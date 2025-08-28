##  🔥04. 파일 내용 확인 및 편집

### ⭐ 1. 파일 전체 내용 보기 - cat ⭐

Concatenate: 파일 내용을 출력합니다.

```bash
# 파일 내용 출력
$ cat /etc/os-release
NAME="Ubuntu"
VERSION="22.04.3 LTS (Jammy Jellyfish)"
ID=ubuntu
ID_LIKE=debian

# 여러 파일 연결하여 출력
$ cat file1.txt file2.txt

# 줄 번호와 함께 출력
$ cat -n /etc/passwd
 1 root:x:0:0:root:/root:/bin/bash
 2 daemon:x:1:1:daemon:/usr/sbin:/usr/sbin/nologin

# 빈 줄 압축하여 출력
$ cat -s file_with_empty_lines.txt

# 탭 문자를 ^I로 표시
$ cat -T file_with_tabs.txt
```

### ⭐ 2. 파일 앞부분 보기 - head ⭐

파일의 처음 몇 줄을 출력합니다

```bash
# 기본 10줄 출력
$ head /var/log/syslog

# 특정 줄 수 출력
$ head -n 5 /etc/passwd
$ head -5 /etc/passwd     # 동일한 결과

# 바이트 단위로 출력
$ head -c 100 file.txt

# 여러 파일의 헤더 출력
$ head -3 *.txt
==> file1.txt <==
line 1
line 2
line 3

==> file2.txt <==
another line 1
another line 2
another line 3
```

### ⭐ 3. 파일 뒷부분 보기 - tail ⭐

파일의 마지막 몇 줄을 출력합니다.

```bash
# 기본 10줄 출력
$ tail /var/log/syslog

# 특정 줄 수 출력
$ tail -n 20 /var/log/auth.log

# 실시간 로그 모니터링 (매우 유용!)
$ tail -f /var/log/syslog

# Ctrl+C로 중단
# 여러 파일 동시 모니터링
$ tail -f /var/log/syslog /var/log/auth.log

# 특정 줄부터 끝까지 출력
$ tail -n +10 file.txt # 10번째 줄부터 끝까지
```

### ⭐ 4. 페이지 단위로 보기 - less/more ⭐

긴 파일을 페이지 단위로 나누어 봅니다.

```bash
# less 사용 (권장)
$ less /var/log/syslog

# more 사용
$ more /etc/passwd
```

##### less 내부 명령어

키 | 기능
---|---
스페이스 | 다음 페이지
b | 이전 페이지
/ | 앞으로 이동
? | 뒤로 검색
검색어 | 위로 이동
n | 다음 검색 결과
N | 이전 검색 결과
q | 파일 종료
h | 파일 시작으로
홈 | 파일 시작으로

### ⭐ 5. 파일에서 문자열 검색 - grep ⭐

Global Regular Expression Print : 파일에서 특정 문자열을 검색합니다

```bash
# 기본 검색
$ grep "ubuntu" /etc/passwd
ubuntu:x:1000:1000:Ubuntu:/home/ubuntu:/bin/bash

# 대소문자 구분 없이 검색
$ grep -i "ubuntu" /etc/passwd

# 줄 번호와 함께 출력
$ grep -n "root" /etc/passwd
1:root:x:0:0:root:/root:/bin/bash

# 매칭되지 않는 줄 출력
$ grep -v "nologin" /etc/

# 여러 파일에서 검색
$ grep "error" /var/log/*.log

# 재귀적 검색 (디렉토리 내 모든 파일)
$ grep -r "TODO" /home/ubuntu/project/

# 정확한 단어만 매칭
$ grep -w "cat" /etc/passwd

# 문자열 앞뒤 줄도 함께 출력
$ grep -A 2 -B 2 "error" /var/log/syslog

# -A: After (뒤 2줄), -B: Before (앞 2줄)
# 매칭된 파일명만 출력
$ grep -l "error" /var/log/*.log

# 매칭 개수 출력
$ grep -c "ubuntu" /etc/passwd
```