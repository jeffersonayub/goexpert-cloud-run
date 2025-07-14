package entity

import (
	"encoding/json"
	"net/http"
)

type Cep struct {
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro,omitempty"`
}

func IsValidCEP(cep string) bool {
	if len(cep) != 8 {
		return false
	}
	for _, c := range cep {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func GetCep(cep string) (localidade string, erro bool, err error) {

	responseCep, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return "", true, err
	}
	defer responseCep.Body.Close()

	var result Cep
	err = json.NewDecoder(responseCep.Body).Decode(&result)
	if err != nil {
		return "", true, err
	}

	return result.Localidade, result.Erro, nil
}
