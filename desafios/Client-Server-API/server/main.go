package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type MoedaResponse struct {
	UsdBrl Moeda `json:"USDBRL"`
}

type Moeda struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
}

func main() {	
	resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		panic(err)
	}
	
	var m MoedaResponse
	error = json.Unmarshal(body, &m)
	if error != nil {
		panic(error)
	}
	println(m.UsdBrl)
	
}