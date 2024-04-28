package bandname_test

import (
	"bandname"
	"testing"
)

func TestCombiningCityAndPetName(t *testing.T) {
	t.Parallel()
	want := "Los Angeles Kale"
	got := bandname.CombiningCityAndPetName("Los Angeles", "Kale")
	if want != got {
		t.Errorf("Want %s, got %s.", want, got)
	}
}

