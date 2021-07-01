package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`[,.!?:;"']`)

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}
	var top int
	finalSlice := make([]string, 0)
	type pairs struct {
		Word  string
		Count int
	}
	collectedStringsSlice := make([]pairs, 0)
	allStrings := make([]string, 0)
	collectedStrings := make(map[string]int)
	splittedStrings := strings.Fields(text)
	for _, str := range splittedStrings {
		splittedStr := re.Split(str, -1)
		if len(splittedStr) > 1 {
			for _, i := range splittedStr {
				if i != "" && i != "-" {
					allStrings = append(allStrings, strings.ToLower(i))
				}
			}
		} else if splittedStr[0] != "-" {
			allStrings = append(allStrings, strings.ToLower(splittedStr[0]))
		}
	}
	for _, str := range allStrings {
		collectedStrings[str]++
	}
	for key, value := range collectedStrings {
		collectedStringsSlice = append(collectedStringsSlice, pairs{Word: key, Count: value})
	}
	sort.Slice(collectedStringsSlice, func(i, j int) bool {
		if collectedStringsSlice[i].Count == collectedStringsSlice[j].Count {
			return collectedStringsSlice[i].Word < collectedStringsSlice[j].Word
		}
		return collectedStringsSlice[i].Count > collectedStringsSlice[j].Count
	})
	if len(collectedStringsSlice) > 10 {
		top = 10
	} else {
		top = len(collectedStringsSlice)
	}
	for _, v := range collectedStringsSlice[:top] {
		finalSlice = append(finalSlice, v.Word)
	}
	return finalSlice
}
