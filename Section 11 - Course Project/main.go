package main
import (
    "github.com/gin-gonic/gin"
    "example.com/rest-api/db"
    "example.com/rest-api/routes"
)

func main() {
    db.InitDB()

    server := gin.Default() // returns an Engine | Configures a http server
    routes.RegisterRoutes(server)



    server.Run(":8080") // localhost:8080
}
