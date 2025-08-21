package main

import (
	"bufio"   // 버퍼를 사용하여 입출력을 효율적으로 처리하는 기능 제공함
	"fmt"     // 표준 입출력 및 문자열 포맷팅 관련 기능 제공함
	"os"      // 운영체제와 상호작용하는 기능 제공함 (표준 입력/출력 등)
	"strconv" // 문자열을 숫자 타입으로, 또는 그 반대로 변환하는 기능 제공함
	"strings" // 문자열 처리 관련 다양한 함수 제공함
)

func main() {
	// 사용자에게 구구단 숫자를 입력받도록 안내함
	fmt.Print("구구단을 출력할 숫자를 입력하세요 (2-9): ")

	// 표준 입력(키보드)으로부터 데이터를 읽기 위한 Reader 객체를 생성함
	reader := bufio.NewReader(os.Stdin)

	// Reader로부터 줄바꿈 문자(\n)가 나올 때까지 한 줄을 읽어옴
	input, _ := reader.ReadString('\n')

	// 읽어온 문자열의 앞뒤 공백(줄바꿈 포함)을 제거함
	input = strings.TrimSpace(input)

	// 입력받은 문자열을 정수(int)로 변환함
	// strconv.Atoi는 문자열을 정수로 변환하며, 변환 실패 시 에러를 함께 반환함
	num, err := strconv.Atoi(input)

	// 에러가 발생했는지 확인함 (즉, 입력이 유효한 숫자가 아닌 경우)
	if err != nil {
		fmt.Println("잘못된 입력임. 숫자를 입력해야 함.")
		return // 에러 메시지 출력 후 프로그램 종료
	}

	// 입력된 숫자가 2에서 9 사이의 유효한 범위에 있는지 확인함
	if num < 2 || num > 9 {
		fmt.Println("2에서 9 사이의 숫자를 입력해야 함.")
		return // 유효하지 않은 범위일 경우 메시지 출력 후 종료
	}

	// 구구단 시작을 알리는 메시지를 출력함
	fmt.Printf("--- %d단 ---\n", num)
	
	// 1부터 9까지 반복하며 구구단 결과를 출력함
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}
