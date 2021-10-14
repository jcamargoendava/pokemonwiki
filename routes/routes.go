// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
package routes

import (
	"context"
	"net/http"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// "github.com/jcamargoendava/pokemonwiki/docs"

	"github.com/gin-gonic/gin"
	pokemonController "github.com/jcamargoendava/pokemonwiki/controllers"
)

func ApiMiddleware(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Set("databaseConn", db)
		c.Set("ctx", ctx)
		c.Next()
	}
}

func StartGin(ctx context.Context) {
	router := gin.Default()
	router.Use(ApiMiddleware(ctx))
	// api := router.Group("/api")

	router.GET("are-you-alive", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "running"})
	})
	router.GET("/pokemons", pokemonController.GetPokemons)
	router.GET("/pokemon/:name", pokemonController.GetPokemon)
	router.POST("/pokemon", pokemonController.CreatePokemon)
	router.PUT("/pokemon/:id", pokemonController.UpdatePokemon)
	router.DELETE("/pokemon/:id", pokemonController.DeletePokemon)

	router.GET("/master/:id", pokemonController.GetMaster)
	router.POST("/master", pokemonController.CreateMaster)
	router.PUT("/master/:id", pokemonController.UpdateMaster)
	router.DELETE("/master/:id", pokemonController.DeleteMaster)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8081")
}
