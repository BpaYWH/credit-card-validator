package main

import (
	"credit-card-validator/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type RequestBody struct {
	CardNumber int `json:"card-number"`
}

type LuhnResponse struct {
	Valid    bool   `json:"valid"`
	CardType string `json:"card-type"`
}

// Handler function
func getRoot(writer http.ResponseWriter, _ *http.Request) {
	fmt.Println(("got / request"))
	io.WriteString(writer, "This is my server")
}

func getValidation(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(("got /validation request"))

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(writer, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	fmt.Println(requestBody.CardNumber)

	var jsonResponse []byte

	isCardValid := utils.IsLuhnValid(int64(requestBody.CardNumber))
	if isCardValid {
		jsonResponse, err = json.Marshal(LuhnResponse{Valid: utils.IsLuhnValid(int64(requestBody.CardNumber)), CardType: utils.CheckIIN(int64(requestBody.CardNumber))})
	} else {
		jsonResponse, err = json.Marshal(LuhnResponse{Valid: utils.IsLuhnValid(int64(requestBody.CardNumber)), CardType: utils.CheckIIN(int64(-1))})
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonResponse)
}
