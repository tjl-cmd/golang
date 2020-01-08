package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			fmt.Fprintf(w, "load login.html failed")
			return
		}

		t.Execute(w, nil)
	} else if method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Printf("username:%s\n", username)
		fmt.Printf("username:%s\n", password)
		if username == "admin" && password == "123456" {
			fmt.Fprintf(w, "username:%s login success\n", username)
		} else {
			fmt.Fprintf(w, "username:%s login load\n", username)

		}
	}
}
func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("listen server failed err:%v\n", err)
		return
	}
}
