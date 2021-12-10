package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	parsedTemplate, err := template.ParseFiles("../../templates/" + tmpl)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
	parsedTemplate.Execute(w, nil)

}
