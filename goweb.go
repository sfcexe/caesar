package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const shift = 7

func caesarCipher(text string, shift int) string {
	var result strings.Builder
	shift = shift % 26 // アルファベット文字の数

	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			char = 'a' + (char-'a'+rune(shift))%26
		} else if char >= 'A' && char <= 'Z' {
			char = 'A' + (char-'A'+rune(shift))%26
		}

		result.WriteRune(char)
	}

	return result.String()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	if r.Method == "POST" {
		// フォームデータを取得
		choice := r.FormValue("choice")
		text := r.FormValue("text")

		var result string

		if choice == "encrypt" {
			result = caesarCipher(text, shift)
		} else if choice == "decrypt" {
			result = caesarCipher(text, -shift)
		}

		data := struct {
			Result string
		}{
			Result: result,
		}

		// 結果をテンプレートに埋め込んで表示
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// GETリクエストの場合は単純にテンプレートを表示
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Webサーバーを起動しました。http://localhost:8080 でアクセス可能です。")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
