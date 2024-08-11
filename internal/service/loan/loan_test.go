package loan

import "testing"

// Test_calculateAmountEachTerm tests the calculateAmountEachTerm function
func Test_calculateAmountEachTerm(t *testing.T) {
	service := NewService(nil)
	type args struct {
		amount float64
		term   uint64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Test case 1",
			args: args{
				amount: 1000,
				term:   1,
			},
			want: []float64{1000.0},
		},
		{
			name: "Test case 2",
			args: args{
				amount: 1000,
				term:   2,
			},
			want: []float64{500.0, 500.0},
		},
		{
			name: "Test case 3",
			args: args{
				amount: 1000,
				term:   3,
			},
			want: []float64{333.33, 333.33, 333.34},
		},
		{
			name: "Test case 4",
			args: args{
				amount: 0,
				term:   3,
			},
			want: []float64{},
		},
		{
			name: "Test case 5",
			args: args{
				amount: 1000,
				term:   0,
			},
			want: []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := service.calculateAmountEachTerm(tt.args.amount, int(tt.args.term))
			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("calculateAmountEachTerm() = %v, want %v", got[i], tt.want[i])
					t.Fail()
				}
			}
		})
	}
}
