package main

import (
	"ca-library-app/internal/composites"
	"ca-library-app/internal/config"
	"ca-library-app/pkg/logging"
	"context"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// entry point  точка входа в приложение
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("get initializing")
	cfg := config.GetConfig()

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("mongodb composite initializing")
	mongoDBC, err := composites.NewMongoDBComposite(context.Background(), "", "", "", "", "", "")
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}
	logger.Info("author composite initializing")
	authorComposite, err := composites.NewAuthorComposite(mongoDBC)
	if err != nil {
		logger.Fatal("author composite failed")
	}

	authorComposite.Handler.Register(router)

	logger.Info("book composite initializing")
	bookComposite, err := composites.NewBookComposite(mongoDBC)
	if err != nil {
		logger.Fatal("book composite failed")
	}
	bookComposite.Handler.Register(router)

	// start
}
