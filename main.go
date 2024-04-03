package main

import (
	"fmt"
	p "gpt/handlers"
	r "gpt/structure"
	"net/http"
	"os"
)

func main() {
	FileServer := http.FileServer(http.Dir("docs"))
	http.Handle("/docs/", http.StripPrefix("/docs/", FileServer))
	r.AllData = p.GatherDataUp("https://groupietrackers.herokuapp.com/api/artists")
	if r.AllData == nil {
		fmt.Println("Failed to gather Data from API")
		os.Exit(1)
	}
	http.HandleFunc("/", p.MainPage)
	http.HandleFunc("/artiste", p.Artistes)
	http.HandleFunc("/search", p.Search)

	fmt.Print("Server running on ")
	println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
