package ecr

import (
	b64 "encoding/base64"
	"log"
	"os"

	awsSession "github.com/aws/aws-sdk-go/aws/session"
	awsECR "github.com/aws/aws-sdk-go/service/ecr"

	_ "github.com/aditmeno/registry-credential-helper/interface"
)

type ECRCredentialHelper struct {
	Token   string
	Session interface{}
}

func GetECRCredentialHelper() *ECRCredentialHelper {
	return &ECRCredentialHelper{}
}

func (credHelper *ECRCredentialHelper) Login() {
	session := login()
	credHelper.Session = session
}

func (credHelper *ECRCredentialHelper) GetToken() string {
	awsRegistry := os.Getenv("AWS_ECR_REGISTRY")
	if awsRegistry == "" {
		log.Fatal("AWS Registry missing, aborting!!")
	}
	credHelper.Token = getToken(awsRegistry, credHelper.Session.(*awsECR.ECR))
	return credHelper.Token
}

func login() *awsECR.ECR {
	awsSession := awsSession.Must(awsSession.NewSession())
	ecrSession := awsECR.New(awsSession)
	return ecrSession
}

func getToken(registryID string, session *awsECR.ECR) string {
	registry := awsECR.GetAuthorizationTokenInput{
		RegistryIds: []*string{
			&registryID,
		},
	}
	tokenOutput, err := session.GetAuthorizationToken(&registry)
	if err != nil {
		log.Fatalf("Encountered an error: %s", err.Error())
	}
	payload := tokenOutput.AuthorizationData[0].AuthorizationToken
	decodedString, _ := b64.StdEncoding.DecodeString(*payload)
	extractedToken := string(decodedString)[4:]
	return extractedToken
}
