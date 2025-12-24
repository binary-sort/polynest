package svg

import "testing"

func TestParsePathRectangle(t *testing.T) {
	d := "M10 10 L20 10 L20 20 L10 20 Z"
	polys, err := ParsePath(d)

	if err != nil {
		t.Fatal(err)
	}

	if len(polys) != 1 {
		t.Fatalf("expected 1 polygon")
	}

	if len(polys[0].Points) != 5 {
		t.Fatalf("expected closed polygon")
	}
}
