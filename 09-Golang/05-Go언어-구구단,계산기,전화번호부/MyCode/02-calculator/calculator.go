package main

import "fmt"

// calculate: 두 숫자와 연산자를 받아 결과(float64)와 에러(error)를 반환
func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("0으로 나눌 수 없습니다.")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("유효하지 않은 연산자입니다.")
	}
}

func main() {
	var a, b float64
	var op string

	// 1) 첫 번째 숫자
	fmt.Print("첫 번째 숫자를 입력하세요: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("잘못된 입력입니다. 숫자를 입력해주세요.")
		return
	}

	// 2) 연산자
	fmt.Print("연산자를 입력하세요 (+, -, *, /): ")
	if _, err := fmt.Scan(&op); err != nil {
		// 연산자 입력 단계에서의 일반 입력 에러는 드묾(문자열이므로),
		// 그래도 안전하게 처리
		fmt.Println("유효하지 않은 연산자입니다.")
		return
	}

	// 3) 두 번째 숫자
	fmt.Print("두 번째 숫자를 입력하세요: ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("잘못된 입력입니다. 숫자를 입력해주세요.")
		return
	}

	// 4) 계산 실행
	result, err := calculate(a, b, op)
	if err != nil {
		// 여기서는 0 나누기 또는 잘못된 연산자 메시지가 그대로 출력됨
		fmt.Println(err.Error())
		return
	}

	// 5) 결과 출력 (예시처럼 보기 좋게 소수점 2자리)
	fmt.Printf("결과: %.2f %s %.2f = %.2f\n", a, op, b, result)
}