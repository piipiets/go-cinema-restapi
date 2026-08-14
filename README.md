## Simple REST API using Go and Gin Framework

### Sample CURL

**1. Create Cinema**
<br> curl --location 'http://localhost:8080/cinema' \
--header 'Content-Type: application/json' \
--data '{
    "name": "XXI Klender",
    "location": "Jakarta",
    "rate": 4.6
  }'

**3. Get All Cinema**
<br> curl --location 'http://localhost:8080/cinema'

**4. Get Cinema By Id**
<br> curl --location 'http://localhost:8080/cinema/1'

**5. Update Cinema**
<br> curl --location --request PUT 'http://localhost:8080/cinema/1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "CGV Paris Van Java Updated",
    "location": "Bandung",
    "rate": 4.8
  }'

**6. Delete Cinema**
<br> curl --location --request DELETE 'http://localhost:8080/cinema/delete/1'
