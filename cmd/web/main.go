package main

import (
	"fmt"
	"github.com/ahmedkhaeld/go-web/pkg/config"
	"github.com/ahmedkhaeld/go-web/pkg/handlers"
	"github.com/ahmedkhaeld/go-web/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// app entry point to access AppConfig
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// assign the cache to the app field
	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting the application %v", portNumber))

	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
