# WAAS
------

## TO DO
- Add validation of data in all insert APIs

## Stack
- gorrila-mux
- gorm
- ```Think something for dependency injection```

## Ideas
- Should use worker for credit/debit? : Probably no, loss of sequence

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