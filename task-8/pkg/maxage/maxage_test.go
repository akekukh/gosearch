package maxage

import "testing"

func TestMaxAge(t *testing.T) {
	tests := []struct {
		name  string
		users []User
		want  int
	}{
		{
			name: "Employee",
			users: []User{
				Employee{age: 10},
				Employee{age: 30},
			},
			want: 30,
		},
		{
			name: "Customer",
			users: []User{
				Customer{age: 10},
				Customer{age: 30},
			},
			want: 30,
		},
		{
			name: "Customer and Employee",
			users: []User{
				Customer{age: 10},
				Employee{age: 30},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.users...); got != tt.want {
				t.Errorf("MaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
