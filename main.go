package main

import (    
    "os"
    "net/http"
)

func main() {	
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		env_hello := os.Getenv("HELLO_WORLD")	
		request_msg := "# Hello ArgoCD # Env_HELLO_WORLD = " + env_hello
		w.Write([]byte(request_msg))		
	})
	http.ListenAndServe(":8181", nil)
}
