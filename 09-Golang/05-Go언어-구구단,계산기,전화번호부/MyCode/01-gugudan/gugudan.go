package main

import "fmt"

func main() {
	var dan int

	// 1. 안내 메시지 출력
	fmt.Print("출력할 단을 입력하세요 (2~9): ")

	// 2. 사용자 입력 받기
	fmt.Scan(&dan)

	// 3. 입력값 검사 (2~9 범위 확인)
	if dan < 2 || dan > 9 {
		fmt.Println("잘못된 입력입니다. 2에서 9 사이의 숫자를 입력해주세요.")
		return // 프로그램 종료
	}

	// 4. 구구단 출력
	fmt.Printf("--- %d단 ---\n", dan)
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d x %d = %2d\n", dan, i, dan*i)
	}
}