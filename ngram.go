package gongram

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// Generate generates Ngrams of given gram size from the input text.
func Generate(text string, gramSize int) ([]string, error) {
	//check if the given gram size is not a positive number.
	if gramSize < 1 {
		return nil, errors.New("please enter a valid gram size")
	}
	text = strings.TrimSpace(text)
	text = sanitize(text)

	re, err := regexp.Compile(`\s+`)
	if err != nil {
		return nil, err
	}
	tokens := re.Split(text, -1)

	// check if the gram size is greater than tokens size
	if len(tokens) < gramSize {
		 return nil, errors.New("gram size is greater than the number of tokens")
	}

	var ngrams []string
	for i := 0 ; i<(len(tokens)-gramSize+1);i++ {
		ngram := strings.Join(tokens[i:i+gramSize], " ")
		ngrams = append(ngrams, ngram)
	}
	return ngrams, nil
}

// sanitize will remove all punctuations from the input string str.
func sanitize(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsPunct(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
