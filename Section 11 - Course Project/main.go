package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
    "example.com/rest-api/db"
    "example.com/rest-api/models"
)

func main() {
    db.InitDB()

    server := gin.Default() // returns an Engine | Configures a http server
    server.GET("/events", getEvents)
    server.POST("/events", createEvent)




    server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
    events := models.GetAllEvents()
    context.JSON(http.StatusOK, events) // To return a response to the server it must be this way
}
func createEvent(context *gin.Context) {
    var event models.Event
    err := context.ShouldBindJSON(&event) // Works similar to 'fmt.Scan()' | data sent must be in same format to 'models.Event'
    if err!=nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
        return
    }

    event.Id = 1
    event.UserId = 1
    event.Save()
    context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}