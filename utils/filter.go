package utils

import (
	"fmt"
	m "gpt/structure"
	"net/http"
	"strconv"
	"strings"
)

func Filter(req *http.Request) []string {
	req.ParseForm()

	fromCD, e1 := strconv.Atoi(req.FormValue("fromCD"))
	toCD, e2 := strconv.Atoi(req.FormValue("toCD"))
	if e1 != nil {
		fromCD = 1957
	}
	if e2 != nil {
		toCD = 2020
	}
	fmt.Println("fromCD:", fromCD)
	fmt.Println("toCD:", toCD)

	f1 := FilterCD(fromCD, toCD)
	fmt.Println("f1:", f1)

	memNum := req.Form["member"]
	f2 := FilterMem(memNum)
	fmt.Println("memNum:", memNum)
	if len(memNum) == 0 {
		f2 = f1
	}
	fmt.Println("f2:", f2)

	fromFA := req.FormValue("fromFA")
	toFA := req.FormValue("toFA")
	if fromFA == "" {
		fromFA = "1957-01-01"
	}
	if toFA == "" {
		toFA = "2025-01-01"
	}
	f3 := FilterFA(fromFA, toFA)
	fmt.Println("f3:", f3)

	loc := req.FormValue("location")
	fmt.Println("loc:", loc)
	f4 := FilterLoc(loc)
	loc = strings.TrimSpace(loc)
	if loc == "" {
		f4 = f1
	}
	fmt.Println("f4:", f4)

	result := []string{}
	for _, v := range m.AllData {
		if HasString(f1, v.Name) && HasString(f2, v.Name) && HasString(f3, v.Name) && HasString(f4, v.Name) {
			result = append(result, v.Name)
		}
	}
	return result
}
func SearchWord(s string, toFind string) bool {
	for i := 0; i <= len(s)-len(toFind); i++ {
		if toFind == s[i:i+len(toFind)] {
			return true
		}
	}
	return false
}

func HasString(ar []string, a string) bool {
	for _, v := range ar {
		if v == a {
			return true
		}
	}
	return false
}
