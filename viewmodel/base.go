package viewmodel

type Base struct {
	Title string
}

func NewBase() Base {
	return Base{
		Title: "Golang First Dive",
	}
}
