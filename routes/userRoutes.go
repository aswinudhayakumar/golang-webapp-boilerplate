package routes

import (
	"encoding/json"
	"net/http"

	"github.com/aswinudhayakumar/golang-boilerplate/controller"
	"github.com/aswinudhayakumar/golang-boilerplate/model"
	"github.com/gorilla/mux"
)

//HandleUserRoutes handles all user routes
func HandleUserRoutes(userController controller.UserController, router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var output model.Output
		output = controller.SayHello(userController.DB)
		json.NewEncoder(w).Encode(output)

	}).Methods("GET")

}
