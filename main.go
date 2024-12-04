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

func (p Product) PriceWithTax() float64 {
	return p.Price * (1 + tax)
}

const templateString = `
{{- "Item Info:" }}
Name: {{ .Name }}
Price: {{ printf "$%.2f" .Price}}
Price With Tax: {{ .PriceWithTax | printf "%.2f" }}
`

func main() {
	p := Product{"Pizza", 10}

	t := template.Must(template.New("templateString").Parse(templateString))
	t.Execute(os.Stdout, p)
}
