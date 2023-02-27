package validators

import (
	"order-service/models"

	"github.com/thedevsaddam/govalidator"
)

func ValidateOrder(orderReq models.OrderReq) interface{} {
	rules := govalidator.MapData{
		"status":        []string{"required", "in:PENDING_INVOICE,INVOICE_GENERATED,SUCCESS,FAILED"},
		"items":         []string{"required"},
		"total":         []string{"numeric"},
		"currency_unit": []string{"required", "in:USD,INR"},
	}

	opts := govalidator.Options{
		Data:  &orderReq,
		Rules: rules,
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		return err
	}
	return nil
}
