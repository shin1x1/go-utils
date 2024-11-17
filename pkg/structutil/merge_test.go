package structutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeStruct(t *testing.T) {
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
		a any
		b map[string]any
	}
	type testCase struct {
		name string
		args args
		want any
	}
	tests := []testCase{
		{
			name: "Merge all fields",
			args: args{
				a: User{
					Name:    "Foo",
					Email:   "a@example.com",
					Age:     10,
					Friends: []string{"Bar", "Boz"},
					Scores:  map[string]float64{"math": 85, "science": 90},
					Address: &Address{
						City:    "City1",
						ZipCode: "1234567",
					},
					IsActive: true,
				},
				b: map[string]any{
					"Name":     "Foo1",
					"Email":    "a1@example.com",
					"Age":      20,
					"Friends":  []string{"Mike"},
					"Scores":   map[string]float64{"aaa": 20},
					"Address":  &Address{City: "City11", ZipCode: "9999999"},
					"IsActive": false,
				},
			},
			want: User{
				Name:    "Foo1",
				Email:   "a1@example.com",
				Age:     20,
				Friends: []string{"Mike"},
				Scores:  map[string]float64{"aaa": 20},
				Address: &Address{
					City:    "City11",
					ZipCode: "9999999",
				},
				IsActive: false,
			},
		},
		{
			name: "Merge partial fields",
			args: args{
				a: Address{
					City:    "City1",
					ZipCode: "1234567",
				},
				b: map[string]any{
					"ZipCode": "11111",
				},
			},
			want: Address{
				City:    "City1",
				ZipCode: "11111",
			},
		},
		{
			name: "Struct pointer",
			args: args{
				a: &Address{
					City:    "City1",
					ZipCode: "1234567",
				},
				b: map[string]any{
					"ZipCode": "11111",
				},
			},
			want: &Address{
				City:    "City1",
				ZipCode: "11111",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MergeStruct(tt.args.a, tt.args.b))
		})
	}
}
