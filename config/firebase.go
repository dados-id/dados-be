package config

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/dados-id/dados-be/exception"
	"google.golang.org/api/option"
)

func NewFireBase() *auth.Client {
	opt := option.WithCredentialsFile("privateKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	exception.FatalIfNeeded(err, "Error initializing app")

	client, err := app.Auth(context.Background())
	exception.FatalIfNeeded(err, "Error getting Auth client")

	return client
}
