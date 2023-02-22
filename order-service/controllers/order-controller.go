package controllers

import (
	"encoding/json"
	"order-service/initializers"
	"order-service/models"
	"order-service/transformers"
	"order-service/validators"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddOrder(ctx *gin.Context) {
	// binding request data
	orderReq := models.OrderReq{}
	ctx.Bind(&orderReq)

	// validate order
	e := validators.ValidateOrder(orderReq)
	if e != nil {
		ctx.JSON(400, e)
		return
	}

	total := 0
	for _, item := range orderReq.Items {
		total = total + (item.Price * item.Quantity)
	}

	itemsString, err := json.Marshal(orderReq.Items)
	if err != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	order := models.Order{
		Status:       orderReq.Status,
		Items:        string(itemsString),
		Total:        total,
		CurrencyUnit: orderReq.CurrencyUnit,
	}

	// saving data into database
	result := initializers.DB.Create(&order)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": result.Error.Error(),
		})
		return
	}

	// getting tranformed reponse
	resp := transformers.GetOrderResponse(order)

	ctx.JSON(200, resp)
}

func UpdateOrder(ctx *gin.Context) {

	order_id, _ := strconv.Atoi(ctx.Param("id"))

	// binding request data
	orderReq := models.OrderReq{}
	ctx.Bind(&orderReq)

	// validate order
	e := validators.ValidateOrder(orderReq)
	if e != nil {
		ctx.JSON(400, e)
		return
	}

	total := 0
	for _, item := range orderReq.Items {
		total = total + (item.Price * item.Quantity)
	}

	itemsString, err := json.Marshal(orderReq.Items)
	if err != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	order := models.Order{
		ID:           order_id,
		Status:       orderReq.Status,
		Items:        string(itemsString),
		Total:        total,
		CurrencyUnit: orderReq.CurrencyUnit,
	}

	// saving data into database
	result := initializers.DB.Save(&order)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": result.Error.Error(),
		})
		return
	}

	// getting tranformed reponse
	resp := transformers.GetOrderResponse(order)

	ctx.JSON(200, resp)
}

func ListOrders(ctx *gin.Context) {
	var orders []models.Order
	// saving data into database
	result := initializers.DB.Find(&orders)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": result.Error.Error(),
		})
		return
	}

	// getting tranformed reponse
	resp := transformers.GetOrderListResponse(orders)

	ctx.JSON(200, resp)
}
