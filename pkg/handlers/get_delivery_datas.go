package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *HTTPHandler) GetStatusOfAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	retrievedData, err := h.database.RetrieveData("SELECT * FROM orders")
	if err != nil {
		h.logger.ErrorLogger.Println("Can't retrieve news from SQL: ", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response := h.retrieveOrders(retrievedData)
	resp, _ := json.Marshal(response)
	fmt.Fprintf(w, string(resp))
}

func (h *HTTPHandler) GetStatusOfOrder(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id := request["id"]
	w.Header().Set("Content-Type", "application/json")
	retrievedData, err := h.database.RetrieveData(fmt.Sprintf("SELECT * FROM orders where id = %s", id))
	if err != nil {
		h.logger.ErrorLogger.Println("Can't retrieve news from SQL: ", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response := h.retrieveOrders(retrievedData)
	resp, _ := json.Marshal(response)
	fmt.Fprintf(w, string(resp))

}
