package get_all_payments

type GetAllPaymentsInterface interface {
	ExecuteGetAllPayments() (string, error)
}
