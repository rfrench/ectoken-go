package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/VerizonDigital/ectoken-go/v3"
)

func main() {
	key := flag.String("key", "", "token key")
	params := flag.String("params", "", "querystring parameters to encrypt")
	token := flag.String("token", "", "token to decrypt")
	decrypt := flag.Bool("decrypt", false, "true")
	verbose := flag.Bool("verbose", false, "verbose mode")

	// parse flags
	flag.Parse()

	if *decrypt == false {
		token, err := v3.Encrypt(*key, *params, *verbose)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(token)
	} else {
		params, err := v3.Decrypt(*key, *token, *verbose)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(params)
	}
}
