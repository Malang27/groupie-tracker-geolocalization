package handlers

import (
	s "gpt/structure"
	u "gpt/utils"
	"html/template"
	"net/http"
)

func MainPage(res http.ResponseWriter, req *http.Request) {
	temp, er := template.ParseFiles("docs/htmlTemplates/index.html")
	if er != nil {
		Err(res, req, http.StatusInternalServerError)
		return
	}
	if req.URL.Path != "/" {
		Err(res, req, http.StatusNotFound)
		return
	}
	f := u.Filter(req)
	result := []s.ArtistData{}
	for _, v := range s.AllData {
		if u.HasString(f, v.Name) {
			result = append(result, v)
		}
	}
	temp.Execute(res, result)
}
