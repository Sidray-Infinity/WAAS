# WAAS
------

## TO DO
- Implement wallet specific mutex to allow concurrent transaction
- Use redis

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