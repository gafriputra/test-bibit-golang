-- this solution is written using MySQL syntax

-- create USER table
CREATE TABLE USER(ID INTEGER, UserName VARCHAR(25), Parent INTEGER);

-- insert data to USER table
INSERT INTO USER(ID, UserName, Parent) 
    VALUES (1, "Ali", 2), (2, "Budi", 0), (3, "Cecep", 1);

-- query to get all user data with their respctive "Creator"
SELECT u.ID, u.UserName, p.UserName AS ParentUserName
FROM user u
LEFT JOIN user p ON p.ID = u.Parent
ORDER BY u.ID ASC