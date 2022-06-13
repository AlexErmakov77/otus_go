package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

type dictWord struct {
	word  string
	count int
}

var word string
var reg = regexp.MustCompile(`[^\r\n\t\f\v\s]+`)

func Top10(str string) []string {
	if len(str) == 0 {
		return nil
	}
	words := make(map[string]int)

	res := reg.FindAllString(str, -1)
	for i := range res {
		word = string(res[i])
		words[word]++
	}
	top := make([]dictWord, 0, len(words))
	for w, c := range words {
		top = append(top, dictWord{w, c})
	}
	sort.Slice(top, func(i int, j int) bool {
		if top[i].count == top[j].count {
			return top[i].word < top[j].word
		}
		return top[i].count > top[j].count
	})
	lenList := len(top)
	if lenList > 10 {
		lenList = 10
	}
	topSlice := make([]string, 0, lenList)
	for i, v := range top {
		if i < lenList {
			topSlice = append(topSlice, v.word)
		}
	}
	return topSlice
}
