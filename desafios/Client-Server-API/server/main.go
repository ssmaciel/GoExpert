package main

import (
	"context"
	"encoding/json"
	"net/http"
	//"io"
	"io/ioutil"
	//"os"
	"time"

	"github.com/ssmaciel/GoExpert/desafios/Client-Server-API/server/database"
	"github.com/ssmaciel/GoExpert/desafios/Client-Server-API/server/internal/entity"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MoedaResponse struct {
	USDBRL struct {
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
	} `json:"USDBRL"`
}
func main() {
	http.HandleFunc("/cotacao", handler)
	http.ListenAndServe(":8080", nil)
}

//func main() {
func handler(w http.ResponseWriter, r *http.Request) {	
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond * 200))
	defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		panic(err)
	}
	var moedaResponse MoedaResponse
	error = json.Unmarshal(body, &moedaResponse)
	if error != nil {
		panic(error)
	}
	
	db, err := gorm.Open(sqlite.Open("moeda.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Moeda{})

	moedaDB := database.NewMoeda(db)
	moeda := entity.NewMoeda(
		moedaResponse.USDBRL.Code,
		moedaResponse.USDBRL.Codein,
		moedaResponse.USDBRL.Name,
		moedaResponse.USDBRL.High,
		moedaResponse.USDBRL.Low,
		moedaResponse.USDBRL.VarBid,
		moedaResponse.USDBRL.PctChange,
		moedaResponse.USDBRL.Bid,
		moedaResponse.USDBRL.Ask,
		moedaResponse.USDBRL.Timestamp,
		moedaResponse.USDBRL.CreateDate,

	)
	moedaDB.Create(moeda)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(moeda)
}