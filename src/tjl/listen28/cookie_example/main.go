package main

import (
	"fmt"
	"net/http"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	/*cookies := r.Cookies()
	for index, cookie := range cookies {
		fmt.Printf("index:%d,cookie:%#v\n", index, cookie)
	}
	*/
	C, err := r.Cookie("SessionId")
	fmt.Printf("cookie:%#v\n,err:%v\n", C, err)
	/*
		cookie := &http.Cookie{
			Name:   "sessionId",
			Value:  "lllkkkdkkkdkkdk",
			MaxAge: 3600,
			Domain: "localhost",
			Path:   "/",
		}
		//在数据返回之前设置cookie，才能生效
		http.SetCookie(w, cookie)*/
	w.Write([]byte("hello"))
}
func main() {
	http.HandleFunc("/", indexHandle)
	http.ListenAndServe(":9090", nil)
}
