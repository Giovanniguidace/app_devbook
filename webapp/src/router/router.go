package router

import (
	"webapp/src/rotas"

	"github.com/gorilla/mux"
)

// Gerar retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
