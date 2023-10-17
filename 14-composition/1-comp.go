package main

import "fmt"

type user struct {
	name string
	age  int
}

type movies struct {
	movieName string
	// anonymous field which have no name
	// anonymous field get there names from the types
	user // embedding
}

type books struct {
	bookName string
	user
}
type perms struct {
	perm string
	user
}

func (u *user) updateAge(id int) {
	u.age = id
}

func main() {
	m := movies{
		movieName: "abc",
		user: user{
			name: "xyz",
			age:  29,
		},
	}
	b := books{bookName: "op tandon", user: user{name: "abc", age: 30}}

	p := perms{perm: "admin", user: user{name: "pol", age: 8}}

	fmt.Println(m.age)
	fmt.Println(b.age)
	fmt.Println(p.age)
p.user.updateAge(0)
	m.updateAge(1)
	fmt.Println(m.age)

}
