## Simple REST API using Go and Gin Framework

### Base URL

```text
https://go-cinema-restapi.up.railway.app
```

### Database Environment

Set these environment variables before running the app:

```env
PGHOST=localhost
PGPORT=5432
PGDATABASE=cinema
PGUSER=postgres
PGPASSWORD=admin
DB_ENGINE=postgres
```

### Sample CURL

**1. Create Cinema**
<br> curl --location 'https://go-cinema-restapi.up.railway.app/cinema' \
--header 'Content-Type: application/json' \
--data '{
    "name": "XXI Klender",
    "location": "Jakarta",
    "rate": 4.6
  }'

**2. Get All Cinema**
<br> curl --location 'https://go-cinema-restapi.up.railway.app/cinema'

**3. Get Cinema By Id**
<br> curl --location 'https://go-cinema-restapi.up.railway.app/cinema/1'

**4. Update Cinema**
<br> curl --location --request PUT 'https://go-cinema-restapi.up.railway.app/cinema/1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "CGV Paris Van Java Updated",
    "location": "Bandung",
    "rate": 4.8
  }'

**5. Delete Cinema**
<br> curl --location --request DELETE 'https://go-cinema-restapi.up.railway.app/cinema/delete/1'
