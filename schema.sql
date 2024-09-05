CREATE TABLE Colleges (
    cName TEXT PRIMARY KEY,
    state TEXT,
    enrollment INT
);
CREATE TABLE Students (
    sID INT PRIMARY KEY,
    sName TEXT,
    GPA REAL,
    sizeHS INT
);
CREATE TABLE Applications (
    sID INT,
    cName TEXT,
    major TEXT,
    decision TEXT,
    metadata JSON,
    PRIMARY KEY (sID, cName),
    FOREIGN KEY (sID) REFERENCES Students(sID),
    FOREIGN KEY (cName) REFERENCES Colleges(cName)
);