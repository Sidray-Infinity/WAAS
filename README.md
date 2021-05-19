# WAAS
------

## Future Enhancement
- Bloom filter to check if username is present
- RPC architecture
- Load data from config file

## Stack
- gorrila-mux
- gorm
- Manual dependency injection

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

## Demo
- Register User
- Create wallet
- Credit amount
- Show row lock by implementing credit and block
- Generate CSV