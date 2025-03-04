package routes

import (
    "net/http"
    "strconv"
    "example.com/rest-api/models"
    "github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err!=nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id.", "error_message": err.Error()})
        return
    }

    event, err := models.GetEventById(eventId)
    if err!=nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.", "error_message": err.Error()})
        return
    }
    context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
    events, err := models.GetAllEvents()
    if err!=nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.", "error_message": err.Error()})
        return
    }

    context.JSON(http.StatusOK, events) // To return a response to the server it must be this way
}

func createEvent(context *gin.Context) {
    var event models.Event
    err := context.ShouldBindJSON(&event) // Works similar to 'fmt.Scan()' | data sent must be in same format to 'models.Event'
    if err!=nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error_message": err.Error()})
        return
    }

    event.UserId = context.GetInt64("userId")

    err = event.Save()
    if err!=nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events.", "error_message": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err!=nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id.", "error_message": err.Error()})
        return
    }

    event, err := models.GetEventById(eventId)
    if err!=nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.", "error_message": err.Error()})
        return
    }

    userId := context.GetInt64("userId")
    if event.UserId != userId {
        context.JSON(http.StatusUnauthorized, gin.H{ "message": "Not authorized to update event."})
        return
    }

    var updateEvent models.Event
    err = context.ShouldBindJSON(&updateEvent) // Works similar to 'fmt.Scan()' | data sent must be in same format to 'models.Event'
    if err!=nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error_message": err.Error()})
        return
    }

    updateEvent.Id = eventId
    err = updateEvent.Update()
    if err!=nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event.", "error_message": err.Error()})
        return
    }
    context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvent(context *gin.Context) {
    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err!=nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id.", "error_message": err.Error()})
        return
    }

    event, err := models.GetEventById(eventId)
    if err!=nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.", "error_message": err.Error()})
        return
    }

    userId := context.GetInt64("userId")
    if event.UserId != userId {
        context.JSON(http.StatusUnauthorized, gin.H{ "message": "Not authorized to delete event."})
        return
    }

    err = event.Delele()
    if err!=nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event.", "error_message": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}