package main

import (
	"log"
	"github.com/hashicorp/vault/api"
)

var VClient *api.Client // global variable

func InitVault(token string) error {
	conf := &api.Config{
		Address: "http://127.0.0.1:8200",
	}

	client, err := api.NewClient(conf)
	if err != nil {
		return err
	}
	VClient = client

	VClient.SetToken(token)

	return nil
}

func main() {
	err := InitVault("s.sCRAYO9kFN3OuQQQOYOrsqdB")
	if err != nil {
		log.Println(err)
	}
	c := VClient.Logical()

	_, err = c.Write("secret/hello", map[string]interface{}{
		"name":     "Priyanshi",
		"username": "priya12",
		"password": "admin1234",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println("Success: Data stored in secret/hello")
}
