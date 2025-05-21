package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyfic/snorlax/api/file"
	"github.com/hyfic/snorlax/api/logger"
	"github.com/hyfic/snorlax/api/util"
)

// Routes handlers
func PingRoute(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "ok"})
	logger.RouteLog(context.ClientIP(), "GET", "PING", false)
}

func CreateFolderRoute(context *gin.Context) {
	var requestBody RequestBody
	if GetBodyFromRequest(context, &requestBody) != nil {
		return
	}

	err := file.CreateFolder(requestBody.Path)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "POST", err.Error(), true)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": requestBody.Path + " created"})
	logger.RouteLog(context.ClientIP(), "POST", "CREATED FOLDER "+requestBody.Path, false)
}

func ViewFolderRoute(context *gin.Context) {
	path, pathErr := GetPathFromParams(context)

	if pathErr != nil {
		return
	}

	files, err := file.ReadFolder(path)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "GET", err.Error(), true)
		return
	}

	context.JSON(http.StatusOK, files)
	logger.RouteLog(context.ClientIP(), "GET", "VIEW FOLDER "+path, false)
}

func GetFileInfoRoute(context *gin.Context) {
	path, pathErr := GetPathFromParams(context)

	if pathErr != nil {
		return
	}

	file, err := file.GetFileInfo(path)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "GET", err.Error(), true)
		return
	}

	context.JSON(http.StatusOK, file)
	logger.RouteLog(context.ClientIP(), "GET", "GET FILE INFO "+path, false)

}

func DownloadRoute(context *gin.Context) {
	path := context.Query("path")
	fileName := context.Query("name")

	if len(path) == 0 || len(fileName) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "path/name is not given."})
		context.Abort()
		logger.RouteLog(context.ClientIP(), "GET", "path/name IS NOT GIVEN IN REQUEST QUERY", true)
		return
	}

	context.FileAttachment(util.StorageFolder+path, fileName)
	logger.RouteLog(context.ClientIP(), "GET", "DOWNLOAD "+path, false)
}

func RenameFileRoute(context *gin.Context) {
	var requestBody PutRequestBody

	if GetBodyFromRequest(context, &requestBody) != nil {
		return
	}

	err := file.RenameFile(requestBody.OldPath, requestBody.NewPath)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "PUT", err.Error(), true)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Renamed " + requestBody.OldPath + " to " + requestBody.NewPath})
	logger.RouteLog(context.ClientIP(), "PUT", "RENAMED "+requestBody.OldPath+" TO "+requestBody.NewPath, false)
}

func DeleteFileRoute(context *gin.Context) {
	path, pathErr := GetPathFromParams(context)

	if pathErr != nil {
		return
	}

	err := file.DeleteFile(path)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "DELETE", err.Error(), true)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "deleted " + path + " successfully"})
	logger.RouteLog(context.ClientIP(), "DELETE", "DELETED "+path, false)
}

func FileUploadRoute(context *gin.Context) {
	fileName := context.PostForm("fileName")
	filePath := context.PostForm("filePath")

	if len(fileName) == 0 || len(filePath) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "fileName or filePath is not provided"})
		logger.RouteLog(context.ClientIP(), "POST", "fileName OR filePath IS NOT GIVEN IN REQUEST QUERY", true)
		return
	}

	uploadedFile, err := context.FormFile("file")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "POST", err.Error(), true)
		return
	}

	err = context.SaveUploadedFile(uploadedFile, util.StorageFolder+filePath+"/"+fileName)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.RouteLog(context.ClientIP(), "POST", err.Error(), true)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Uploaded file successfully."})
	logger.RouteLog(context.ClientIP(), "POST", "UPLOADED "+fileName+" TO "+filePath, false)
}
