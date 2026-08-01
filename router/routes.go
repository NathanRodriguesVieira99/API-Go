package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
* Pointer (*) é o endereco do que está sendo usado.
 */
func initializeRoutes(router *gin.Engine) {
	// * Cria um grupo de rotas.
	v1 := router.Group("/api/v1")
	{
		/*
		 * Declara uma rota.
		 */
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET opening",
			})
		})

		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST opening",
			})
		})

		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE opening",
			})
		})

		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT opening",
			})
		})

		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET openings",
			})
		})

	}

}
