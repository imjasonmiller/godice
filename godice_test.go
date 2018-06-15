package godice

import (
	"reflect"
	"testing"
)

func TestCompareString(t *testing.T) {
	tests := map[struct{ strA, strB string }]float64{
		{"George Orwell", "George Orwell"}: 1.0,
		{"Carl Sagan", "cArL sAgAn"}:       1.0,
		{"Aldous Huxley", "Isaac Asimov"}:  0.0,
		{"A", "A"}:                         1.0,
		{"", ""}:                           0.0,
	}

	for test, val := range tests {
		res := CompareString(test.strA, test.strB)

		if res != val {
			t.Errorf("inputted: %+v, expected %f, got: %v", test, val, res)
		}
	}
}

func TestBigrams(t *testing.T) {
	tests := map[string]map[Bigram]bool{
		"Mirabela Calin": map[Bigram]bool{
			Bigram{77, 105}:  true,
			Bigram{97, 32}:   true,
			Bigram{32, 67}:   true,
			Bigram{108, 97}:  true,
			Bigram{67, 97}:   true,
			Bigram{97, 108}:  true,
			Bigram{105, 114}: true,
			Bigram{114, 97}:  true,
			Bigram{97, 98}:   true,
			Bigram{98, 101}:  true,
			Bigram{101, 108}: true,
			Bigram{108, 105}: true,
			Bigram{105, 110}: true,
		},
		"გია ჩხეიძე": map[Bigram]bool{
			Bigram{4334, 4308}: true,
			Bigram{4308, 4312}: true,
			Bigram{4306, 4312}: true,
			Bigram{4312, 4304}: true,
			Bigram{4304, 32}:   true,
			Bigram{32, 4329}:   true,
			Bigram{4329, 4334}: true,
			Bigram{4312, 4331}: true,
			Bigram{4331, 4308}: true,
		},
		"Αλέξανδρος": map[Bigram]bool{
			Bigram{941, 958}: true,
			Bigram{958, 945}: true,
			Bigram{961, 959}: true,
			Bigram{913, 955}: true,
			Bigram{955, 941}: true,
			Bigram{945, 957}: true,
			Bigram{957, 948}: true,
			Bigram{948, 961}: true,
			Bigram{959, 962}: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			bigrams := Bigrams(name)
			if !reflect.DeepEqual(test, bigrams) {
				t.Errorf("inputted: %s, expected %v, got: %v", name, test, bigrams)
			}
		})
	}
}
