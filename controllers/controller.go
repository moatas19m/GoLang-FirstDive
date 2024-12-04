package controllers

import (
	"html/template"
)

var (
	homeController home
)

func StartUp(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	homeController.registerRoutes()
}
