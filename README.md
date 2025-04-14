# ledger

To run the program: go run *.go


Assumptions:
Only supports these features:
 - Ability to record money movements (ie: deposits and withdrawals)
 - View current balance
 - View transaction history
 - Create accounts

The id of an account for simplicity will be an auto-incrementing id (int64), which means only 2^64 - 1 accounts can be created.
There is only a need to return the full transaction history so it does not support filtering the history by dates, transaction types or anything else.
Also the transaction history is stored in an array so it will be returned in ascending order by the order the transaction was created.
Only two types of transactions are supported: deposits and withdrawals.



API Documentation:

POST /accounts

Body: 

{
  name: string
}

Response:

{
    id: 0
}

Example:

curl --location 'localhost:8181/accounts' \
--header 'Content-Type: application/json' \
--data '{
    "name": "AA"
}'


POST /accounts/:accountId/transactions

Body: 

{
    date: string RFC3339 format (ex: 2025-04-20T15:04:05Z)
    type: string
    amount: float64
}

Example: 

curl --location 'localhost:8181/accounts/0/balance'


GET /accounts/:accountId/balance

Response:

{
    balance: float64
}

Example:

curl --location 'localhost:8181/accounts/0/transactions' \
--header 'Content-Type: application/json' \
--data '{
    "date": "2025-04-20T15:04:05Z",
    "type": "WITHDRAWAL",
    "amount": 7
}'


GET /accounts/:accountId/transactionHistory

Response: 

[
 {
    date: string RFC3339 format
    type: string 
    amount: float64
 }
]

Example: 

curl --location 'localhost:8181/accounts/0/transactionHistory'
