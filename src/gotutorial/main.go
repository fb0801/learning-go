package main

import ( 
    "net/http"
    "fmt"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
    http.HandleFunc("/health")
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
    fmt.Println("request made")
}