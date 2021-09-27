package handlers

import (
	"delivery-validation/pkg/database"
	"delivery-validation/pkg/logger"
	"delivery-validation/pkg/router"
	"fmt"
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
	h.Router.RegisterHandler("/orders", h.GetStatusOfAllOrders, "GET")
	h.Router.RegisterHandler("/orders", h.AddNewOrder, "POST")
	h.Router.RegisterHandler("/orders/id={id}", h.GetStatusOfOrder, "GET")
	h.Router.RegisterHandler("/orders/id={id}", h.PostUpdateOnDelivery, "POST")
}

func (h *HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
