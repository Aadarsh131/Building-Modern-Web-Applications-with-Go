package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Aadarsh131/Building-Modern-Web-Applications-with-Go/pkg/config"
	"github.com/Aadarsh131/Building-Modern-Web-Applications-with-Go/pkg/handlers"
	"github.com/alexedwards/scs/v2"
)

const portnumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd

	//store it in application wide config, so that we can access it from any package
	app.Session = session

	// http.HandleFunc("/", handlers.Home)
	// http.HandleFunc("/About", handlers.About)
	// http.HandleFunc("/Hello", handlers.About)
	// http.HandleFunc("/Add", handlers.Add)
	// http.HandleFunc("/Divide", handlers.Divide)


	app.Name = "Aadarsh" 
	app.IsRich = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println("Starting the server on port", portnumber)

	// serv := &http.Server{ //http.Server without the & would also work, but we are using the ref
	// 	Addr:    portnumber,
	// 	Handler: routes(),
	// }
	// err := serv.ListenAndServe()

	//OR
	err := http.ListenAndServe(portnumber, routes())
	if err != nil {
		log.Fatal(err)
	}
}
