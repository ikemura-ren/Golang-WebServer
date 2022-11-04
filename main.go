package main

import (
	"fmt"
	"net/http"
)

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func SecondHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello" + r.FormValue("hoge"))
	}
	if r.Method == "POST" {
		r.ParseForm()
		text := r.Form.Get("text")
		fmt.Fprintf(w, "Hello" + text)
	}
}

func main() {
	http.HandleFunc("/", FirstHandler) // Hello, Worldを返すハンドラ
	http.HandleFunc("/second", SecondHandler) // 値を送ってHello + [param]を返すハンドラ（GET）
	http.ListenAndServe(":8080", nil)
}