package main

import (
	"fmt"
	"log"
	"github.com/DaminAlaskarovTC/bookings/pkg/config"
	"github.com/DaminAlaskarovTC/bookings/pkg/handlers"
	"github.com/DaminAlaskarovTC/bookings/pkg/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
// 
var session *scs.SessionManager

func main() {
	


	//change this to true when in productionn
	app.InProduction = false


	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	//checked of de connction https is in plaats van http in production moet dit true zijn maar wij gebruiken nu localhost en die is http
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//weg gecomment voor lecture 36
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	// 	n,err := fmt.Fprintf(w ,"Hello sdsdf world!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes: %d",n))
	// })

	fmt.Println(fmt.Sprintf("starting application on port: %s", portNumber))
	// weg voor lec36
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}

// iets met een mod file
