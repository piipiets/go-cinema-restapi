## Simple REST API using Go and Gin Framework

### Sample CURL

1. Create Cinema
curl --location 'http://localhost:8080/cinema' \
--header 'Content-Type: application/json' \
--data '{
    "name": "XXI Klender",
    "location": "Jakarta",
    "rate": 4.6
  }'

2. Get All Cinema
curl --location 'http://localhost:8080/cinema'

3. Get Cinema By Id
curl --location 'http://localhost:8080/cinema/1'

4. Update Cinema
curl --location --request PUT 'http://localhost:8080/cinema/1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "CGV Paris Van Java Updated",
    "location": "Bandung",
    "rate": 4.8
  }'

5. Delete Cinema
curl --location --request DELETE 'http://localhost:8080/cinema/delete/1'