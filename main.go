package main

import (
	"fmt"
	"net/http"
)

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func SecondHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello" + r.FormValue("hoge"))
}

func main() {
	http.HandleFunc("/", FirstHandler) // Hello, Worldを返すハンドラ
	http.HandleFunc("/second", SecondHandler) // 値を送ってHello + [param]を返すハンドラ（GET）
	http.ListenAndServe(":8080", nil)
}