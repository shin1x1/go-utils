package assert

import "testing"

func TestStructFields_Success(t *testing.T) {
	type Address struct {
		City    string
		ZipCode string
	}
	type User struct {
		Name     string
		Email    string
		Age      int
		Friends  []string
		Scores   map[string]float64
		Address  *Address
		IsActive bool
	}

	type args struct {
		actual   any
		expected map[string]any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				actual: Address{
					City:    "city",
					ZipCode: "123",
				},
				expected: map[string]any{
					"City":    "city",
					"ZipCode": "123",
				},
			},
		},
		{
			name: "test",
			args: args{
				actual: Address{
					City:    "city",
					ZipCode: "123",
				},
				expected: map[string]any{
					"ZipCode": "123",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EqualStructFields(t, tt.args.actual, tt.args.expected)
		})
	}
}
