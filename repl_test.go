package main


import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "this    IS  A teST    of  	TE-sTIng",
			expected: []string{"this", "is", "a", "test", "of", "te-sting"},
		},
		{
			input:    "what  if    we    use numbers 223454",
			expected: []string{"what", "if", "we", "use", "numbers", "223454"},
		},
		{
			input:    "mewtwo charizard",
			expected: []string{"mewtwo", "charizard"},
		},
		{
			input:    "Charizard	 Gengar Arcanine   Bulbasaur Blaziken  Umbreon Lucario Gardevoir 	Eevee Dragonite",
			expected: []string{"charizard", "gengar", "arcanine", "bulbasaur", "blaziken", "umbreon", "lucario", "gardevoir", "eevee", "dragonite"},
		},
		// add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanInput(%q) produced %v, expected %v", c.input, actual, c.expected)
			}

		}
	}
}
