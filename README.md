# WAAS
------

## TO DO
- Add validation of data in all insert APIs
- Check if username / KYC is duplicated
- Dependency Injection 

## Stack
- gorrila-mux
- gorm
- ```Think something for dependency injection```

## Ideas
- Should use worker for credit/debit? : Probably no, loss of sequence

## Swagger
- swagger generate spec -o ./swagger.yaml --scan-models
- swagger serve -F=swagger swagger.yaml

## Refactoring
- Controller
    - Routes
    - Model specific handlers

- Domain
    - Logic part (likely to be empty)

- Model
    - Entity
    - Impl : Interaction with DB
    - Views : Request params -> Entity (or the other way)

## References
### Distributed locks
- https://kylewbanks.com/blog/distributed-locks-using-golang-and-redis
