package domain

type User struct {
	Name string
	Age  int
}

func GetMatt() *User {
	return &User{"Matt", 33}
}

func GetJohn() *User {
	return &User{"John", 42}
}
