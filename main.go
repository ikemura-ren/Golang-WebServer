package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"time"
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
func ThirdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(4)
		switch num{
		case 0:
			fmt.Fprintf(w, "大吉")
		case 1:
			fmt.Fprintf(w, "中吉")
		case 2:
			fmt.Fprintf(w, "小吉")
		case 3:
			fmt.Fprintf(w, "凶")
		}	
	}
}

func ForceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	http.HandleFunc("/", FirstHandler) // Hello, Worldを返すハンドラ
	http.HandleFunc("/second", SecondHandler) // 値を送ってHello + [param]を返すハンドラ（GET）
	http.HandleFunc("/third", ThirdHandler) // GETリクエストでおみくじ結果を返すハンドラ（GET）
	http.ListenAndServe(":8080", nil)
}