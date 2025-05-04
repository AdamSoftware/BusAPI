-- set up bus code


CREATE TABLE Routes (
    RouteID INTEGER PRIMARY KEY AUTOINCREMENT,
    BusId INT NOT NULL,
    GeoJson TEXT NOT NULL, 
    RouteName TEXT NOT NULL
);

CREATE TABLE Users (
    UserID INTEGER PRIMARY KEY AUTOINCREMENT UnIQUE,
    EmployeeID INT NOT NULL,
    EmployeeRoles INT NOT NULL,
    Username TEXT NOT NULL,
    Password TEXT NOT NULL,
    FOREIGN KEY (EmployeeID) REFERENCES Employees(EmployeeID),
    FOREIGN KEY (EmployeeRoles) REFERENCES EmployeeRoles(EmployeeRoleID)
);

CREATE TABLE Employees (
    EmployeeID INTEGER PRIMARY KEY AUTOINCREMENT,
    FirstName TEXT NOT NULL,
    LastName TEXT NOT NULL,
    Email TEXT NOT NULL,
    Phone TEXT
);

CREATE TABLE Buses (
    BusID INTEGER PRIMARY KEY AUTOINCREMENT,  
    RouteID INT NOT NULL,
    EmployeeID INT NOT NULL,
    BusNumber TEXT NOT NULL,
    Capacity INT NOT NULL,
    FOREIGN KEY (RouteID) REFERENCES Routes(RouteID),
    FOREIGN KEY (EmployeeID) REFERENCES Employees(EmployeeID)
);

CREATE TABLE Students (
    StudentID INTEGER PRIMARY KEY AUTOINCREMENT,
    FirstName TEXT NOT NULL,
    LastName TEXT NOT NULL,
    Photo TEXT,
    SchoolID INT NOT NULL,
    FOREIGN KEY (SchoolID) REFERENCES Schools(SchoolID)
);

CREATE TABLE BusStudents (
    BusStudentID INTEGER PRIMARY KEY AUTOINCREMENT,
    BusID INT NOT NULL,
    StudentID INT NOT NULL,
    FOREIGN KEY (BusID) REFERENCES Buses(BusID),
    FOREIGN KEY (StudentID) REFERENCES Students(StudentID)
);

-- Enum tables

CREATE TABLE EmployeeRoles (
    EmployeeRoleID INTEGER PRIMARY KEY AUTOINCREMENT,
    RoleName TEXT NOT NULL
);

CREATE TABLE Schools (
    SchoolID INTEGER PRIMARY KEY AUTOINCREMENT,
    SchoolName TEXT NOT NULL
);



-- Insert Employee Roles
INSERT INTO EmployeeRoles (EmployeeRoleID,RoleName) VALUES
(1,'Driver'),
(2,'Administrator');

-- Insert Schools
INSERT INTO Schools (SchoolID,SchoolName) VALUES
(1,'Central High School'),
(2,'Westside Academy'),
(3,'Eastview College');

-- Insert Employees
INSERT INTO Employees (EmployeeID,FirstName, LastName, Email, Phone) VALUES
(1,'John', 'Doe', 'johndoe@example.com', '123-456-7890'),
(2,'Jane', 'Smith', 'janesmith@example.com', '987-654-3210'),
(3,'Emily', 'Johnson', 'emilyj@example.com', '555-123-4567');

-- Insert Users (No Explicit UserID)
INSERT INTO Users (UserID, EmployeeID, EmployeeRoles, Username, Password) VALUES
(1,1, 1, 'johndoe', 'hashed_password_1'),
(2,2, 1, 'janesmith', 'hashed_password_2'),
(3,3, 2, 'emilyj', 'hashed_password_3');

-- Insert Students
INSERT INTO Students (StudentID,FirstName, LastName, Photo, SchoolID) VALUES
(1,'Michael', 'Brown', 'michael.jpg', 1),
(2,'Sarah', 'Miller', 'sarah.jpg', 2),
(3,'David', 'Wilson', 'david.jpg', 3);

-- Insert Buses
INSERT INTO Buses (BusID,RouteID, EmployeeID, BusNumber, Capacity) VALUES
(1,1, 1, 'BC123', 40),
(2,2, 2, 'YZ789', 50),
(3,1, 3, 'MN456', 35);

-- Insert Bus-Student Assignments
INSERT INTO BusStudents (BusStudentID, BusID, StudentID) VALUES
(1,1, 1), 
(3,1, 2),
(2,2, 3);

-- Insert Routes
INSERT INTO Routes (RouteID,BusID, GeoJson, RouteName) VALUES
(1,1, '{
    "type": "Feature",
    "geometry": {
        "type": "LineString",
        "coordinates": [
            [-88.4154, 44.2619], 
            [-88.4094, 44.2719]
        ]
    },
    "properties": {
        "name": "Route A",
        "color": "#FF0000"
    }
}', 'Route A'),

(2,2, '{
    "type": "Feature",
    "geometry": {
        "type": "LineString",
        "coordinates": [
            [-88.4154, 44.2619], 
            [-88.4194, 44.2819]
        ]
    },
    "properties": {
        "name": "Route B",
        "color": "#0000FF"
    }
}', 'Route B');

