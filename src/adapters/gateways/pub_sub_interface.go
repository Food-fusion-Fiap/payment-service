package gateways

type PubSubInterface interface {
	NotifyPaymentApproved(orderId string) error
	NotifyPaymentError(orderId string) error
}
