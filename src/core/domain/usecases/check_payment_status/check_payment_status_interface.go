package check_payment_status

type CheckPaymentStatusUseCaseInterface interface {
	ExecuteCheckPaymentStatus(orderId string) (string, error)
}
