package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyfic/snorlax/api/database"
	"github.com/hyfic/snorlax/api/logger"
	"github.com/hyfic/snorlax/api/models"
)

// Routes handlers
func GetAllDestination(context *gin.Context) {
	rows, err := database.DB.Query("SELECT * FROM servers")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "GET", err.Error(), true)
		return
	}

	defer rows.Close()

	var result []models.Destination
	for rows.Next() {
		var destination models.Destination
		err = rows.Scan(&destination.Id, &destination.Name, &destination.Url)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			logger.RouteLog(context.ClientIP(), "GET", err.Error(), true)
			return
		}

		result = append(result, destination)
	}

	context.JSON(http.StatusOK, result)
	logger.RouteLog(context.ClientIP(), "GET", "SHOW DESTINATIONS", false)
}

func GetDestinationByName(context *gin.Context) {
	name := context.Param("name")
	smt, err := database.DB.Prepare("SELECT * FROM servers WHERE name = ?")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "GET", err.Error(), true)
		return
	}

	var destination models.Destination
	err = smt.QueryRow(name).Scan(&destination.Id, &destination.Name, &destination.Url)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "GET", err.Error(), true)
		return
	}

	context.JSON(http.StatusOK, destination)
	logger.RouteLog(context.ClientIP(), "GET", "SHOW DESTINATION "+destination.Name, false)
}

func CreateDestination(context *gin.Context) {
	var destination models.Destination
	err := context.BindJSON(&destination)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "POST", err.Error(), true)
		return
	}

	sql := `INSERT INTO servers(name, url) 
            VALUES (?, ?);`
	_, err = database.DB.Exec(sql, destination.Name, destination.Url)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "POST", err.Error(), true)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": destination.Name + " created"})
	logger.RouteLog(context.ClientIP(), "POST", "CREATED SERVER "+destination.Name, false)
}

func UpdateDestination(context *gin.Context) {
	var destination models.Destination
	err := context.BindJSON(&destination)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "POST", err.Error(), true)
		return
	}

	sql := `UPDATE servers SET name = ? url = ? WHERE id = ?`
	_, err = database.DB.Exec(sql, destination.Name, destination.Url, destination.Id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "POST", err.Error(), true)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": destination.Name + " created"})
	logger.RouteLog(context.ClientIP(), "POST", "CREATED SERVER "+destination.Name, false)
}

func DeleteDestination(context *gin.Context) {
	name := context.Param("name")
	smt, err := database.DB.Prepare("DELETE FROM servers WHERE name = ?")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "DELETE", err.Error(), true)
		return
	}

	_, err = smt.Exec(name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "DELETE", err.Error(), true)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Destination deleted"})
	logger.RouteLog(context.ClientIP(), "DELETE", "DELETE DESTINATION "+name, false)
}
