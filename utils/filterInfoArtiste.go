package utils

import (
	"fmt"
	n "gpt/structure"
	"strconv"
	"strings"
	"time"
)

func FilterMem(a []string) []string {
	b := []int{}
	for _, v := range a {
		i, _ := strconv.Atoi(v)
		b = append(b, i)
	}
	c := []string{}
	for _, v := range b {
		for _, g := range n.AllData {
			if len(g.Members) == v {
				c = append(c, g.Name)
			}
		}
	}
	return c
}

func FilterFA(from, to string) []string {
	a := []string{}
	date1, e1 := time.Parse("2006-01-02", from)
	date2, e2 := time.Parse("2006-01-02", to)
	fmt.Println("fromFA:", date1)
	fmt.Println("toFA:", date2)
	if e1 != nil || e2 != nil {
		fmt.Println("Time parse problem")
	}
	if date2.Before(date1) {
		date1, date2 = date2, date1
	}
	for _, v := range n.AllData {
		FAdate, e3 := time.Parse("02-01-2006", v.FirstAlbum)
		if e3 != nil {
			fmt.Println("FADate parse problem")
		}
		if date2.After(FAdate) && date1.Before(FAdate) {
			fmt.Println(date1, "<", FAdate, "<", date2)
			a = append(a, v.Name)
		} else {
			fmt.Println("ELSE:")
			fmt.Println(date1, "<", FAdate, "<", date2)
		}
	}
	return a
}

func FilterCD(min, max int) []string {
	if max < min {
		min, max = max, min
	}
	a := []string{}
	for _, v := range n.AllData {
		if v.CreationDate >= min && v.CreationDate <= max {
			a = append(a, v.Name)
		}
	}
	return a
}

func FilterLoc(a string) []string {
	b := []string{}
	for _, artist := range n.AllData {
		for key := range artist.Concerts {
			if SearchWord(strings.ToLower(key), strings.ToLower(a)) {
				b = append(b, artist.Name)
			}
		}
	}
	return b
}
