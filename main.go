package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aditmeno/registry-credential-helper/ecr"
	"github.com/aditmeno/registry-credential-helper/gcr"
	registryInterface "github.com/aditmeno/registry-credential-helper/interface"
)

func main() {
	if os.Getenv("AWS_WEB_IDENTITY_TOKEN_FILE") != "" {
		ecrHelper := ecr.GetECRCredentialHelper()
		RegistryMain(ecrHelper)
	} else if os.Getenv("GOOGLE_CLOuD_PROVIDER") != "" {
		gcrHelper := gcr.GetGCRCredentialHelper()
		RegistryMain(gcrHelper)
	} else {
		log.Fatal("Couldn't determine Cloud Provider")
	}
}

func RegistryMain(rh registryInterface.RegistryHelper) {
	rh.Login()
	token := rh.GetToken()
	fmt.Println(token)
}
