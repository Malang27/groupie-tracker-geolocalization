package handlers

import (
	"net/http"
	"text/template"
	s "gpt/structure"
)

func Search(res http.ResponseWriter, req *http.Request) {
	temp, e1 := template.ParseFiles("docs/htmlTemplates/search.html")
	if e1 != nil {
		Err(res, req, http.StatusInternalServerError)
		return
	}
	temp.Execute(res, s.AllData)
}
