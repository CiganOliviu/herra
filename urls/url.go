package urls

import (
	"net/http"
	"html/template"
)

type UserRefference struct {
    FirstName string
    LastName string
}

func Index (Writer http.ResponseWriter, Request *http.Request) {
  user_refference := UserRefference{Title: "Amazing News Aggregator", News: "some news"}
  setTemplate, _ := template.ParseFiles("views/index.html")
  setTemplate.Execute(Writer, user_refference)
}
