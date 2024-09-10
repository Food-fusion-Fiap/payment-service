# Integração entre order-service e payment-service
Para integrá-los, é necessário pegar o DNS do LoadBalancer do order-service na AWS e substituir a string em [orderServiceAddress](https://github.com/Food-fusion-Fiap/payment-service/blob/dea633fb7ff887c82e7bbd1abbfa22beebb5fa02/src/infra/external/order_service/get_order.go#L17).

# Como rodar localmente
- `docker compose up` para rodar a base (postgres e pgadmin) (arquivo docker-compose.yml)
- em src/infra/db/gorm/gorm.go, commente a string de produção e descomente a string local
- rode a aplicação pela IDE ou pelo comando `go run ./main.go`

# Como criar mocks de interfaces
- rode `mockery --srcpkg=. --inpackage --all` no src para criar a mock de todas as interfaces
- Se quiser criar apenas um mock, use `mockery --srcpkg=. --name=NameOfInterface --inpackage --filename=mock_NameOfInterface.go`

# Fast Food FIAP - Tech Challenge - Payment Microservice (Entrega da Fase 4)
SonarCloud: https://sonarcloud.io/summary/new_code?id=Food-fusion-Fiap_payment-service
![image](https://github.com/user-attachments/assets/e830b5dd-3e12-4156-842e-39a8d2e43163)
