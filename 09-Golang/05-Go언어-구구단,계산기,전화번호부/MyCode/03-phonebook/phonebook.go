// file: phonebook.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 전화번호부: 이름(string) -> 전화번호(string)
var phoneBook = make(map[string]string)

// 한 번 만든 리더를 계속 재사용 (입력 혼선 방지)
var reader = bufio.NewReader(os.Stdin)

// 한 줄 입력 + 공백 제거
func readLine(prompt string) string {
	fmt.Print(prompt)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

// 1. 연락처 추가
func addContact() {
	name := readLine("이름을 입력하세요: ")
	if name == "" {
		fmt.Println("이름이 비어 있습니다. 추가를 취소합니다.")
		return
	}
	phone := readLine("전화번호를 입력하세요: ")
	if phone == "" {
		fmt.Println("전화번호가 비어 있습니다. 추가를 취소합니다.")
		return
	}
	phoneBook[name] = phone
	fmt.Printf("'%s'의 연락처가 추가되었습니다.\n", name)
}

// 2. 연락처 조회
func viewContact() {
	name := readLine("이름을 입력하세요: ")
	if phone, ok := phoneBook[name]; ok {
		fmt.Printf("이름: %s, 전화번호: %s\n", name, phone)
	} else {
		fmt.Println("연락처를 찾을 수 없습니다.")
	}
}

// 3. 연락처 삭제
func deleteContact() {
	name := readLine("삭제할 이름을 입력하세요: ")
	if _, ok := phoneBook[name]; ok {
		delete(phoneBook, name)
		fmt.Printf("'%s'의 연락처가 삭제되었습니다.\n", name)
	} else {
		fmt.Println("연락처를 찾을 수 없습니다.")
	}
}

// 4. 모든 연락처 보기
func listAllContacts() {
	if len(phoneBook) == 0 {
		fmt.Println("등록된 연락처가 없습니다.")
		return
	}
	fmt.Println("-- 모든 연락처 --")
	for name, phone := range phoneBook {
		fmt.Printf("이름: %s, 전화번호: %s\n", name, phone)
	}
}

func main() {
	for {
		// 메뉴 출력
		fmt.Println("\n--- 전화번호부 ---")
		fmt.Println("1. 연락처 추가")
		fmt.Println("2. 연락처 조회")
		fmt.Println("3. 연락처 삭제")
		fmt.Println("4. 모든 연락처 보기")
		fmt.Println("5. 종료")

		// 메뉴 선택 입력(문자열 -> 정수 변환)
		choiceStr := readLine("메뉴를 선택하세요: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("잘못된 입력입니다. 숫자를 입력해주세요.")
			continue
		}

		// 메뉴 분기
		switch choice {
		case 1:
			addContact()
		case 2:
			viewContact()
		case 3:
			deleteContact()
		case 4:
			listAllContacts()
		case 5:
			fmt.Println("프로그램을 종료합니다.")
			return
		default:
			fmt.Println("유효하지 않은 메뉴입니다. 1~5 중에서 선택하세요.")
		}
	}
}
