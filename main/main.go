package main

import (
	"flag"
	"log"
	"os"
	"practiceTwo/database"
	"practiceTwo/entities"
	"gorm.io/gorm"
	_ "practiceTwo/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files" // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type app struct {
	server *gin.Engine
	InfoLog *log.Logger
	ErrorLog *log.Logger
	Database *gorm.DB
}

var GinApp = app{
	server: GetServer(), //server.go
	InfoLog: GetInfoLog(), //logs.go
	ErrorLog: GetErrorLog(),//logs.go
}

func (app app) connectDatabase() {
	Db, err := database.OpenConnection() // database.go
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
	app.InfoLog.Println("Database connected")
	app.Database = Db
}
var migrate = flag.Bool("migrate", false, "Migrate table to DB")
var seed = flag.Uint("seed", 0, "Populate DB with fake data x values where -seed=x")

// @title           Practice Two api doc
// @version         1.0
// @description     This is a practice rest api with go.

// @host      localhost:7000
// @BasePath  /api/
func main() {
	flag.Parse()
	err := godotenv.Load()
	GinApp.connectDatabase()

	if *migrate {
		err = database.MigrateEntities(entities.Song{})
		if err != nil {
			GinApp.ErrorLog.Fatal(err)
		}
	}
	if *seed >  0 {
		database.Seed(*seed, entities.NewSong())
	}
	if err != nil {
		GinApp.ErrorLog.Fatal(err)
	}
	GetRoutes() //routes.go
	GinApp.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = GinApp.server.Run(os.Getenv("Port"))

	if err != nil {
		GinApp.ErrorLog.Fatal(err)
	}
}