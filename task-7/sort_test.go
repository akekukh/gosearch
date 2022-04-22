package task_7

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

const (
	MaxRandLimit = 1_000
)

func TestSort_Ints(t *testing.T) {
	array := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	sort.Ints(array)
	got := array
	want := []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSort_Strings(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"test1", []string{}, []string{}},
		{"test2", nil, nil},
		{"test3", []string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}, []string{"", "%*&^*&^&", "***", "Hello", "bar", "f00", "foo", "foo"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			sort.Strings(input)
			got := input
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSort_Ints(b *testing.B) {
	data := generateRandomIntArray(b.N)
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}

func BenchmarkSort_Strings(b *testing.B) {
	data := generateRandomStringArray(b.N)
	for i := 0; i < b.N; i++ {
		sort.Strings(data)
	}
}

func generateRandomIntArray(arrayLen int) []int {
	var a []int
	for i := 0; i < arrayLen; i++ {
		a = append(a, rand.Intn(MaxRandLimit))
	}
	return a
}

func generateRandomStringArray(arrayLen int) []string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]string, arrayLen)
	for i := range b {
		b[i] = string(letterBytes[rand.Intn(len(letterBytes))])
	}
	return b
}
