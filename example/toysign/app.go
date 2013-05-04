package main

import (
	"fmt"
	"github.com/bufio/toys/locale"
	"github.com/bufio/toys/view"
	"labix.org/v2/mgo"
	"net/http"
)

func main() {
	// Configuration variable
	var (
		// cnnStr the connection string to MongoDB
		cnnStr = "localhost"
		// langRoot the path to language folder in file system
		langRoot = "language"
		// langDefaultSet the default language set
		langDefaultSet = "en"
		// tmplDefaultSet the path to template folder in files system
		tmplRoot = "template"
		// tmplDefaultSet the default template set
		tmplDefaultSet = "default"
		// rsrcRoot the path to static folder in file system
		rsrcRoot = "statics"
		// rsrcPrefix the URL path for static file server
		rsrcPrefix = "/statics/"
		//toysignPath the URL path for toysign
		toysignPath = "/user/"
	)

	//database session
	dbsess, err := mgo.Dial(cnnStr)
	if err != nil {
		panic(err)
	}
	defer dbsess.Close()

	//multi language support
	lang := locale.NewLang(langRoot)
	if err := lang.Parse(langDefaultSet); err != nil {
		fmt.Println(err.Error())
	}

	//template for cms
	tmpl := view.NewView(tmplRoot)
	tmpl.SetLang(lang)
	tmpl.HandleResource(rsrcPrefix, rsrcRoot)
	if err := tmpl.Parse(tmplDefaultSet); err != nil {
		fmt.Println(err.Error())
	}

	http.Handle(toysignPath, Handler(toysignPath, dbsess, tmpl))
	http.ListenAndServe("localhost:8080", nil)
}