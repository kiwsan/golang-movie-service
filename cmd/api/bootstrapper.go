package startup

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/containers"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/middlewares"
	_ "github.com/kiwsan/golang-movie-service/docs"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	router = gin.Default()
)

// @title Movie API
// @version 1.0
// @description Movies API service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func StartApplication() {

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	router.Use(middlewares.ErrorHandler())

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile(".", true)))

	goValidator := validator.New()

	// Setup route group for the API
	v1 := router.Group("/api/v1")
	{

		v1.GET("/healthcheck", containers.GetHealthCheckController().HealthCheck)
		v1.GET("/movies", containers.GetMovieFindAllController().Browse)
		v1.GET("/movies/:id", containers.GetMovieFindController().Find)
		v1.POST("/movies", containers.GetMovieCreateController(goValidator).Post)
		v1.PATCH("/movies/:id", containers.GetMovieUpdateController(goValidator).Put)
		v1.DELETE("/movies/:id", containers.GetMovieDeleteController().Delete)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(); err != nil {
		logger.Errorf("error running server", err)
	} else {
		logger.Info("Start the application...")
	}
}
