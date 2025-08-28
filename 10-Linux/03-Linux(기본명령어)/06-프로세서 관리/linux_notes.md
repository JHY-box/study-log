##  🔥06. 프로세스 관리


### ⭐ 1. 실행 중인 프로세스 확인 - ps ⭐

Process Status : 현재 실행 중인 프로세스를 확인합니다.

```bash
# 현재 터미널의 프로세스
$ ps
    PID  TTY       TIME  CMD
   1234  pts/0 00:00:00  bash
   1456  pts/0 00:00:00  ps

# 모든 프로세스 상세 정보
$ ps aux
USER    PID %CPU %MEM    VSZ   RSS   TTY STAT START  TIME    COMMAND
root      1  0.0 0.1  167400 11904     ?   Ss 10:00  0:01 /sbin/init
ubuntu 1234  0.0 0.0   21532  5120 pts/0   Ss 10:30  0:00      -bash

# 트리 형태로 프로세스 관계 표시
$ ps auxf

# 특정 프로세스 검색
$ ps aux | grep apache
$ ps aux | grep mysql

# 실시간 프로세스 모니터링
$ top
$ htop # 더 사용자 친화적 (설치 필요)
```

##### ps aux 출력 설명

컬럼 | 의미
---|---
USER | 프로세스 소유자
PID | 프로세스 ID
%CPU | CPU 사용률
%MEM | 메모리 사용률
VSZ | 가상 메모리 크기
RSS | 실제 메모리 사용량
TTY | 터미널
STAT | 프로세스 상태
START | 시작 시간
TIME | CPU 사용 시간
COMMAND | 실행 명령어

### ⭐ 2. 프로세스 종료 - kill ⭐

프로세스를 종료합니다.

```bash
# 기본 종료 시그널 (TERM)
$ kill 1234

# 강제 종료 (KILL)
$ kill -9 1234
$ kill -KILL 1234

# 프로세스명으로 종료
$ killall firefox
$ killall -9 apache2

# 패턴으로 프로세스 찾아 종료
$ pkill -f "python script.py"
$ pkill -u username # 특정 사용자의 모든 프로세스

# 프로세스에 다양한 시그널 전송
$ kill -HUP 1234 # 재시작 시그널
$ kill -USR1 1234 # 사용자 정의 시그널
```

### ⭐ 3. 백그라운드 실행 ⭐

```bash
# 명령어를 백그라운드에서 실행
$ python long_script.py &
[1] 2345

# 실행 중인 작업 확인
$ jobs
[1]+ Running python long_script.py &

# 백그라운드 작업을 포어그라운드로
$ fg %1

# 실행 중인 프로세스를 백그라운드로
# Ctrl+Z로 일시정지 후
$ bg %1

# nohup: 터미널 종료 후에도 계속 실행
$ nohup python long_script.py > output.log 2>&1 &
```