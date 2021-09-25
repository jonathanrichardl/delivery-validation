package handlers

import (
	"delivery-validation/pkg/database"
	"delivery-validation/pkg/models"
	"fmt"
)

func (h *HTTPHandler) retrieveRequirements(order *models.Orders, id int) {
	retrievedRequirements, err := h.database.RetrieveData(fmt.Sprintf("SELECT * FROM requirements WHERE order_id = %d", id))
	if err != nil {
		h.logger.ErrorLogger.Println("Error retrieving order requirement from SQL: ", err.Error())
	}

	for retrievedRequirements.Data.Next() {
		var requirement models.Requirements
		err := retrievedRequirements.Data.Scan(&requirement.Status, &requirement.ExpectedOutcome, &requirement.Reqest)
		if err != nil {
			h.logger.ErrorLogger.Println("Error retrieving order requirement: ", err.Error())
		}
		order.Requirements = append(order.Requirements, requirement)

	}
}

func (h *HTTPHandler) retrieveOrders(retrievedData *database.RetrievedData) []models.Orders {
	var response []models.Orders
	for retrievedData.Data.Next() {
		var each models.Orders
		err := retrievedData.Data.Scan(&each.Id, &each.Title)
		if err != nil {
			h.logger.ErrorLogger.Println("Error retrieving data: ", err.Error())
		}
		h.retrieveRequirements(&each, each.Id)
		response = append(response, each)
	}
	return response
}
