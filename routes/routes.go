package routes

import (
	"micro-feature-pemilu/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		user := api.Group("/users")
		{
			user.GET("", controllers.GetAllUser)
			user.GET("/:id", controllers.GetOneUser)
		}

		article := api.Group("/articles")
		{
			article.GET("", controllers.GetAllArticles)
			article.GET("/:id", controllers.GetOneArticles)
		}

		partai := api.Group("/partais")
		{
			partai.GET("/", controllers.GetAllPartais)
			partai.GET("/:id", controllers.GetOnePartais)
		}

		paslon := api.Group("/paslons")
		{
			paslon.GET("/", controllers.GetAllPaslons)
			paslon.GET("/:id", controllers.GetOnePaslons)
		}

		vote := api.Group("/votes")
		{
			vote.GET("/", controllers.GetAllVotes)
			vote.GET("/:id", controllers.GetOneVotes)
		}
	}

	r.Run()
}
