package godsc

import (
	"sort"
	"strings"
)

// Bigram holds Unicode code points
type Bigram struct{ r0, r1 rune }

// Bigrams of a string
func Bigrams(str string) map[Bigram]bool {
	bigrams := make(map[Bigram]bool)
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
	bigrams := make(map[Bigram]bool)
	words := strings.Fields(str)

	for _, word := range words {
		for key := range Bigrams(word) {
			bigrams[key] = true
		}
	}

	return bigrams
}

// CompareTwoStrings returns the score of two strings
func CompareTwoStrings(strA, strB string) float64 {
	// Return if strings are identical
	if strA == strB {
		return 1.0
	}

	// Return if strings have no length
	if len(strA) == 0 && len(strB) == 0 {
		return 0.0
	}

	// Return if strings only have 1 char
	if len(strA) == 1 && len(strB) == 1 {
		return 0.0
	}

	bigramsA, bigramsB := bigramsForWords(strA), bigramsForWords(strB)

	var intersection float64

	for a := range bigramsA {
		if bigramsB[a] {
			intersection++
		}
	}

	return 2.0 * intersection / float64(len(bigramsA)+len(bigramsB))
}

// A Match struct is part of Matches
type Match struct {
	Text  string
	Score float64
}

// Matches are returned from FindBestMatch
type Matches struct {
	Candidates []Match
	BestMatch  Match
}

// FindBestMatch returns a struct with a slice of Candidates of type Match
// It also contains a single BestMatch of type Match
func FindBestMatch(str string, candidates []string) Matches {
	scores := []Match{}

	for _, candidate := range candidates {
		scores = append(scores, Match{
			candidate,
			CompareTwoStrings(str, candidate),
		})
	}

	sort.Slice(scores, func(a, b int) bool {
		return scores[a].Score > scores[b].Score
	})

	return Matches{scores, scores[0]}
}
