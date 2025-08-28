##  🔥09. 압축 및 아카이브

### ⭐ 1. tar 아카이브 ⭐

```bash
# 압축 생성
$ tar -czf archive.tar.gz folder/
$ tar -czf backup_$(date +%Y%m%d).tar.gz /home/ubuntu/

# 압축 해제
$ tar -xzf archive.tar.gz

# 압축 내용 확인 (해제하지 않고)
$ tar -tzf archive.tar.gz

# 특정 디렉토리에 압축 해제
$ tar -xzf archive.tar.gz -C /tmp/

# 진행 상황 표시
$ tar -czvf archive.tar.gz folder/
```

### ⭐ 2. zip 압축 ⭐

```bash
# zip 압축 생성
$ zip -r archive.zip folder/

# 압축 해제
$ unzip archive.zip

# 특정 디렉토리에 해제
$ unzip archive.zip -d /tmp/

# 압축 내용 확인
$ unzip -l archive.zip
```
