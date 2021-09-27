package handlers

import (
	"delivery-validation/pkg/database"
	"delivery-validation/pkg/models"
	"fmt"
)

func (h *HTTPHandler) retrieveRequirements(order *models.Orders, id int) {
	retrievedRequirements, err := h.database.RetrieveData(fmt.Sprintf("SELECT requirementid, request, expectedoutcome, status FROM requirements WHERE order_id = %d", id))
	if err != nil {
		h.logger.ErrorLogger.Println("Error retrieving order requirement from SQL: ", err.Error())
	}

	for retrievedRequirements.Data.Next() {
		var requirement models.Requirements
		err := retrievedRequirements.Data.Scan(&requirement.Requirementid, &requirement.Request,
			&requirement.ExpectedOutcome, &requirement.Status)
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

func (h *HTTPHandler) insertOrderAndRetrieveID(newOrder models.Orders) int {
	id := h.assignIdToOrder()
	h.database.AddData(fmt.Sprintf("INSERT INTO orders (id, title) VALUES (%d,'%s');", id, newOrder.Title))
	for _, requirement := range newOrder.Requirements {
		h.database.AddData(fmt.Sprintf("INSERT INTO requirements (request,expectedoutcome,order_id,status) VALUES ('%s', '%s', %d, 0)",
			requirement.Request, requirement.ExpectedOutcome, id))
	}
	return id
}

func (h *HTTPHandler) assignIdToOrder() int {
	var count int
	retrievedData, err := h.database.RetrieveData("SELECT COUNT(*) FROM orders")
	if err != nil {
		h.logger.ErrorLogger.Println("Database Error: ", err.Error())
	}
	retrievedData.Data.Next()
	err = retrievedData.Data.Scan(&count)
	if err != nil {
		h.logger.ErrorLogger.Println("Error retrieving data: ", err.Error())
	}
	count++
	return count
}

func (h *HTTPHandler) validateRequirements(id string, form models.ProgressForm) {
	retrievedData, err := h.database.RetrieveData(fmt.Sprintf("SELECT * FROM orders where id = %s", id))
	if err != nil {
		h.logger.ErrorLogger.Println("Database Error : ", err.Error())
		return
	}
	order := h.retrieveOrders(retrievedData)[0]
	for _, resp := range form.Fufillments {
		for _, requirement := range order.Requirements {
			if requirement.Requirementid == int(resp.Requirementid) {
				if resp.Outcome == requirement.ExpectedOutcome {
					h.statusComplete(requirement.Requirementid)
				}
				break
			}
		}
	}

}

func (h *HTTPHandler) statusComplete(requirementID int) {
	if err := h.database.UpdateData(fmt.Sprintf("UPDATE requirements SET status = 1 WHERE requirementid = %d;",
		requirementID)); err != nil {
		h.logger.ErrorLogger.Println("Can't modify database : ", err.Error())
	}
}
