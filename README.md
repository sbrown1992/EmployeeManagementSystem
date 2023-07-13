# EmployeeManagementSystem
Tech Challenge for Scotland based employer

## Requirements

* Docker
* Docker Compose

## Setup

* create a `.env` file in the same directory as `docker-compose.yml`. This will contain the
password needed for the database
```
MSSQL_SA_PASSWORD=yourStrongPassword1234
```

## Running

In order to run the app, run `docker compose up`. This will start all of the necessary services
