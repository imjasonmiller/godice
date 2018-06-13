package godice

import (
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
	// Return if strings have no length.
	if len(strA) == 0 && len(strB) == 0 {
		return 0.0
	}

	// Return if strings only have 1 char.
	if len(strA) == 1 && len(strB) == 1 {
		return 0.0
	}

	// Return if strings are identical.
	if strA == strB {
		return 1.0
	}

	bigramsA, bigramsB := bigramsForWords(strA), bigramsForWords(strB)

	var intersection float64

	for bigramA := range bigramsA {
		// Find bigram of A in bigrams of B
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
func CompareStrings(str string, candidates []string) Matches {
	scores := []Match{}

	for _, candidate := range candidates {
		scores = append(scores, Match{
			candidate,
			CompareString(str, candidate),
		})
	}

	sort.Slice(scores, func(a, b int) bool {
		return scores[a].Score > scores[b].Score
	})

	return Matches{
		BestMatch:  scores[0],
		Candidates: scores,
	}
}
