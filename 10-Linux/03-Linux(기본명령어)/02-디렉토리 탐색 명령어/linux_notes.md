## 🔥 02. 디렉토리 탐색 명령어

### ⭐ 1. 현재 위치 확인 - pwd ⭐

Print Working Directory: 현재 작업 중인 디렉토리 경로를 출력합니다.

```bash
# 현재 위치 확인
$ pwd
/home/ubuntu

# 심볼릭 링크의 실제 경로 확인
$ pwd -P
/home/ubuntu
```

### ⭐ 2. 디렉토리 이동 - cd ⭐

Change Directory: 디렉토리를 이동합니다.

```bash
# 홈 디렉토리로 이동
$ cd
$ pwd
/home/ubuntu

# 홈 디렉토리로 이동(동일)
$ cd ~
$ pwd
/home/ubuntu

# 특정 디렉토리로 이동
$ cd /home/ubuntu
$ pwd
/home/ubuntu

# 특정 디렉토리로 이동
$ cd /var/log
$ pwd
/var/log

# 상위 디렉토리로 이동
$ cd ..
$ pwd
/var

# 이전 디렉토리로 이동
$ cd -
$ pwd
/var/log

# 루트 디렉토리로 이동
$ cd /
$ pwd
/
```

#### 상대 경로 vs 절대 경로

```bash
# 절대 경로 (루트부터 시작)
$ cd /home/ubuntu/Documents

# 상대 경로 (현재 위치 기준)
$ cd Documents        # 현재 위치의 Documents 폴더로 이동
$ cd ./Downloads      # 현재 폴더의 Downloads 폴더로 이동
$ cd ../../etc        # 두 단계 상위의 etc 폴더로 이동
```



## ⭐ 03. 디렉토리 내용 확인 - ls ⭐

List: 디렉토리 내용을 출력합니다.

```bash
# 기본 출력
$ ls
Desktop  Documents  Downloads  Pictures

# 상세 정보 출력
$ ls -l
total 16
drwxr-xr-x 2 ubuntu ubuntu 4096 Aug  6 10:30 Desktop
drwxr-xr-x 2 ubuntu ubuntu 4096 Aug  6 10:30 Documents
drwxr-xr-x 2 ubuntu ubuntu 4096 Aug  6 10:30 Downloads
drwxr-xr-x 2 ubuntu ubuntu 4096 Aug  6 10:30 Pictures

# 숨김 파일까지 표시
$ ls -a
.  ..  .bashrc  .profile  Desktop  Documents  Downloads

# 파일 크기를 읽기 쉬운 단위로
$ ls -lh
total 16K
drwxr-xr-x 2 ubuntu ubuntu 4.0K Aug  6 10:30 Desktop
drwxr-xr-x 2 ubuntu ubuntu 4.0K Aug  6 10:30 Documents

# 숨김 파일까지 상세 정보 출력
$ ls -al

# 최신 수정 순으로 정렬
$ ls -lt

# 크기 순으로 정렬
$ ls -lS

# 확장자/파일타입별 색상 표시(환경에 따라 다름)
$ ls --color=auto
```

### ls -l 출력 정보 해석

```bash
drwxr-xr-x 2 ubuntu ubuntu 4096 Aug  6 10:30 Documents
```

부분 | 의미 | 설명
---|---|---
첫 글자 (예: `d`) | 파일 타입 | `d` = 디렉토리, `-` = 일반 파일, `l` = 심볼릭 링크
권한 (`rwxr-xr-x`) | 권한 | 사용자/그룹/기타 순서로 읽기(r), 쓰기(w), 실행(x) 권한 표시
숫자 (`2`) | 링크 수 | 하드링크 수(디렉토리의 경우 서브디렉토리 수 등)
소유자 | 소유자, 그룹 | 소유자 계정명 / 그룹명 (`ubuntu ubuntu`)
크기 (`4096`) | 파일 크기 | 바이트 단위 (옵션 `-h`로 읽기 쉬운 단위 표시)
수정 시간 (`Aug 6 10:30`) | 마지막 수정 시각 | 최근 수정된 날짜와 시간
이름 (`Documents`) | 파일/디렉토리 이름 | 항목 이름

참고
- 출력 상단의 `total 16`은 블록 단위 총합(파일 시스템에 따라 다름)입니다.
- 권한 문자열을 자세히 보면 소유자·그룹·기타의 세 부분으로 나뉘며, 필요한 경우 `chmod`, `chown` 명령으로 변경합니다.