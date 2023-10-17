package main

import "fmt"

type Profile struct {
	langName string
}

func (p *Profile) updateProfile(lang string) {
	p.langName = lang
}

func main() {
	u1 := Profile{langName: "java"}

	u1.updateProfile("golang")
	fmt.Println(u1)
}
