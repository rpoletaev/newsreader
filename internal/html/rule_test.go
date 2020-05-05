package html

import (
	"os"
	"testing"
)

func TestRBC(t *testing.T) {
	f, err := os.Open("./fixtures/rbc.html")
	if err != nil {
		t.Fatal("on open file:", err)
	}

	r := Rule{
		Root: ".l-row span[itemprop='itemListElement']",
	}

}
