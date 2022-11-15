package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"time"
	"strconv"  
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
	header := r.Header
	fmt.Fprintln(w,header)

}
func FifthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		a := r.Form.Get("num")
		b, _ := strconv.Atoi(a)
		num := 100 + b
		fmt.Fprint(w, num)
	}
}

func SixthHandler(w http.ResponseWriter, r *http.Request) {
	//リクエストボディを返すハンドラ
	body := r.Body
	fmt.Fprintln(w, body)
}


func main() {
	http.HandleFunc("/", FirstHandler) // Hello, Worldを返すハンドラ
	http.HandleFunc("/second", SecondHandler) // 値を送ってHello + [param]を返すハンドラ（GET）
	http.HandleFunc("/third", ThirdHandler) // GETリクエストでおみくじ結果を返すハンドラ（GET）
	http.HandleFunc("/force",ForceHandler) // GETリクエストでリクエストヘッダーの情報を返すハンドラ
	http.HandleFunc("/fifth",FifthHandler) // POSTリクエストから数値を受け取って計算して返すハンドラ
	http.HandleFunc("/sixth",SixthHandler) // GETリクエストでリクエストボディの情報を返すハンドラ
	http.ListenAndServe(":8080", nil)
}