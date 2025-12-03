package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	// テスト対象のハンドラを呼び出し
	handler := func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("APIテスト成功")); err != nil {
			t.Errorf("レスポンス書き込みエラー: %v", err)
		}
	}

	handler(w, req)

	res := w.Result()
	defer func() {
		if err := res.Body.Close(); err != nil {
			t.Logf("Body.Close エラー: %v", err)
		}
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("レスポンス読み取りエラー: %v", err)
	}

	if string(body) != "APIテスト成功" {
		t.Errorf("期待値と異なります: got %s", string(body))
	}
}

func TestServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello")); err != nil {
			t.Errorf("レスポンス書き込みエラー: %v", err)
		}
	}))
	defer func() {
		t.Log("テストサーバを終了します")
		ts.Close() // Closeはerrorを返さない
	}()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("HTTPリクエストエラー: %v", err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			t.Logf("Body.Close エラー: %v", err)
		}
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("レスポンス読み取りエラー: %v", err)
	}

	if string(body) != "Hello" {
		t.Errorf("期待値と異なります: got %s", string(body))
	}
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 6 // ← これを6に変える

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}