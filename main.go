package main

import "fmt"

func main() {
	u := User{"Rachel", "Oyugi"}
	fmt.Println(u.Greeting())
	p:=Occupation{"Software Developer", "Information Technology"}
	fmt.Println(p.Job())

	
}
type User struct {
	FirstName, LastName string
}

type Occupation struct{
	Title, Department string
}
func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}
func (p Occupation) Job() string{
	return fmt.Sprintf("Rachel is a %s  in the %s  Department", p.Title, p.Department)
}
