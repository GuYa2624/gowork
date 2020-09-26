package main

import (
	"fmt"
	"net/http"
	_"github.com/go-sql-driver/mysql"
)

type MyHandler struct {}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的处理器处理请求")
}

func main()  {
	myHandler := MyHandler{}

	http.Handle("/myHandler", &myHandler)

	http.ListenAndServe(":8080", nil)
}
