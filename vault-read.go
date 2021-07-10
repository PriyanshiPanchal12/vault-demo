package main

import (
	"fmt"
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

	secret, err := c.Read("secret/hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	// secret, err = c.Write("secret/hello", "name: Priyanshi")
	// if err != nil {
	// 	log.Println(err)
	// }
	fmt.Println("secret", secret)
}