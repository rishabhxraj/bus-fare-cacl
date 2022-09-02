package fare

import "testing"

func TestCalculateFare(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{"Distance-30", 30, 47},
		{"Distance-10", 10, 17},
		{"Distance-26", 26, 43},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateFare(tt.args); got != tt.want {
				t.Errorf("CalculateFare() = %v, want %v", got, tt.want)
			}
		})
	}
}
