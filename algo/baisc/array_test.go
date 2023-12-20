package basic

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	type args struct {
		n int
	}
	var tests []struct {
		name string
		args args
		want []int
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name string
			args args
			want []int
		}{name: "", args: args{n: i}, want: make([]int, i)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Array(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Array() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayDelete(t *testing.T) {
	type args struct {
		array []int
		index int
	}
	var tests []struct {
		name string
		args args
		want []int
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name string
			args args
			want []int
		}{name: "", args: args{array: make([]int, i), index: i / 2}, want: make([]int, i-1)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayDelete(tt.args.array, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayExpand(t *testing.T) {
	type args struct {
		array []int
		n     int
	}
	var tests []struct {
		name string
		args args
		want []int
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name string
			args args
			want []int
		}{name: "", args: args{array: make([]int, i), n: i / 2}, want: make([]int, i+i/2)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayExpand(tt.args.array, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayExpand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayFind(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	var tests []struct {
		name string
		args args
		want int
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name string
			args args
			want int
		}{name: "", args: args{array: make([]int, i), target: i / 2}, want: i / 2})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayFind(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("ArrayFind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayInsert(t *testing.T) {
	type args struct {
		array []int
		index int
		value int
	}
	var tests []struct {
		name string
		args args
		want []int
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name string
			args args
			want []int
		}{name: "", args: args{array: make([]int, i), index: i / 2, value: i}, want: make([]int, i+1)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayInsert(tt.args.array, tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayRandomAccess(t *testing.T) {
	type args struct {
		array []int
	}
	var tests []struct {
		name string
		args args
		want int
	}
	for i := 0; i < 100; i++ {
		tests = append(tests, struct {
			name string
			args args
			want int
		}{name: "", args: args{array: make([]int, i)}, want: 0})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayRandomAccess(tt.args.array); got != tt.want {
				t.Errorf("ArrayRandomAccess() = %v, want %v", got, tt.want)
			}
		})
	}
}
