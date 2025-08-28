##  🔥08. 시스템 정보 확인

### ⭐ 1. 시스템 리소스 ⭐

```bash
# 메모리 사용량
$ free -h
          total   used   free shared buff/cache available
   Mem:   3.8Gi  1.2Gi  1.8Gi  180Mi      800Mi     2.4Gi
  Swap:   2.0Gi     0B  2.0Gi

# 디스크 사용량
$ df -h
 Filesystem    Size  Used  Avail  Use%  Mounted on
 /dev/sda1      20G  8.5G    11G   45%           /
 tmpfs         1.9G     0   1.9G    0%    /dev/shm

# 디렉토리 크기
$ du -sh /var/log
156M /var/log
$ du -h --max-depth=1 /var/
4.0K /var/backups
156M /var/log
8.0K /var/mail

# CPU 정보
$ lscpu
Architecture: x86_64
CPU op-mode(s): 32-bit, 64-bit
CPU(s): 2

# 시스템 부하
$ uptime
 10:45:23 up 2:30, 1 user, load average: 0.15, 0.20, 0.18

# 실시간 시스템 모니터링
$ htop
$ top
```

### ⭐ 2. 시스템 정보 ⭐

```bash
# 운영체제 정보
$ uname -a
Linux ubuntu-vm 5.15.0-78-generic #85-Ubuntu SMP Fri Jul 7 15:25:09 UTC 2023 x86_64 x86_64
x86_64 GNU/Linux

# 배포판 정보
$ lsb_release -a
$ cat /etc/os-release

# 하드웨어 정보
$ lshw -short
$ lsblk   # 블록 디바이스
$ lsusb   # USB 디바이스
$ lspci   # PCI 디바이스
```


