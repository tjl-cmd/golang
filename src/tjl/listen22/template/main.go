package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type UserInfo struct {
	Name string
	Sex  string
	Age  int
}

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			fmt.Fprintf(w, "load login.html failed")
			return
		}
		/*u := UserInfo{
			Name: "Mary",
			Sex:  "男",
			Age:  19,
		}*/
		//t.Execute(w, "mary")
		//t.Execute(w, u)
		m := make(map[string]interface{})
		m["username"] = "Mary"
		m["sex"] = "男"
		m["age"] = 19
		t.Execute(w, m)
		t.Execute(os.Stdout, m)
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
