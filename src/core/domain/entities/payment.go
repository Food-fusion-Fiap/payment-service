package entities

type Payment struct {
	ID            uint   `json:"id"`
	OrderID       uint   `json:"orderId"`
	QrCode        string `json:"qrCode"`
	PaymentStatus string `json:"paymentStatus"`
	CreatedAt     string `json:"createdAt"`
}
