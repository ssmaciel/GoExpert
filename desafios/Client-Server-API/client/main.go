package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"time"
)


type MoedaResponse struct {
	Valor        string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond * 300))
	defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
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

	valorStr := fmt.Sprintf("Dólar: %v", moedaResponse.Valor)
	println(valorStr)

	
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(valorStr))
	// tamanho, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	f.Close()
}