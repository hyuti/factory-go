package main

import (
	"fmt"

	"github.com/hyuti/factory-go/factory"
)

type Group struct {
	ID int
}

type User struct {
	ID     int
	Groups []*Group
}

var UserFactory = factory.NewFactory(
	&User{},
).SeqInt("ID", func(n int) (any, error) {
	return n, nil
})

func main() {
	for i := 1; i <= 3; i++ {
		user := UserFactory.MustCreateWithOption(map[string]any{
			"Groups": []*Group{
				{i}, {i + 1},
			},
		}).(*User)
		fmt.Println("ID:", user.ID)
		for _, group := range user.Groups {
			fmt.Println(" Group.ID:", group.ID)
		}
	}
}
