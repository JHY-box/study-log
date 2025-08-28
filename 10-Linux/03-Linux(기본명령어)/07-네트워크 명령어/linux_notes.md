##  🔥07. 네트워크 명령어

### ⭐ 1. 네트워크 연결 확인 - ping ⭐

```bash
# 기본 ping
$ ping google.com
PING google.com (172.217.175.46): 56 data bytes
64 bytes from 172.217.175.46: icmp_seq=0 time=23.123 ms

# 횟수 제한
$ ping -c 4 google.com

# IPv6 사용
$ ping6 google.com

# 큰 패킷 전송
$ ping -s 1024 google.com
```

### ⭐ 2. 네트워크 정보 확인 ⭐

```bash
# IP 주소 확인 (최신 명령어)
$ ip addr show
$ ip a

# 구버전 명령어 (호환성용)
$ ifconfig

# 라우팅 테이블 확인
$ ip route show
$ route -n

# DNS 정보 확인
$ nslookup google.com
$ dig google.com

# 네트워크 연결 상태
$ netstat -tuln
$ ss -tuln # 최신 명령어
```

### ⭐ 3. 파일 다운로드 ⭐

```bash
# wget 사용
$ wget https://example.com/file.zip
$ wget -O custom_name.zip https://example.com/file.zip

# curl 사용
$ curl -O https://example.com/file.zip
$ curl -o custom_name.zip https://example.com/file.zip

# 재귀적 다운로드
$ wget -r -np https://example.com/folder/

# 백그라운드 다운로드
$ wget -b https://example.com/large_file.zip
```