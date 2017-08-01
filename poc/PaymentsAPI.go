package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
)

type PaymentsAPI struct {


}

func (a *PaymentsAPI) getPaymentById( paymentId string ) PaymentDTO{
	conf := getConfigs()

	apiUrlTemplate := conf.get( CONFIG_API_URL) + "payments/%v?access_token=%s"
	accessToken := conf.get( CONFIG_ACCESS_TOKEN)

	apiUrl := fmt.Sprintf(apiUrlTemplate, paymentId, accessToken)

	resp, err := http.Get( apiUrl )

	if err != nil {
		panic( err )
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Println( string(body) )

	var dto PaymentDTO

	err = json.Unmarshal(body, &dto)

	if err != nil {
		panic( err )
	}

	return dto
}
