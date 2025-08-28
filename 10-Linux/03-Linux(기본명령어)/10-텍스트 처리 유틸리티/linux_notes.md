##  🔥10. 텍스트 처리 유틸리티

### ⭐ 1. 텍스트 정렬 및 중복 제거 ⭐

```bash
# 파일 내용 정렬
$ sort names.txt
$ sort -n numbers.txt # 숫자 정렬
$ sort -r names.txt # 역순 정렬

# 중복 제거
$ sort names.txt | uniq
$ sort names.txt | uniq -c # 개수와 함께

# 직접 중복 제거
$ sort -u names.txt
```

### ⭐ 2. 텍스트 변환 ⭐

```bash
# 문자 변환
$ tr 'a-z' 'A-Z' < file.txt # 소문자를 대문자로
$ tr -d ' ' < file.txt # 공백 제거

# 열 추출
$ cut -d: -f1 /etc/passwd # : 구분자로 첫 번째 필드
$ cut -c1-10 file.txt # 1-10번째 문자

# 워드 카운트
$ wc file.txt # 줄, 단어, 문자 수
$ wc -l file.txt # 줄 수만
$ wc -w file.txt # 단어 수만
```