package main

import (	
	"log/slog"
	"context"
	"net/http"
	"strings"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"os"
	"errors"
)

func initFirebaseAuth(ctx context.Context, logger *slog.Logger) *auth.Client {

	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_SERVICE_ACCOUNT_JSON")))

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}

	return client
}

func (app *application) verifyIDToken(ctx context.Context, r *http.Request) (*auth.Token, error) {

	idToken := r.Header.Get("Authorization")
    splitToken := strings.Split(idToken, "Bearer ")
    if len(splitToken) < 2 {
    	err := errors.New("authorization token not exists")
    	app.logger.Error(err.Error())
    	return nil, err
    }
    idToken = splitToken[1]

	token, err := app.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		app.logger.Error(err.Error())
		return nil, err
	}

	app.logger.Info("IDToken verification success", "Verified ID token", token)

	return token, nil
}
