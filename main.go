package main

import (
	"html/template"
	"os"
)

const tax = 6.75 / 100

type Product struct {
	Name  string
	Price float64
}

const templateString = `
{{- "Item Info:" }}
Name: {{ .Name }}
Price: {{ printf "$%.2f" .Price}}
Price With Tax: {{ calctax .Price | printf "%.2f" }}
`

func main() {
	p := Product{"Pizza", 10}

	funcMap := template.FuncMap{}
	funcMap["calctax"] = func(price float64) float64 {
		return price * (1 + tax)
	}

	t := template.Must(template.New("").Funcs(funcMap).Parse(templateString))
	t.Execute(os.Stdout, p)
}
