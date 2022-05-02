package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"github.com/DaminAlaskarovTC/bookings/pkg/config"
	"github.com/DaminAlaskarovTC/bookings/pkg/models"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

//NewTemplates sets teh config for the new package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData( td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template using html/template package
func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// tc = app.TemplateCache

	//laat zien of de templates kunne laden in de cache
	// _, err := RenderTemplateTest(w)
	// if err != nil {
	// 	fmt.Println("Error getting template cache: ", err)
	// }

	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// als er geen html is(dus bijvoorbeeld idsgagd.page.html) dan krijgt ok false als er wel iets is gevonden dan krijgt t die waarde
	t, ok := tc[html]
	// als !true (ok is bool)
	if !ok {
		log.Fatal("could not get template from template cache ( ok) ")
	}

	buf := new(bytes.Buffer)


	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser : ", err)
		return
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/" + html)

	// err := parsedTemplate.Execute(w, nil)
}

// CreateTemplateCache creates a template cache as map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("current page: ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
