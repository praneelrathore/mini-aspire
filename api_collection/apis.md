# APIs for the project

Adding curl of all apis which can be used to test the code and understand the flow of the project.



## User

### Register User

```sh
curl --location '127.0.0.1:8081/v1/user/register' \
--header 'Content-Type: application/json' \
--data '{
    "name": "test_user",
    "phone": 9876543210,
    "password": "password"
}'
```

### Submit Loan Request

```sh
curl --location '127.0.0.1:8081/v1/user/loan/request/submit' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 1,
    "amount": 11000,
    "term": 3,
    "date": "2024-08-11"
}'
```

### Get Loan Requests

```sh
curl --location '127.0.0.1:8081/v1/user/loan/request/get?user_id=1'
```

### Submit Loan Installment

```sh
curl --location '127.0.0.1:8081/v1/user/loan/request/repay' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 1,
    "loan_application_id": 1,
    "amount": 3666.67,
    "loan_repayment_id": 3
}'
```


## Admin

### Admin Registration

```sh
curl --location '127.0.0.1:8081/v1/admin/register' \
--header 'Content-Type: application/json' \
--data '{
    "name": "test_admin",
    "password": "password"
}'
```

### Get Submitted Loans

```sh
curl --location '127.0.0.1:8081/v1/admin/loan/requests?admin_id=1'
```

### Approve Loan/Reject Loan

```sh
curl --location '127.0.0.1:8081/v1/admin/loan/approve' \
--header 'Content-Type: application/json' \
--data '{
    "loan_id": 1,
    "admin_id": 1,
    "status": 2
}'
```
To reject loan, simply change the status to 4.