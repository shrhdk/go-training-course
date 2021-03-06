package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Marshal it:
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatal(err)
	}

	// Create Decoder:
	buf := bytes.NewReader(data)
	dec := NewDecoder(buf)

	// Decode it:
	for {
		token, err := dec.Token()
		if err == io.EOF {
			return
		}
		fmt.Printf("%s %s\n", reflect.TypeOf(token), token)
	}
}
