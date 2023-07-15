CREATE SCHEMA EMS

CREATE TABLE EMS.department
(
    id int NOT NULL IDENTITY(1,1),
    name varchar(255)
    
    CONSTRAINT PK_Department PRIMARY KEY (ID)
)

CREATE TABLE EMS.employee
(
    id int NOT NULL IDENTITY(1,1),
    first_name varchar(255),
    last_name varchar(255),
    username varchar(255), 
    password varchar(255),
    email varchar(255),
    dob date,
    department_id int,
    position varchar(255),

    CONSTRAINT PK_Employee PRIMARY KEY (id),
    CONSTRAINT FK_EmployeeDepartment FOREIGN KEY (department_id)
    REFERENCES department(id)
)