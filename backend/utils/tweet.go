package utils

import (
	"regexp"
	"strings"
	"unicode"
)

var frequentWords = map[string]bool{
	"but":   true,
	"the":   true,
	"in":    true,
	"a":     true,
	"this":  true,
	"these": true,
	"of":    true,
	"and":   true,
	"not":   true,
	"i":     true,
	"just":  true,
	"was":   true,
	"with":  true,
	"or":    true,
	"were":  true,
	"he":    true,
	"she":   true,
	"did":   true,
	"are":   true,
	"is":    true,
	"am":    true,
	"we":    true,
	"to":    true,
	"be":    true,
	"you":   true,
	"if":    true,
}

func TokenizeTweet(tweet string) []string {
	lt := strings.ToLower(tweet)
	r := strings.NewReplacer("\"", "", "…", "")
	lt = r.Replace(lt)
	words := strings.FieldsFunc(lt, ff)
	return words
}

func CleanTweet(tweet string) string {
	urlRe := regexp.MustCompile("https?://[^\\s]+")
	tweet = urlRe.ReplaceAllLiteralString(tweet, "")

	rtRe := regexp.MustCompile("RT @[^\\s]+")
	tweet = rtRe.ReplaceAllLiteralString(tweet, "")

	return tweet
}

func DeleteFrequentWords(words []string) []string {
	newWords := make([]string, 0, len(words))

	for _, word := range words {
		if _, ok := frequentWords[word]; !ok {
			newWords = append(newWords, word)
		}
	}

	return newWords
}

func ff(c rune) bool {
	return unicode.IsSpace(c) || (unicode.IsPunct(c) && c != '\'')
}
