package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	s "gpt/structure"
)


func Artistes(res http.ResponseWriter, req *http.Request) {
	temp, er := template.ParseFiles("docs/htmlTemplates/artistes.html")
	if er != nil {
		log.Fatal(er)
		Err(res, req, http.StatusInternalServerError)
		return
	}
	name := req.FormValue("name")
	for _, v := range s.AllData {
		if v.Name == name {
			fmt.Println(v)
			temp.Execute(res, v)
			break
		}
	}
	return
}
