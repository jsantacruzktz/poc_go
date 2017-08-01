package main

const CONFIG_API_URL string = "apiUrl";
const CONFIG_PAYMENTS_URI string = "paymentsUri"
const CONFIG_ACCESS_TOKEN string = "accessToken"

type Configs struct {

	config map[string]string

}



func (c *Configs) get(key string) string {

	mainConfigs := map[string]string {
		CONFIG_API_URL      : "https://api.mercadolibre.com/",
		CONFIG_PAYMENTS_URI : "payments/",
		CONFIG_ACCESS_TOKEN : "ADM-601-080109-c2e6f235d254eb649770213c2f97a582__N_C__-gmazzaglia-62867623",

	}

	c.config = mainConfigs

	return c.config[key];
}

func getConfigs() Configs{
	return Configs{}
}