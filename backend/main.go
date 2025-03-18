package backend

import (
	"fmt"

	"github.com/samuelmaxi/back-go-barbershop-core/backend/config"
	"github.com/samuelmaxi/back-go-barbershop-core/backend/utils/user"
)

func ServerDataBase() {
	var database config.DataBase
	server := database.Connect()
	user := user.User{
		ID:       1,
		UserName: "Samuel",
		Password: "123",
	}

	user.Register(server)

	fmt.Println("Hello word back end")
}
