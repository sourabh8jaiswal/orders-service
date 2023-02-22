package constants

var OrderStatus = struct {
	PENDING_INVOICE   string
	INVOICE_GENERATED string
	SUCCESS           string
	FAILED            string
}{
	PENDING_INVOICE:   "PENDING_INVOICE",
	INVOICE_GENERATED: "INVOICE_GENERATED",
	SUCCESS:           "SUCCESS",
	FAILED:            "FAILED",
}

var CurrencyUnits = struct {
	USD string
	INR string
}{
	USD: "USD",
	INR: "INR",
}
