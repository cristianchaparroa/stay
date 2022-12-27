package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "stay/cmd/docs"
	"stay/core/app"
	"stay/core/datasources/sql"
	"stay/core/server"
	"stay/internal/adapters/handlers"
	"stay/internal/adapters/repositories"
	"stay/internal/domain"
)

//	@title			Stay API documentation
//	@version		1.0
//	@description	It contains of the stay api documentation
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
		return
	}

	c := app.NewAppConfig(8080)

	conn, err := sql.NewMySQLConnection()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	userReader := repositories.NewUserReader(conn)
	propertyReader := repositories.NewPropertyReader(conn)
	propertyWriter := repositories.NewPropertyWriter(conn)

	propertyUseCase := domain.NewPropertyUseCase(
		userReader,
		propertyWriter,
		propertyReader,
	)

	hs := []server.Handler{
		handlers.NewProperties(propertyUseCase),
	}

	s := server.New(c, hs)
	s.Run()
}
