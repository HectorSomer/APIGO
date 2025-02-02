package main

import (
	"api-hexagonal/src/products/infraestructure"
	"api-hexagonal/src/sells/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("No se cargó el .env")
    }
	routers:= gin.Default()
	infraestructure.Init(routers)
    dependencies.Init(routers)
	if err := routers.Run(); err != nil {
		panic(err)
	}
}