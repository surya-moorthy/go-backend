package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter,r *http.Request) {
	 if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"parseform() error: %v",err)
	 }

	 fmt.Fprintf(w,"Post Request Successful")
    
	 name := r.FormValue("name")
	 address := r.FormValue("address")

	 fmt.Fprintf(w,"name : %s \n address : %s \n",name,address )
}

func helloHandler(w http.ResponseWriter,r *http.Request)  {
  if r.URL.Path != "/hello" {
	http.Error(w,"404 not found",http.StatusNotFound)
	return
  }

  if r.Method != "GET" {
	http.Error(w,"method is not supported",http.StatusNotFound)
	return
  }

  fmt.Fprintf(w,"hello!")
}

func main() {
    fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080",nil); err != nil {
           log.Fatal(err)
	}

}