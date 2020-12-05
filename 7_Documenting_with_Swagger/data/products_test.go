package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Tea",
		Price: 2,
		SKU: "abc-abc-avvv",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
