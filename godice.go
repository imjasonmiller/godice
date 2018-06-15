package godice

import (
	"errors"
	"sort"
	"strings"
)

// Bigram holds the UTF-8 code points, known as runes.
type Bigram struct{ r0, r1 rune }

// Bigrams returns a map with each bigram initialized as true.
func Bigrams(str string) map[Bigram]bool {
	bigrams := map[Bigram]bool{}
	last := rune(0)

	for index, char := range str {
		if index > 0 {
			bigrams[Bigram{last, char}] = true
		}
		last = char
	}

	return bigrams
}

// Bigrams for multiple words, split by whitespace
func bigramsForWords(str string) map[Bigram]bool {
	bigrams := map[Bigram]bool{}
	words := strings.Fields(str)

	for _, word := range words {
		for key := range Bigrams(word) {
			bigrams[key] = true
		}
	}

	return bigrams
}

// CompareString returns the score of two strings
func CompareString(strA, strB string) float64 {
	// Return if either string has no length.
	if len(strA) == 0 || len(strB) == 0 {
		return 0.0
	}

	// Return if strings are identical.
	if strA == strB {
		return 1.0
	}

	bigramsA := bigramsForWords(strings.ToLower(strA))
	bigramsB := bigramsForWords(strings.ToLower(strB))

	var intersection float64

	// Find bigram of A in bigrams of B.
	for bigramA := range bigramsA {
		if bigramsB[bigramA] {
			intersection++
		}
	}

	return 2.0 * intersection / float64(len(bigramsA)+len(bigramsB))
}

// Match holds the original text and its final score.
type Match struct {
	Text  string
	Score float64
}

// Matches contains a best match and all candidates, sorted by score.
type Matches struct {
	BestMatch  Match
	Candidates []Match
}

// CompareStrings handles multiple strings
func CompareStrings(str string, candidates []string) (Matches, error) {
	scores := []Match{}

	if len(candidates) < 1 {
		return Matches{}, errors.New("slice for CompareStrings must contain at least one element")
	}

	for _, candidate := range candidates {
		scores = append(scores, Match{
			Text:  candidate,
			Score: CompareString(str, candidate),
		})
	}

	sort.Slice(scores, func(a, b int) bool {
		return scores[a].Score > scores[b].Score
	})

	return Matches{
		BestMatch:  scores[0],
		Candidates: scores,
	}, nil
}
