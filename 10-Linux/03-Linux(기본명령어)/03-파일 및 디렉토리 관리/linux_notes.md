##  🔥03. 파일 및 디렉토리 관리 

### ⭐ 1. 디렉토리 생성 - mkdir ⭐

Make Directory: 새 디렉토리를 생성합니다.

```bash
# 단일 디렉토리 생성
$ mkdir project
$ ls
project

# 여러 디렉토리 동시 생성
$ mkdir web mobile api
$ ls
api  mobile  project  web

# 중간 경로까지 자동 생성
$ mkdir -p project/frontend/src/components
$ ls -R project/
project/:
frontend

project/frontend:
src

project/frontend/src:
components

project/frontend/src/components:

# 권한 지정하여 생성
$ mkdir -m 755 public_folder
```

### ⭐ 2. 파일 생성 - touch ⭐

Touch: 빈 파일을 생성하거나 파일의 타임스탬프를 변경합니다.

```bash
# 빈 파일 생성
$ touch index.html
$ ls -l index.html
-rw-r--r-- 1 ubuntu ubuntu 0 Aug  6 10:45 index.html

# 여러 파일 동시 생성
$ touch style.css script.js README.md
$ ls
index.html  script.js  style.css  README.md

# 파일 시간 업데이트
$ touch existing_file.txt

# 특정 시간으로 설정
$ touch -t 202501010000 newfile.txt
```

### ⭐ 3. 파일/디렉토리 복사 - cp ⭐

Copy: 파일이나 디렉토리를 복사합니다.

```bash
# 파일 복사
$ cp index.html backup.html
$ ls
backup.html  index.html

# 디렉토리로 복사
$ mkdir backup
$ cp index.html backup/
$ ls backup/
index.html

# 여러 파일 복사
$ cp *.html *.css backup/

# 디렉토리 전체 복사
$ cp -r project project_backup
$ ls
project  project_backup

# 원본 속성 유지하여 복사
$ cp -p original.txt copy.txt

# 대상덮어쓰기 확인(묻기)
$ cp -i file1.txt file2.txt
cp: overwrite 'file2.txt'? y
```

### ⭐ 4. 파일/디렉토리 이동/이름변경 - mv ⭐

Move: 파일을 이동하거나 이름을 변경합니다.

```bash
# 파일 이름 변경
$ mv old_name.txt new_name.txt
$ ls
new_name.txt

# 파일 이동
$ mkdir documents
$ mv new_name.txt documents/
$ ls documents/
new_name.txt

# 여러 파일 이동
$ mv *.txt docs/
# 디렉토리 이름 변경
$ mv documents docs
$ ls
docs

# 대화형 이동 (덮어쓰기 확인)
$ mv -i source.txt target.txt
```

### ⭐ 5. 파일/디렉토리 삭제 - rm ⭐

Remove: 파일이나 디렉토리를 삭제합니다.

```bash
# 파일 삭제
$ rm unwanted_file.txt

# 여러 파일 삭제
$ rm file1.txt file2.txt
$ rm *.bak

# 대화형 삭제 (확인 요청)
$ rm -i important_file.txt
rm: remove regular file 'important_file.txt'? n

# 강제 삭제 (확인 없이)
$ rm -f file.txt

# 디렉토리 삭제 (재귀적)
$ rm -r old_directory
$ rm -rf unwanted_folder   # 강제 삭제

# 빈 디렉토리만 삭제
$ rmdir empty_folder
```

⚠ 주의: `rm -rf`는 매우 강력한 명령어입니다. 특히 `rm -rf /` 같은 명령은 시스템을 파괴할 수 있으니 주의하세요.