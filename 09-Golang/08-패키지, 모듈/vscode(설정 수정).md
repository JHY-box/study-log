#### 모듈과 패키지 실행을 위한 VSCode 설정 수정

Code Runner 익스텐션은 Alt+Ctrl+N을 누를 때 go 소스파일의 절대 경로를 실행하게 되어 있음
모듈 단위는 명령어가 모듈 디렉토리 위치에서 실행되어야 하므로 설정을 수정함으로서 모듈 폴더 이동 후 코드를 실행하게
변경해야 함
이미 설정 파일에 code-runner.executorMap이 있다면 기존 내용을 수정
```
"code-runner.executorMap": {
    //... 생략 ...
   
    "go": "cd $dir && go run $fileName",

    //... 생략 ...
}
```
<b>없다면 아래 내용만 추가</b>
기존 항목과는 콤마(,)로 구분해야 함

```
code-runner.executorMap": {
    "go": "cd $dir && go run $fileName"
}
```