package router

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hyfic/snorlax/api/logger"
	"github.com/hyfic/snorlax/api/util"
)

var route *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	route = gin.New()
}

func StartServer(port int32) {
	route.Use(cors.Default())

	route.GET("/ping", PingRoute) // route to check if server is up

	// Routes
	fileApi := route.Group("file")

	fileApi.GET("/view-folder", ViewFolderRoute)
	fileApi.GET("/get-file-info", GetFileInfoRoute)
	fileApi.GET("/download", DownloadRoute)

	fileApi.POST("/create-folder", CreateFolderRoute)
	fileApi.POST("/upload", FileUploadRoute)

	fileApi.PUT("/rename-file", RenameFileRoute)
	fileApi.DELETE("/delete-file", DeleteFileRoute)

	destinationApi := route.Group("destination")
	destinationApi.GET("/all", GetAllDestination)
	destinationApi.GET("/:name", GetDestinationByName)
	destinationApi.POST("/create", CreateDestination)
	destinationApi.PUT("/update", UpdateDestination)
	destinationApi.DELETE("/delete/:name", DeleteDestination)

	// set storage as static folder
	fileApi.Static("/storage", util.StorageFolder)

	fileApi.Use()

	// listen server on port
	addr := fmt.Sprintf(":%v", port)

	fmt.Println("=======================")
	logger.Success(fmt.Sprintf("[+] SERVER STARTED AT PORT %v", port))
	logger.Info(fmt.Sprintf("[i] http://127.0.0.1" + addr))

	ips, err := util.LocalIP()

	if err == nil {
		for _, ip := range ips {
			logger.Info(fmt.Sprintf("[i] http://" + ip.String() + addr))
		}
	}
	fmt.Println("=======================")

	route.Run(addr)
}
