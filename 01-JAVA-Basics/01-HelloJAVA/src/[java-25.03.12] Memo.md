# Java 시작하기

```java
public class App2 {

    // 메서드 블록 --> 프로그램의 시작점
    public static void main(String[] args) {

        String a = "Jayyyyy";
        String b = "Gooooooo";
        // 괄호 안의 메시지를 터미널에 메시지를 출력하는 명령어
        System.out.println("Hello, World!");
        System.out.println("안녕하세요. 자바!!");

        Java(b);
        Go(a);
    }

    public static void Java (String zzz) {
        // 괄호 안의 메시지를 터미널에 메시지를 출력하는 명령어
        System.out.println("zzz의 값: " + zzz);
        System.out.println("Hello, World!");
        System.out.println("안녕하세요. 자바!!1111111111");
    }

    public static void Go (String b) {
        // 괄호 안의 메시지를 터미널에 메시지를 출력하는 명령어
        System.out.println("String b의 값: " + b);
        System.out.println("Hello, World!");
        System.out.println("안녕하세요. 자바!!2222222222");
    }
}
```
### 당연히 알고 있어야하는것들.
 <b>1. public에 class는 단 하나만 존재.⭐</b>
    - public 클래스는 한 파일에 단 하나만 존재할 수 있다.

 <b>2. public static void main(String[] args)</b>
    - Java는 무조건 main 메서드가 먼저 실행됨.⭐

 <b>3. public static void main(String[] args)
       public static void Java (String zzz)
       public static void Go (String b)</b>
    - 여기서 main, Java, Go 의 명칭은 ⭐메서드(method)라고 불림.
    - (String zzz), (String b)는 ⭐매개변수(parameter)

 <b>4. System.out.println("Hello, World!");</b>
    - println  =  print + Line  =  "출력하고 줄바꿈해라" 라는 뜻.

### 실행 순서
1. main 메서드 실행
 - a에 "Jayyyyy" 저장
 - b에 "Gooooooo" 저장
 - "Hello, World!" 출력
 - "안녕하세요. 자바!!" 출력

2. Java(b) 호출
 - b 의 값 "Gooooooo"가 Java 메서드의 zzz로 전달됨
 - "zzz의 값: Gooooooo" 출력
 - "Hello, World!" 출력
 - "안녕하세요. 자바!!1111111111" 출력

3. Go(a) 호출
 - a 의 값 "Jayyyyy"가 Go 메서드의 b로 전달됨
 - "String b의 값: Jayyyyy" 출력
 - "Hello, World!" 출력
 - "안녕하세요. 자바!!2222222222" 출력
 


#### 정리
 - main에서 두 문자열을 출력
 - Java 메서드에 "Gooooooo"를 전달해서 관련 메시지 출력
 - Go 메서드에 "Jayyyyy"를 전달해서 관련 메시지 출력
 - 각 메서드에서 전달받은 값을 출력함
 - 변수 이름과 메서드 매개변수 이름은 달라도 값이 잘 전달됨