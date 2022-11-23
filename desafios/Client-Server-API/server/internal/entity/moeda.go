package entity

import "github.com/google/uuid"

type Moeda struct {
	ID    	string
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

func NewMoeda(code string, codein string, name string, high string, low string, varBid string, pctChange string, bid string, ask string, timestamp string, create_date string) *Moeda {
	moeda := &Moeda{
		ID: uuid.New().String(),
		Code: code, 
		Codein: codein,
		Name: name, 
		Low: low, 
		VarBid: varBid,
		PctChange: pctChange,
		Bid: bid, 
		Ask: ask, 
		Timestamp: timestamp,
		CreateDate: create_date,
	}

	return moeda
}