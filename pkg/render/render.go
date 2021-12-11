package render

import (
	"bytes"
	"fmt"
	"github.com/ahmedkhaeld/go-web/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//var functions a map of functions that can be used in templates e.g. format a date
// some time we will create our own functions and pass them to the template
var functions = template.FuncMap{}

// app pointer will have access to the configuration to access TemplateCache or other AppConfig fields
var app *config.AppConfig

// NewTemplates  set app to the AppConfig when it is called to use the TemplateCache
func NewTemplates(a *config.AppConfig) {
	app = a
}
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	var templateCache map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()

	}

	theTemplate, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("could not get the template from teh templateCache")
	}

	aBuffer := new(bytes.Buffer)

	_ = theTemplate.Execute(aBuffer, nil)

	_, err := aBuffer.WriteTo(w)
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
