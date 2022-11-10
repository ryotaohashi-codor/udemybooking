package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

var portNumber = ":8080"

func main(){
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("cannnot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))

	srv := &http.Server {
		Addr: "127.0.0.1"+portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}