package main

import (
	"clickshopUser/awsgo"
	"clickshopUser/db"
	"clickshopUser/models"

	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.AWSinit()

	if !ValidParameters() {
		fmt.Printf("Error sending parameters, it should be 'SecretName")
		err := errors.New("Error sending parameters, it should be 'SecretName")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error reading Secret" + err.Error())
		return event, err
	}

	err = db.SignUp(data)
	return event, err
}

func ValidParameters() bool {
	var takeParameter bool
	_, takeParameter = os.LookupEnv(("Secretname"))
	return takeParameter
}
