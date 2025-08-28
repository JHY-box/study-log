##  🔥05. 파일 권한 관리


### ⭐ 1.권한 확인 ⭐

```bash
$ ls -l example.txt
-rw-r--r-- 1 ubuntu ubuntu 1024 Aug 6 10:30 example.txt
```

##### 권한 표기법 해석
```bash
-rw-r--r--
│││ │││ │││
│││ │││ └── 기타 사용자 (others)
│││ └────── 그룹 (group)
└────────── 소유자 (owner)
```

|문자 | 의미 | 수치
|---|---|---
|r | 읽기 (read) | 4
|w | 쓰기 (write) | 2
|x | 실행 (execute) | 1
|- | 권한 없음 | 0


### ⭐ 2.권한 변경 - chmod ⭐

Change Mode : 파일이나 디렉토리의 권한을 변경합니다.

```bash
# 수치 방식으로 권한 설정
$ chmod 755 script.sh
$ ls -l script.sh
-rwxr-xr-x 1 ubuntu ubuntu 128 Aug 6 10:30 script.sh

# 문자 방식으로 권한 설정
$ chmod u+x script.sh # 소유자에게 실행 권한 추가
$ chmod g-w file.txt # 그룹에서 쓰기 권한 제거
$ chmod o=r file.txt # 기타 사용자에게 읽기 권한만 설정
$ chmod a+r file.txt # 모든 사용자에게 읽기 권한 추가

# 재귀적으로 권한 변경
$ chmod -R 755 /var/www/html/

# 일반적인 권한 설정 예시
$ chmod 644 document.txt # 파일: 소유자 읽기/쓰기, 기타 읽기만
$ chmod 755 directory/ # 디렉토리: 소유자 전체, 기타 읽기/실행
$ chmod 600 private_key # 비밀키: 소유자만 읽기/쓰기
```

### ⭐ 3.소유자 변경 - chown ⭐

Change Owner: 파일이나 디렉토리의 소유자를 변경합니다.

```bash
# 소유자만 변경
$ sudo chown newuser file.txt

# 소유자와 그룹 동시 변경
$ sudo chown newuser:newgroup file.txt

# 그룹만 변경 (chgrp와 동일)
$ sudo chown :newgroup file.txt

# 재귀적으로 변경
$ sudo chown -R apache:apache /var/www/html/

# 참조 파일과 동일하게 설정
$ sudo chown --reference=other_file.txt target_file.txt
```