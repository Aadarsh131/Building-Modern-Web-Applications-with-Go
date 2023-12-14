package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portnumber = ":8080"

func main() {

	session := scs.New()
	session.Lifetime = 24 * time.Hour

	// http.HandleFunc("/", handlers.Home)
	// http.HandleFunc("/About", handlers.About)
	// http.HandleFunc("/Hello", handlers.About)
	// http.HandleFunc("/Add", handlers.Add)
	// http.HandleFunc("/Divide", handlers.Divide)

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
