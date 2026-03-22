package domain

type User struct {
	ID      int
	Version int

	FulltName   string
	PhoneNumber *string
	// Email     string
	// Password  string
}
