package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password321`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginpsword := `password3214`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpsword))
	if err != nil {
		fmt.Println("you can't login")
		return

	}
	fmt.Println("you are logged in")
}
