package render

import (
	"fmt"
	"net/http"
)

//RenderTemplate renders the templates
func RenderTemplate(w http.ResponseWriter, tmpl string) { //lowercase named functions can only be accesible within the same package (private field), hence making is UpperCase to make it accesible to any pacakage
	// parsedTemplatePointer, _ := template.ParseFiles("./templates/" + tmpl)
	// err := parsedTemplatePointer.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error parsing template: ", err)
	// 	return
	// }
	fmt.Fprintln(w,"from RnderTemplateFunc")
}

