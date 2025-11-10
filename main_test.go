package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ハンドラー部分を関数化してテスト可能にする
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("APIテスト成功"))
}

// テストコード
func TestHandler_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("期待するステータスコード %d, 実際 %d", http.StatusOK, res.StatusCode)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "APIテスト成功") {
		t.Errorf("レスポンスに期待する文字列が含まれていません: %s", body)
	}
}

func TestHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("期待するステータスコード %d, 実際 %d", http.StatusMethodNotAllowed, res.StatusCode)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "Method not allowed") {
		t.Errorf("エラーメッセージが異なります: %s", body)
	}
}
