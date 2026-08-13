## Simple REST API using Go and Gin Framework

### Sample CURL
curl --location 'http://localhost:8080/cinema' \
--header 'Content-Type: application/json' \
--data '{
    "name": "XXI Klender",
    "location": "Jakarta",
    "rate": 4.6
  }'
