package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculate 함수는 두 개의 float64 숫자와 하나의 문자열 연산자를 받아 계산 결과를 반환함.
// 이 단계에서는 에러 처리를 포함하지 않고 순수 계산 로직에 집중함.
func calculate(num1, num2 float64, operator string) float64 {
	var result float64 // 계산 결과를 저장할 
	
	// Go의 switch 문은 Java보다 유연함. break를 명시하지 않아도 자동으로 다음 case로 넘어가지 않음.
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2 // 이 단계에서는 0으로 나누는 경우를 처리하지 않음
	}
	return result // 계산 결과를 반환함
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 첫 번째 숫자 입력 처리
	fmt.Print("첫 번째 숫자를 입력하세요: ")
	input1, _ := reader.ReadString('\n')
	num1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println("잘못된 입력임. 숫자를 입력해야 함.")
		return
	}

	// 연산자 입력 처리
	fmt.Print("연산자를 입력하세요 (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	// 두 번째 숫자 입력 처리
	fmt.Print("두 번째 숫자를 입력하세요:") // 줄바꿈 추가
	input2, _ := reader.ReadString('\n')
	num2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println("잘못된 입력임. 숫자를 입력해야 함.")
		return
	}
	
	// calculate 함수 호출 및 결과 출력
	// calculate 함수가 반환하는 단일 결과 값을 받음.
	result := calculate(num1, num2, operator)
	fmt.Printf("계산 결과: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
