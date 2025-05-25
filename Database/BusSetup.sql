
-- ROUTES
CREATE TABLE Routes (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    DepartureTime DATETIME NOT NULL,
    BusId INT NOT NULL,
    GeoJson TEXT NOT NULL, 
    RouteName TEXT,
    FOREIGN KEY (BusId) REFERENCES Buses(Id)
);

CREATE TABLE RouteGroups (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    Name TEXT NOT NULL
);

CREATE TABLE RouteGroupRoutes (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    RouteGroupId INT NOT NULL,
    RouteId INT NOT NULL,
    FOREIGN KEY (RouteGroupId) REFERENCES RouteGroups(Id),
    FOREIGN KEY (RouteId) REFERENCES Routes(Id)
);

CREATE TABLE Stops (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    StopName TEXT NOT NULL,
    Latitude REAL NOT NULL,
    Longitude REAL NOT NULL,
    StopTypeId INT,
    FOREIGN KEY (StopTypeId) REFERENCES StopType(Id)
);

CREATE TABLE StopType (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    TypeName TEXT NOT NULL
);

CREATE TABLE RouteStops (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    RouteId INT NOT NULL,
    StopId INT NOT NULL,
    StopOrder INT NOT NULL,
    FOREIGN KEY (RouteId) REFERENCES Routes(Id),
    FOREIGN KEY (StopId) REFERENCES Stops(Id)
);

-- USERS
CREATE TABLE Users (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    EmployeeId INT NOT NULL,
    EmployeeRoleId INT NOT NULL,
    Username TEXT NOT NULL,
    Password TEXT NOT NULL,
    FOREIGN KEY (EmployeeId) REFERENCES Employees(Id),
    FOREIGN KEY (EmployeeRoleId) REFERENCES EmployeeRoles(Id)
);

-- EMPLOYEES
CREATE TABLE Employees (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    FirstName TEXT NOT NULL,
    LastName TEXT NOT NULL,
    Email TEXT NOT NULL,
    Phone TEXT
);

-- BUSES
CREATE TABLE Buses (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,  
    EmployeeId INT NOT NULL,
    BusNumber TEXT NOT NULL,
    Capacity INT NOT NULL,
    FOREIGN KEY (EmployeeId) REFERENCES Employees(Id)
);

CREATE TABLE BusRouteGroups (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    BusId INT NOT NULL,
    RouteGroupId INT NOT NULL,
    FOREIGN KEY (BusId) REFERENCES Buses(Id),
    FOREIGN KEY (RouteGroupId) REFERENCES RouteGroups(Id)
);

-- STUDENTS
CREATE TABLE Students (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    FirstName TEXT NOT NULL,
    LastName TEXT NOT NULL,
    Photo TEXT,
    Address TEXT NOT NULL,
    Latitude REAL,
    Longitude REAL,
    SchoolId INT NOT NULL,
    FOREIGN KEY (SchoolId) REFERENCES Schools(Id)
);

-- BUS-STUDENT RELATION
CREATE TABLE BusStudents (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    BusId INT NOT NULL,
    StudentId INT NOT NULL,
    FOREIGN KEY (BusId) REFERENCES Buses(Id),
    FOREIGN KEY (StudentId) REFERENCES Students(Id)
);

-- ENUM: EMPLOYEE ROLES
CREATE TABLE EmployeeRoles (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    RoleName TEXT NOT NULL
);

-- ENUM: SCHOOLS
CREATE TABLE Schools (
    Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    SchoolName TEXT NOT NULL
);

---------------------------------------------------------
-- INSERTS
---------------------------------------------------------

-- Insert Employee Roles
INSERT INTO EmployeeRoles (Id, RoleName) VALUES
(1, 'Driver'),
(2, 'Administrator');

-- Insert Schools
INSERT INTO Schools (Id, SchoolName) VALUES
(1, 'Central High School'),
(2, 'Westside Academy'),
(3, 'Eastview College');

-- Insert Employees
INSERT INTO Employees (Id, FirstName, LastName, Email, Phone) VALUES
(1, 'John', 'Doe', 'johndoe@example.com', '123-456-7890'),
(2, 'Jane', 'Smith', 'janesmith@example.com', '987-654-3210'),
(3, 'Emily', 'Johnson', 'emilyj@example.com', '555-123-4567');

-- Insert Users
INSERT INTO Users (Id, EmployeeId, EmployeeRoleId, Username, Password) VALUES
(1, 1, 1, 'johndoe', 'hashed_password_1'),
(2, 2, 1, 'janesmith', 'hashed_password_2'),
(3, 3, 2, 'emilyj', 'hashed_password_3');

-- Insert Students
INSERT INTO Students (Id, FirstName, LastName, Photo, Address, Latitude, Longitude, SchoolId) VALUES
(1, 'Michael', 'Brown', 'michael.jpg', '1212 Test St', 44.2619, -88.4154, 1),
(2, 'Sarah', 'Miller', 'sarah.jpg', '1434 Test St', 44.2719, -88.4094, 2),
(3, 'David', 'Wilson', 'david.jpg', '1989 Test St', 44.2819, -88.4194, 3);

-- Insert Buses
INSERT INTO Buses (Id, EmployeeId, BusNumber, Capacity) VALUES
(1, 1, 'BC123', 40),
(2, 2, 'YZ789', 50),
(3, 3, 'MN456', 35);

-- Insert Bus-Student Assignments
INSERT INTO BusStudents (Id, BusId, StudentId) VALUES
(1, 1, 1), 
(2, 1, 2),
(3, 2, 3);

-- Insert RouteGroups
INSERT INTO RouteGroups (Id, Name) VALUES
(1, 'Morning Routes'),
(2, 'Afternoon Routes');

-- Insert BusRouteGroups (assign buses to route groups)
INSERT INTO BusRouteGroups (Id, BusId, RouteGroupId) VALUES
(1, 1, 1),
(2, 1, 2),
(3, 2, 1);

-- Insert Routes
INSERT INTO Routes (Id, DepartureTime, BusId, GeoJson, RouteName) VALUES
(1, '2025-05-25 07:00:00', 1, '{
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

(2, '2025-05-25 15:00:00', 2, '{
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

-- Insert RouteGroupRoutes (link routes to route groups)
INSERT INTO RouteGroupRoutes (Id, RouteGroupId, RouteId) VALUES
(1, 1, 1),
(2, 2, 2);

-- Insert StopTypes
INSERT INTO StopType (Id, TypeName) VALUES
(1, 'Pickup Point'),
(2, 'Dropoff Point'),
(3, 'Cross Stop');

-- Insert Stops
INSERT INTO Stops (Id, StopName, Latitude, Longitude, StopTypeId) VALUES
(1, 'Main St & 1st Ave', 44.2620, -88.4155, 1),
(2, 'Central Park', 44.2700, -88.4100, 3),
(3, 'Elm Street', 44.2800, -88.4190, 2);

-- Insert RouteStops (assign stops to routes with order)
INSERT INTO RouteStops (Id, RouteId, StopId, StopOrder) VALUES
(1, 1, 1, 1),
(2, 1, 2, 2),
(3, 2, 3, 1);

