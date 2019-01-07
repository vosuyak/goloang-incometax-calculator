package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// router = ExpensesRouter(router)
	router = SalaryRouter(router)

	return router
}
