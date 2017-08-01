package main

type ApiClient struct {

	Payments PaymentsAPI

}

func newApiClient() ApiClient{
	return ApiClient{
		Payments: PaymentsAPI{},
	}
}





