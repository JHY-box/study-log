package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv" // 3단계에서 추가
	"strings" // 3단계에서 추가
)

/** 1. Todo 자원을 나타내는 구조체 정의 */
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

/** 2. 데이터베이스를 대신할 인메모리 슬라이스 */
var todos = []Todo{
	{ID: 1, Title: "Learn Go", Completed: false},
	{ID: 2, Title: "Build REST API", Completed: false},
}

/** 6. PK 역할을 수행할 ID 변수 (2단계에서 추가) */
var nextID = 3

/** 3. /todos 경로의 요청을 처리하는 핸들러 (1~3단계 모두 포함) */
func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// 1단계: GET /todos (전체 조회)
	case http.MethodGet:
		jsonBytes, err := json.Marshal(todos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)

	// 2단계: POST /todos (새 항목 추가) — form 데이터 사용
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		title := r.PostFormValue("title")
		if title == "" {
			http.Error(w, "Title is required", http.StatusBadRequest)
			return
		}
		completed := r.PostFormValue("completed") == "true"

		newTodo := Todo{
			ID:        nextID,
			Title:     title,
			Completed: completed,
		}
		nextID++

		todos = append(todos, newTodo)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newTodo)

	// 3단계: PUT /todos/{id} (수정) — URL에서 ID 파싱
	case http.MethodPut:
		// /todos/{id} 에서 {id} 추출
		idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// form 데이터 사용 (title, completed)
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		title := r.PostFormValue("title")
		completed := r.PostFormValue("completed") == "true"

		for i := range todos {
			if todos[i].ID == id {
				todos[i].Title = title
				todos[i].Completed = completed

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(todos[i])
				return
			}
		}
		http.Error(w, "Todo not found", http.StatusNotFound)

	// 3단계: DELETE /todos/{id} (삭제) — URL에서 ID 파싱
	case http.MethodDelete:
		idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		for i, t := range todos {
			if t.ID == id {
				todos = append(todos[:i], todos[i+1:]...)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		http.Error(w, "Todo not found", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func main() {
	// `/todos/`로 등록해야 /todos/1 같은 path 파라미터를 받음
	http.HandleFunc("/todos/", todosHandler)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}