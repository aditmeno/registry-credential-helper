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
	if os.Getenv("AWS_CLOUD_PROVIDER") != "" || os.Getenv("AWS_ACCESS_KEY_ID") != "" || os.Getenv("AWS_SDK_LOAD_CONFIG") != "" {
		ecrHelper := ecr.GetECRCredentialHelper()
		RegistryMain(ecrHelper)
	} else if os.Getenv("GOOGLE_CLOUD_PROVIDER") != "" {
		gcrHelper := gcr.GetGCRCredentialHelper()
		RegistryMain(gcrHelper)
	} else {
		log.Fatal("Couldn't determine Cloud Provider!")
	}
}

func RegistryMain(rh registryInterface.RegistryHelper) {
	rh.Login()
	token := rh.GetToken()
	fmt.Println(token)
}
