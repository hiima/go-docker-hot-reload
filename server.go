package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ホットリロードの挙動を試したいときは、このメッセージを適当なものに変更して保存すれば良い
		// 保存したときにホットリロードが走り、 http://localhost:9999 で変更が反映されていれば OK
		fmt.Fprintf(w, "Hello, World")
	})

	fmt.Println("open http://localhost:9999")
	http.ListenAndServe(":9999", nil)
}
