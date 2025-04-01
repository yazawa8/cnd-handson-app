# User Service

## Description
## Quick Start
user登録をする
```bash
curl -X POST http://localhost:8080/user/register \
     -H "Content-Type: application/json" \
     -d '{
           "email": "test@example.com",
           "password": "securepassword"
         }'
```     
userログインをする
```bash
curl -X POST http://localhost:8080/user/login \
     -H "Content-Type: application/json" \
     -d '{
           "email": "test@example.com",
           "password": "securepassword"
         }'
```
