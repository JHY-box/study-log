package main

import "fmt"

type Notifier interface {
	Send(message string) // <--- 여기 !!! ⭐⭐⭐
}

type EmailNotifier struct {
	Recipient string
}

func (e *EmailNotifier) Send(message string) {
	fmt.Printf("'%s' 주소로 이메일 발송: %s", e.Recipient, message)
}
func main() {

	var n Notifier // <-- 여기!! ⭐⭐⭐

	n = &EmailNotifier{Recipient: "test@example.com"}

	n.Send("안녕하세요! Go 인터페이스 테스트입니다.")
}
