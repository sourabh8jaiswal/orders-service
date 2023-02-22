package transformers

import (
	"encoding/json"
	"order-service/models"
)

type OrderResponse struct {
	ID           int                `json:"id"`
	Status       string             `json:"status"`
	Items        []models.OrderItem `json:"items"`
	Total        int                `json:"total"`
	CurrencyUnit string             `json:"currency_unit"`
}

func GetOrderResponse(order models.Order) OrderResponse {
	var items []models.OrderItem
	json.Unmarshal([]byte(order.Items), &items)

	resp := OrderResponse{
		ID:           order.ID,
		Status:       order.Status,
		Items:        items,
		Total:        order.Total,
		CurrencyUnit: order.CurrencyUnit,
	}
	return resp
}

func GetOrderListResponse(orders []models.Order) []OrderResponse {
	var resp []OrderResponse = []OrderResponse{}
	for _, el := range orders {
		order := GetOrderResponse(el)
		resp = append(resp, order)
	}
	return resp
}
