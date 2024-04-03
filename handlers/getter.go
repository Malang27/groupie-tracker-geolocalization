package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	s "gpt/structure"
)

func GatherDataUp(link string) []s.ArtistData {
	data1 := GetData(link)
	Artists := []s.ArtistData{}
	e := json.Unmarshal(data1, &Artists)
	if e != nil {
		log.Fatal(e)
		return nil
	}
	for i := 0; i < len(Artists); i++ {
		r := s.Relation{}
		json.Unmarshal(GetData(Artists[i].Relation), &r)
		Artists[i].Concerts = r.Concerts
	}
	return Artists
}
func GetData(link string) []byte {
	data1, e1 := http.Get(link)
	if e1 != nil {
		log.Fatal(e1)
		return nil
	}
	data2, e2 := ioutil.ReadAll(data1.Body)
	if e2 != nil {
		log.Fatal(e2)
		return nil
	}
	return data2
}
