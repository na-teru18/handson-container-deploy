package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		message := fmt.Sprintf("APIテスト成功（ver1.0）：%s", currentTime)
		w.Write([]byte(message))
	})

	fmt.Println("Server started on port 80")
	http.ListenAndServe(":80", nil)
}
