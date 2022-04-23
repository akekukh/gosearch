package older

// Employee - employee user
type Employee struct {
	age int
}

// Customer - customer user
type Customer struct {
	age int
}

// User - returns older person
func User(users ...interface{}) interface{} {
	var max int
	var olderUser interface{}

	for _, user := range users {
		if u, ok := user.(Employee); ok {
			if u.age > max {
				olderUser = u
				max = u.age
			}
		} else if u, ok := user.(Customer); ok {
			if u.age > max {
				olderUser = u
				max = u.age
			}
		}
	}

	return olderUser
}
