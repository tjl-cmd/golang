package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                //解析参数，默认是不会解析
	fmt.Fprintf(w, "%v", r.Form) //这些信息输出到服务器端的打印信息
	fmt.Fprintf(w, "path:%s\n", r.URL.Path)
	fmt.Fprintf(w, "schema:%s\n", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, " "))
	}
	json.Marshal(w)
	fmt.Fprintf(w, "hello woman") //这个写到w的是输出到客户端的

}
func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("listenAndServe:", err)
	}
}
