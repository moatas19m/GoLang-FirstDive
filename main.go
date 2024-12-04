package main

import (
	"GoLang-FirstDive/viewmodel"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		template := templates[requestedFile+".html"]
		context := viewmodel.NewBase()
		if template != nil {
			template.Execute(w, context)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/css/", http.FileServer(http.Dir("./css/")))
	http.ListenAndServe(":9000", nil)
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/header.html", basePath+"/footer.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Cannot open template directory: " + err.Error())
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		panic("Cannot read template directory: " + err.Error())
	}

	for _, file := range files {
		f, err := os.Open(basePath + "/content/" + file.Name())
		if err != nil {
			panic("Cannot open template file: " + file.Name() + err.Error())
		}

		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Cannot read template file: " + file.Name() + err.Error())
		}

		f.Close()

		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Cannot parse template: " + file.Name() + err.Error())
		}
		result[file.Name()] = tmpl
	}
	return result
}
