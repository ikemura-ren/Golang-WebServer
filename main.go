package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, 世界")
}

func main() {
  	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.HandleFunc("/test", handler2) 
	http.Handle("/tes", String("こんにちは"))  
	http.ListenAndServe(":8080", nil)
}