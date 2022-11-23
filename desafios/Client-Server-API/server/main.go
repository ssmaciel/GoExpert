package main

import (
	"context"
	"encoding/json"
	"net/http"
	"io"
	"io/ioutil"
	"os"
	"time"
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
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond * 20000))
	defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	//req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
	body, error := ioutil.ReadAll(res.Body)
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