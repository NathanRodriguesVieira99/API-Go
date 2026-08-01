package router

import (
	"github.com/gin-gonic/gin"
)

/*
* Em GO quando se cria uma funcao, variável etc com a primeira letra maiuscula, é exportado automaticamente.
* Quando se cria uma funcao, variável etc com a primeira letra minuscula, é uma variável local.
 */
func Initialize() {

	// Inicialia um router com as configs padrões do framework gin gonic
	router := gin.Default()

	// Inicializa as rotas do projeto
	initializeRoutes(router)

	// roda o projeto na porta 8080 por default
	router.Run(":8080")
}
