# Requirements

## Product

1. Setup Docker Services []
2. Create an Azure SQL Database []
3. Create an Employee GraphQL Schema []
4. API Endpoints []
5. GraphQL Operations []
6. JWT Middleware []
7. Email Notification []
8. Bonus []

## Delivery

1. Code exists on Github []
1. A Readme file exists with docs (and links to related docs) []
1. A Postman collection (or similar) exists for testing endpoints []
1. A Docker Compose file exists for setting up services []
1. Any scripts requried to setup the database []

## Details

### Setup Docker Services

* A docker compose file exists which starts [X]
    * An Azure SQL Database [X]
    * A Mailhog SMTP Server [X]
* How to start and stop the services is documented [X]
* The Azure SQL Database has persistent storage []

### Create an Azure SQL Database

* All of the following happens via automation []
* An Employee Table has been setup []
    * ID (Primary Key, automatically increments, not null) []
    * First Name []
    * Last Name []
    * Username []
    * Password []
    * Email []
    * Date of Birth []
    * DepartmentID (Foreign key to Department Table) []
    * Position []
* A Department Table has been setup []
    * ID (Primary Key, automatically increments, not null) []
    * Name []
* How to manage the database is documented []
* A sample set of testdata exists []

### Create an Employee GraphQL Schema

* A schema exists which describes an employee like above []

### API Endpoints

* A webserver exists and listens on a port []
* The default port is documented []

#### Login

* Code exists to handle a POST request at `/login` []
* The endpoint only accepts `username` and `password` fields []
* Upon successful login, return a valid JWT for the user with status 200 []
* (bonus) Upon failed login
    * If username or password was not provided, return 400 bad request []
    * If authentication fails, return 403 forbidden []
    * If the error is unknown, return 500 internal server error []
* There are curl script(s) to test this endpoint []
* This endpoint is documented

#### Employee

* Code exists to handle GET, POST, PUT, and DELETE requests at `/employee` []
    * Full detail is in the GraphQL Operations section
* Each of these requests requires a valid JWT []
    * Full detail is in the JWT Middleware section []
* The errors returned use standard HTTP Codes []
* There are curl scripts to test these endpoints []
    * GET []
    * POST []
    * PUT []
    * DELETE []
* These endpoints are documented []
    * GET []
    * POST []
    * PUT []
    * DELETE []

### GraphQL Operations

* The following operations are supported
    * Get a list of all employees []
        * With filtering []
        * With sorting []
        * With Pagination []
    * Get a specific employee by id []
    * Get the current employee by context []
    * Add a new employee []
    * Update an existing employee []
    * Delete an employee []
* TODO add more details on individual endpoints []

### JWT Middleware

* Middleware exists to check the validity of a JWT []
* If no JWT is provided, return 401 Unauthorized []
* If an invalid JWT is provided, return 403 Forbidden []
* Document somewhere that the actual requirements doc states that `The request should only be
accepted if the JWT is provided or is valid.` This seems incorrect since we only want to accept
valid JWTs, not simply all those provided []

### Email Notification

* The `Add New Employee` functionality also emails the new employee []
* The email is branded and professional in appearance []
    * Sbrown Industries branding exists []
