package main

import "net/http"

func main() {	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World / Hello ArgoCD"))		
	})
	http.ListenAndServe(":8181", nil)
}
