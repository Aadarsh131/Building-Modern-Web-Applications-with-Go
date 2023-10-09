package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

const portnumber = ":8080"

func Hello(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello")
	fmt.Println(fmt.Sprintf("Bytes Written: %d", n))
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 4)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("Added result: %d", sum))
}

// addValues adds two values and returns the result (Convention is to use the function name at the start)
func addValues(x, y int) int { //a function name with first character as Lower case means that the function is only accessible within the package it was created in (same as private and public fields in OOPS langauge)
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	val, err := divideValues(200.34, 0)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("%d", err))
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f = %f", 200.34, 234.2, val))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero!")
		return 0, err
	}
	return x / y, nil
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplatePointer, _ := template.ParseFiles("./templates/"+tmpl)
	err:=parsedTemplatePointer.Execute(w,nil)
	if(err != nil){
		fmt.Println("error parsing template: ", err)
	}
	return
}
func main() {
	fmt.Println("Hello")

	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)
	http.HandleFunc("/Hello", Hello)
	http.HandleFunc("/Add", Add)
	http.HandleFunc("/Divide", Divide)

	fmt.Println("Starting the server on port", portnumber)
	_ = http.ListenAndServe(portnumber, nil)
}
