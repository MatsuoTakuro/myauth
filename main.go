package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/idtoken"
)

var (
	GoogleCliID = os.Getenv("GOOGLE_CLIENT_ID")
	IDToken     = os.Getenv("ID_TOKEN")
)

func main() {

	tkValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		fmt.Println("fails to create id_token validator: ", err)
		return
	}

	p, err := tkValidator.Validate(context.Background(), IDToken, GoogleCliID)
	if err != nil {
		fmt.Println("validate err: ", err)
		return
	}

	fmt.Println(p.Claims["name"])

}
