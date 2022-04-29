package maxage

// User - interface for user with age
type User interface {
	Age() int
}

// Employee - employee user
type Employee struct {
	age int
}

// Age - getter for Employee age field
func (e Employee) Age() int {
	return e.age
}

// Customer - customer user
type Customer struct {
	age int
}

// Age - getter for Customer age field
func (c Customer) Age() int {
	return c.age
}

// MaxAge - return max age
func MaxAge(users ...User) int {
	var max int

	for _, user := range users {
		if max < user.Age() {
			max = user.Age()
		}
	}

	return max
}
