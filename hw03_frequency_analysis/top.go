package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

type Entry struct {
	Key   string
	Value int
}

func Top10(text string) []string {
	if text == "" {
		return []string{}
	}
	strs := strings.Fields(text)

	counter := make(map[string]int)
	entries := []Entry{}

	for i := 0; i < len(strs); i++ {
		convertedStr, okConv := SimplifyString(strs[i])
		if !okConv {
			continue
		}
		_, ok := counter[convertedStr]
		if !ok {
			counter[convertedStr] = 0
		} else {
			counter[convertedStr]++
		}
	}

	for k, v := range counter {
		entries = append(entries, Entry{k, v})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Value == entries[j].Value {
			return entries[i].Key < entries[j].Key
		}
		return entries[i].Value > entries[j].Value
	})

	var result []string
	for i := 0; i < 10 && i < len(entries); i++ {
		result = append(result, entries[i].Key)
	}
	return result
}

func SimplifyString(text string) (string, bool) {
	str := strings.ToLower(text)
	runes := []rune(str)

	if len(runes) > 0 && unicode.IsPunct(runes[0]) {
		runes = runes[1:]
	}
	if len(runes) > 0 && unicode.IsPunct(runes[len(runes)-1]) {
		runes = runes[:len(runes)-1]
	}

	if len(runes) == 0 {
		return "", false
	}
	return string(runes), true
}
