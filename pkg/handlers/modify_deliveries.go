package handlers

import (
	"delivery-validation/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *HTTPHandler) AddNewOrder(w http.ResponseWriter, r *http.Request) {
	h.logger.InfoLogger.Println("New order received")
	var order models.Orders
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.ErrorLogger.Println("Error reading request: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request"))
	}
	err = json.Unmarshal(req, &order)
	if err != nil {
		h.logger.ErrorLogger.Println("Error unmarshalling request data: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request"))
	}
	id := h.insertOrderAndRetrieveID(order)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("201 - Order '%s' has been added, keep track on your order here at localhost:8080/orders/id=%d", order.Title, id)))
}

func (h *HTTPHandler) PostUpdateOnDelivery(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id := request["id"]
	w.Header().Set("Content-Type", "application/json")
	var form models.ProgressForm
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.ErrorLogger.Println("Error reading request: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request"))
	}
	err = json.Unmarshal(req, &form)
	if err != nil {
		h.logger.ErrorLogger.Println("Error unmarshalling request data: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request"))
	}
	h.validateRequirements(id, form)

}
