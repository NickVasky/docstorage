package app

import (
	"context"
	"log"
	"net/http"

	"github.com/NickVasky/docstorage/internal/closer"
	"github.com/NickVasky/docstorage/internal/codegen/apicodegen"
	"github.com/NickVasky/docstorage/internal/config"
)

type App struct {
	httpServer      *http.Server
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDependencies(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runHttpServer()
}

func (a *App) initDependencies(ctx context.Context) error {
	initFuncs := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHttpServer,
	}

	for _, initFunc := range initFuncs {
		err := initFunc(ctx)
		if err != nil {
			return err
		}
	}

	return nil

}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initHttpServer(ctx context.Context) error {

	docAPIController := apicodegen.NewDocumentsAPIController(a.serviceProvider.DocAPIService(ctx))
	authAPIController := apicodegen.NewAuthAPIController(a.serviceProvider.AuthAPIService())

	router := a.newRouter(docAPIController, authAPIController)

	httpServer := &http.Server{
		Addr:    a.serviceProvider.HttpConfig().Address(),
		Handler: router,
	}

	a.httpServer = httpServer

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) runHttpServer() error {
	log.Printf("Server is running on %s", a.serviceProvider.HttpConfig().Address())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
