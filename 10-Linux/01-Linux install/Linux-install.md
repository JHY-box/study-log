#### Intel Mac 혹은 Windows https://ubuntu.com/download/server
#### MAC M1/M2/M3/M4 https://ubuntu.com/download/server/arm
```
2025년 업데이트: 이 자료를 작성할 당시 22.04 버전으로 테스트 했으나 2025년 8월 기준 최신 LTS 버전은 24.04 LTS
(Noble Numbat) 입니다. 프로덕션 환경에서는 LTS 버전 사용을 강력히 권장합니다.

```

# 1. 설치 이미지 준비
- 최신 LTS 버전: 24.04 LTS (Noble Numbat) → 10년 지원
- 안정성 중시 시: 22.04 LTS (Jammy Jellyfish)
- 다운로드 링크 제공 (Intel/AMD, Apple Silicon 구분됨)

# 2. 실습 환경 구성
- 실제 서버 환경: 서버 1대 + 여러 클라이언트 접속
- 개인 학습 시: 가상머신(VM) 활용
- Host OS: 실제 컴퓨터 운영체제
- Guest OS: 가상머신에 설치된 운영체제

- 권장 VM 소프트웨어 (2025년 기준)
  - Windows: VMware Workstation Pro (무료, 성능 최고), VirtualBox(무료)
  - Intel Mac: VMware Fusion (무료), Parallels Desktop (유료)
  - Apple Silicon: UTM (무료, ARM 최적화), Parallels Desktop (유료)

# 3. Ubuntu 설치 과정
1. Try or Install Ubuntu Server 선택
2. 언어: English 선택
3. 업데이트: Skip (시간 단축)
4. 키보드: Korean 선택 가능
5. 설치 옵션: 기본값 유지
6. 네트워크 설정 (생략⭐) -그냥 넘어가자.
   - 기본: DHCP(자동 IP)
   - 필요 시: 고정 IP 직접 입력 (Subnet, Gateway 등)
7. 프록시 서버 등록 (기본값 유지) - 추가적인 프록시 서버 연결 없이 다음으로 진행한다.
8. 미러 서버 주소 변경 (기본값 유지) :기본값을 그대로 둔 상태로 다음으로 넘어간다.
9. 디스크 파티셔닝: 기본값(자동)
10. 사용자 계정 입력
11. SSH 설치: 권장
12. 추가 패키지: 선택 안 해도 됨
13. 설치 완료 후 재부팅

# 4. 설치 후
- 로그인: 설치 시 입력한 ID/비밀번호 사용
- 비밀번호 입력 시 화면 표시 안 됨 (보안상 정상)
- 시스템 종료: shutdown now