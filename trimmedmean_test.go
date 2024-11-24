package trimmedmean

import "testing"

func TestTrimmedMean(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mean, err := TrimmedMean(data, 0.1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := 5.5 // Mean of [2, 3, 4, 5, 6, 7, 8, 9]
	if mean != expected {
		t.Errorf("expected %v, got %v", expected, mean)
	}
}
