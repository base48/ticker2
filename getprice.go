package main

import (
		"encoding/json"
		"github.com/go-resty/resty/v2"
)

func GetPrice() string {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://api.bitfinex.com/v1/pubticker/btcusd")
	if err == nil {
		type Pars struct {
			Last	string `json:"last_price"`
		}

		pars := Pars{}
		json.Unmarshal(resp.Body(), &pars)
		return pars.Last
	}
	return "fuck up"
}

