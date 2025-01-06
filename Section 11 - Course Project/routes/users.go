package routes
import (
    "net/http"
    "github.com/gin-gonic/gin"
    "example.com/rest-api/models"
    "example.com/rest-api/utils"
)

func signup(context *gin.Context) {
    var user models.User
    err := context.ShouldBindJSON(&user) // Works similar to 'fmt.Scan()' | data sent must be in same format to 'models.Event'
    if err!=nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error_message": err.Error()})
        return
    }

    err = user.Save()
    if err!=nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user.", "error_message": err.Error()})
        return
    }
    context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

func login(context *gin.Context) {
    var user models.User
    err := context.ShouldBindJSON(&user)
    if err!=nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error_message": err.Error()})
        return
    }

    err = user.ValidadeCredentials()
    if err!=nil {
        context.JSON(http.StatusUnauthorized, gin.H{ "message": "Could not authenticate user.", "error_message": err.Error()})
        return
    }

    token, err := utils.GenerateToken(user.Email, user.Id)
    if err!=nil {
        context.JSON(http.StatusInternalServerError, gin.H{ "message": "Could not authenticate user.", "error_message": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "login successfull.", "token": token})
}