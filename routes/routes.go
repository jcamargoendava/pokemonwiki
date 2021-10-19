// @title Pokemon API Swagger
// @version 1.0
// @description This is a REST API that can do all the CRUD operations like create, update, delete or display pokemons and their masters. These are JSON objects that are saved in a database hosted on a mongoDB Atlas cluster.
// @host localhost:8081
// @BasePath /api/
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

	pokemonController "github.com/jcamargoendava/pokemonwiki/controllers"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	docs "github.com/jcamargoendava/pokemonwiki/docs"
	// "../docs"

	"github.com/gin-gonic/gin"
)

func ApiMiddleware(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Set("databaseConn", db)
		c.Set("ctx", ctx)
		c.Next()
	}
}

func StartGin(ctx context.Context) {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Pokemon API Swagger"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(ApiMiddleware(ctx))

	api := router.Group("/api")
	{
		api.GET("are-you-alive", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "running"})
		})
		api.GET("/pokemons", pokemonController.GetPokemons)
		api.GET("/pokemon/:name", pokemonController.GetPokemon)
		api.POST("/pokemon", pokemonController.CreatePokemon)
		api.PUT("/pokemon/:id", pokemonController.UpdatePokemon)
		api.DELETE("/pokemon/:id", pokemonController.DeletePokemon)

		// api.GET("/masters", pokemonController.GetMasters)
		api.GET("/master/:id", pokemonController.GetMaster)
		api.POST("/master", pokemonController.CreateMaster)
		api.PUT("/master/:id", pokemonController.UpdateMaster)
		api.DELETE("/master/:id", pokemonController.DeleteMaster)
	}

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8081")

}
