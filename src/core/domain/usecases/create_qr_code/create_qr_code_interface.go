package create_qr_code

type CreateQrCodeInterface interface {
	ExecuteCreateQrCode(orderId uint) (string, error)
}
