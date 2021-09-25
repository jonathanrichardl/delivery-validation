package handlers

import (
	"delivery-validation/pkg/database"
	"delivery-validation/pkg/logger"
	"delivery-validation/pkg/router"
	"net/http"
)

type HTTPHandler struct {
	database *database.DBInstance
	Router   *router.RouterInstance
	logger   *logger.LoggerInstance
}

func NewHttpHandlers(DatabaseInstance *database.DBInstance, RouterInstance *router.RouterInstance, LoggerInstance *logger.LoggerInstance) *HTTPHandler {
	return &HTTPHandler{
		database: DatabaseInstance,
		Router:   RouterInstance,
		logger:   LoggerInstance,
	}
}

func (h *HTTPHandler) RegisterAllHandlers() {
	h.Router.RegisterHandler("/", h.Index, "GET")
	h.Router.RegisterHandler("/deliveries", h.GetStatusOfAllOrders, "GET")
	//h.Router.RegisterHandler("/deliveries", h.AddNewOrder, "POST")
	h.Router.RegisterHandler("/deliveries/id={id}", h.GetStatusOfOrder, "GET")
	//h.Router.RegisterHandler("/deliveries/id={id}", h.PostUpdateOnDelivery, "POST")
}

func (h *HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {

}
