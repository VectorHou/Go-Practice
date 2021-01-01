package main

import (
	"net/http"
	"fmt"
)

func BcjClient(input []string) (output []bool, err error){
	output = append(output, true)
	err = nil
	return
}

func BcjServer(){
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

func main(){
	BcjServer()
}