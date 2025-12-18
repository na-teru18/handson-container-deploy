package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// バージョン番号
const version = "4.0"

// func unusedFunc() {}

// Go言語では、通常、ファイル内で使われる補助的な関数を main 関数よりも先に定義するため、main 関数の直前に記述
func Add(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		currentTime := time.Now().Format("2006-01-02 15:04:05")
		message := fmt.Sprintf("APIテスト成功（ver%s）：%s", version, currentTime)

		// Writeの返り値をチェック
		if _, err := w.Write([]byte(message)); err != nil {
			log.Printf("レスポンス書き込みエラー: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	fmt.Printf("サーバを起動しました（ポート: 80 var%s）", version)

	// ListenAndServeの返り値をチェック
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("サーバ起動エラー: %v", err)
	}
}
