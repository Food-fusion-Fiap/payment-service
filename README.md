# Integration between order-service and payment-service
To integrate both services, it is needed to get order-service LoadBalancer DNS from AWS and substitute string in [orderServiceAddress](https://github.com/Food-fusion-Fiap/payment-service/blob/dea633fb7ff887c82e7bbd1abbfa22beebb5fa02/src/infra/external/order_service/get_order.go#L17).


# How to run locally
- docker compose up to run postgres and pgadmin (file docker-compose.yml)
- on src/infra/db/gorm/gorm.go, comment production db string and descoment local development connection string

# How to create interface mocks
- run `mockery --srcpkg=. --inpackage --all` at source to create mock to all interfaces
- If you want to create only one mock, use `mockery --srcpkg=. --name=NameOfInterface --inpackage --filename=mock_NameOfInterface.go`

# Fast Food FIAP - Tech Challenge - Payment Microservice

SonarCloud: https://sonarcloud.io/summary/new_code?id=Food-fusion-Fiap_payment-service

![image](https://github.com/user-attachments/assets/e830b5dd-3e12-4156-842e-39a8d2e43163)
