package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//var functions a map of functions that can be used in templates e.g. format a date
// some time we will create our own functions and pass them to the template
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	theTemplate, ok := templateCache[tmpl]
	if !ok {
		log.Fatal(err)
	}

	aBuffer := new(bytes.Buffer)

	_ = theTemplate.Execute(aBuffer, nil)

	_, err = aBuffer.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing to the browser", err)
	}

}

// CreateTemplateCache return a map that has the parsed templates include the layouts
func CreateTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	//filepath.Glob get the location of template pages.
	pagesPath, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	// for loop extract the page name for the pages' path.
	for _, page := range pagesPath {
		pageName := filepath.Base(page)

		templateSet, err := template.New(pageName).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		// check template matches any layouts
		matches, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}
		cache[pageName] = templateSet
	}
	return cache, nil
}
