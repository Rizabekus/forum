CREATE TABLE comments(
    Name VARCHAR(30) NOT NULL,
    Text VARCHAR(200) NOT NULL,
    Id VARCHAR(40) NOT NULL,
    -- FOREIGN KEY (Id) REFERENCES posts(Id)

);