package urls

import (
	"net/http"
	"html/template"
)

type NewsAggPage struct {
    Title string
    News string
}

func Index (w http.ResponseWriter, r *http.Request) {
  p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
  t, _ := template.ParseFiles("views/index.html")
  t.Execute(w, p)
}
