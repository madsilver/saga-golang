package model

type EventType string

const (
	OrderCreated            EventType = "ORDER_CREATED"
	ProductValidated        EventType = "PRODUCT_VALIDATED"
	ProductValidationFailed EventType = "PRODUCT_VALIDATION_FAILED"
	PaymentProcessed        EventType = "PAYMENT_PROCESSED"
	PaymentFailed           EventType = "PAYMENT_FAILED"
	InventoryUpdated        EventType = "INVENTORY_UPDATED"
	InventoryFailed         EventType = "INVENTORY_FAILED"
	SagaCompleted           EventType = "SAGA_COMPLETED"
	SagaFailed              EventType = "SAGA_FAILED"
)

type Event struct {
	SagaID    string      `json:"saga_id"`
	OrderID   string      `json:"order_id"`
	Type      EventType   `json:"type"`
	Payload   interface{} `json:"payload"`
	Timestamp int64       `json:"timestamp"`
}
