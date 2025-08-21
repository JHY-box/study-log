package main

import (
	"bufio"   // 버퍼를 사용하여 입출력을 효율적으로 처리하는 기능 제공함
	"fmt"     // 표준 입출력 및 문자열 포맷팅 관련 기능 제공함
	"os"      // 운영체제와 상호작용하는 기능 제공함 (표준 입력/출력 등)
	"strconv" // 문자열을 숫자 타입으로, 또는 그 반대로 변환하는 기능 제공함
	"strings" // 문자열 처리 관련 다양한 함수 제공함
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 1. 첫 번째 숫자 입력 처리
	fmt.Print("첫 번째 숫자를 입력하세요: ")
	input1, _ := reader.ReadString('\n')

	// 문자열을 float64 타입으로 변환함. 64는 비트 크기를 의미함.
	num1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println("잘못된 입력임. 숫자를 입력해야 함.")
		return
	}

	// 2. 연산자 입력 처리
	fmt.Print("연산자를 입력하세요 (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	// 3. 두 번째 숫자 입력 처리
	fmt.Print("두 번째 숫자를 입력하세요: ")
	input2, _ := reader.ReadString('\n')
	num2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println("잘못된 입력임. 숫자를 입력해야 함.")
		return
	}
	
	// 입력받은 값들을 확인용으로 출력함
	fmt.Printf("입력된 값: 숫자1=%.2f, 연산자='%s', 숫자2=%.2f\n", num1, operator, num2)
}
