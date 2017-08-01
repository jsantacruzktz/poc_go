package main

import (
	"time"
	"log"
)

func main() {
/*
	log.Println( "********** STARTING COMMON SEARCH **********")
	executeStraightFowardSearchs()
*/

	log.Println( "\n\n********** STARTING SEARCH ROUTINES **********\n\n")
	executeInRoutines()


}

func executeStraightFowardSearchs() {
	apiClient := newApiClient()

	listOfPaymentsId := getPaymentIds()

	start := time.Now()
	for _, paymentId := range listOfPaymentsId {

		eachSearchStart := time.Now()
		payment := apiClient.Payments.getPaymentById( paymentId )
		eachSearchElapsed := time.Since(eachSearchStart)

		log.Printf( "Responsed: %+v (Took %v) \n", payment, eachSearchElapsed )
	}

	elapsed := time.Since(start)

	log.Printf( "Search all payments took: %v", elapsed)
}



func executeInRoutines() {
	getPaymentIds := getPaymentIds()

	mainChan := make(chan PaymentDTO)

	for _, eachPaymentId := range getPaymentIds {

		go func(paymentId string){
			apiClient := newApiClient()
			eachSearchStart := time.Now()
			payment := apiClient.Payments.getPaymentById( paymentId )
			eachSearchElapsed := time.Since(eachSearchStart)

			log.Printf( "Responsed: %+v (Took %v) \n", payment, eachSearchElapsed )

			mainChan <- payment
		}(eachPaymentId)
	}

	start := time.Now()
	for i := 1; ; i++ {
		log.Printf( "Routine just posted (%d): %+v  \n", i, <-mainChan)

		if(i == len(getPaymentIds)){
			elapsed := time.Since(start)
			log.Printf( "Routines have done the work took %v. \n", elapsed )
			break
		}
	}
}

func getPaymentIds() []string {
	return []string{"857912893", "841528742", "857806577", "841528756", "857912914", "857806596", "857912918", "854328788", "841463883", "841463875"}
}
