package main

import (
	"encoding/json"
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

    var req struct {
        ProductId string `json:"product_id"`
        FirstName string `json:"first_name"`
        LastName string `json:"last_name"`
        Address1 string `json:"address_1"`
        Address2 string `json:"address_2"`
        City string `json:"city"`
        State string `json:"state"`
        Zip string `json:"zip"`
        Country string `json:"country"`

    }

    err := json.NewDecoder(request.Body).Decode(&req)

    if err != nil {
        http.Error(writer, err)
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