package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // Parse Request
	fmt.Println(r.Form) // Form info
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	// logic
	var index uint32 = (uint32)(rand.Uint32()) % 3
	var response string = "<p>Hello, World!</p><img src='http://localhost:8889/images/img_%d.jpg'>"
	fmt.Fprintf(w, "<html><body>"+response+"</body></html>\n", index)
}
func main() {
	http.HandleFunc("/abc", sayHelloName)           // assign callback
	err := http.ListenAndServe("0.0.0.0:9090", nil) // assign address
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
