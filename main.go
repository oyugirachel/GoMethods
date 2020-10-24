package main

import "fmt"

func main() {
	u := User{"Rachel", "Oyugi"}
	fmt.Println(u.Greeting())

	
}
type User struct {
	FirstName, LastName string
}

func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}
