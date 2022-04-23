package older

import (
	"testing"
)

func TestUser(t *testing.T) {
	tests := []struct {
		name  string
		users []interface{}
		want  interface{}
	}{
		{
			name: "Employee",
			users: []interface{}{
				Employee{age: 10},
				Employee{age: 30},
			},
			want: Employee{age: 30},
		},
		{
			name: "Customer",
			users: []interface{}{
				Customer{age: 10},
				Customer{age: 30},
			},
			want: Customer{age: 30},
		},
		{
			name: "Customer and Employee",
			users: []interface{}{
				Customer{age: 10},
				Employee{age: 30},
			},
			want: Employee{age: 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := User(tt.users...); got != tt.want {
				t.Errorf("User() = %v, want %v", got, tt.want)
			}
		})
	}
}
