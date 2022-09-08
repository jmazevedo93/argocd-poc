package main

import (
    "fmt"
    "os"
    "net/http"
)

func main() {
	fmt.Println("BAR:", os.Getenv("AWS_SECRET_ACCESS_KEY"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello Full Cycle!!!</h1>"))
		fmt.Println("BAR:", )
	})
	http.ListenAndServe(":8090", nil)
}
