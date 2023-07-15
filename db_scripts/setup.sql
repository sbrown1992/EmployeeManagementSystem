CREATE SCHEMA EMS

CREATE TABLE EMS.Department
(
    ID int NOT NULL IDENTITY(1,1),
    Name varchar(255)
    
    CONSTRAINT PK_Department PRIMARY KEY (ID)
)

CREATE TABLE EMS.Employee
(
    ID int NOT NULL IDENTITY(1,1),
    FirstName varchar(255),
    LastName varchar(255),
    Username varchar(255), 
    Password varchar(255),
    Email varchar(255),
    DepartmentID int,
    Position varchar(255),

    CONSTRAINT PK_Employee PRIMARY KEY (ID),
    CONSTRAINT FK_EmployeeDepartment FOREIGN KEY (DepartmentID)
    REFERENCES Department(ID)
)