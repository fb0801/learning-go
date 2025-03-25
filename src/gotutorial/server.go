package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
    http.HandleFunc("/health", handleHealth)
    
    log.Println(v...:"listening on localhost")
    var err error = http.ListenAndServe(addr: "localhost:4242", handler:nil)

    if err != nil {
        log.Fatal(err)
    }
}

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
    if request.Method != "POST" {
        http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
        return 
    }
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
    response := []byte("server is up and running")

    _, err := writer.Write(response)
    if err != nil{
        fmt.Println(err)
    }

}