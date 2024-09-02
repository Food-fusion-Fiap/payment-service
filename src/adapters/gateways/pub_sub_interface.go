package gateways

type PubSubInterface interface {
	NotifyPaymentApproved(orderId uint) error
	NotifyPaymentError(orderId uint) error
}
