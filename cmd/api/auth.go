package main

import (	
	"log/slog"
	"context"
	firebase "firebase.google.com/go/v4"
)

const (
	firebaseProjectId = "books-library-6afdb"
)

func initFirebaseAuth(ctx context.Context, logger *slog.Logger) *auth.Client {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		panic(err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}

	return client
}

func (app *application) verifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {

	token, err := app.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		app.logger.Error(err.Error())
		return nil, err
	}

	app.logger.Info("IDToken verification success", "Verified ID token", token)

	return token, nil
}
