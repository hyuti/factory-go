package main

import (
	"fmt"

	"github.com/hyuti/factory-go/factory"
)

type User struct {
	ID       int
	Name     string
	Location string
}

// 'Location: "Tokyo"' is default value.
var UserFactory = factory.NewFactory(
	&User{Location: "Tokyo"},
).SeqInt("ID", func(n int) (any, error) {
	return n, nil
}).Attr("Name", func(args factory.Args) (any, error) {
	user := args.Instance().(*User)
	return fmt.Sprintf("user-%d", user.ID), nil
})

func main() {
	for i := 0; i < 3; i++ {
		user := UserFactory.MustCreate().(*User)
		fmt.Println("ID:", user.ID, " Name:", user.Name, " Location:", user.Location)
	}
}
