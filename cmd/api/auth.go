package main

import (	
	"log/slog"
	"context"
	"net/http"
	"strings"
	firebase "firebase.google.com/go/v4"
)

const (
	firebaseProjectId = "books-library-6afdb"
)

func initFirebaseAuth(ctx context.Context, logger *slog.Logger) *auth.Client {
	
	var cfg firebase.Config
	cfg.ProjectID = firebaseProjectId

	app, err := firebase.NewApp(ctx, cfg)
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
    idToken = splitToken[1]

	token, err := app.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		app.logger.Error(err.Error())
		return nil, err
	}

	app.logger.Info("IDToken verification success", "Verified ID token", token)

	return token, nil
}
