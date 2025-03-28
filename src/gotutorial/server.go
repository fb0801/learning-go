package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"github.com/stripe/stripe-go/v81/paymentintent"
)


func main() {

    stripe.key=""

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
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return 
    }
    params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(req.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},

}
paymentIntent, err := paymentintent.New(params)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
    paymentintent.ClientSecret

    var response struct {
        ClientSecret string `json:"clientSecret"`
    }
    respone.ClientSecret - paymentintent.ClientSecret

    var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")

	_, err = io.Copy(writer, &buf)
	if err != nil {
		fmt.Println(err)
	}
    
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
    response := []byte("server is up and running")

    _, err := writer.Write(response)
    if err != nil{
        fmt.Println(err)
    }

}

func calculateOrderAmount(productId string) int64 {
    switch productId {
	case "Forever Pants":
		return 26000
	case "Forever Shirt":
		return 15500
	case "Forever Shorts":
		return 30000
	}
	return 0
}