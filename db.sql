DROP TABLE EmailAddress;
CREATE TABLE EmailAddress (
    emailaddress_id INTEGER NOT NULL
    , emailaddress VARCHAR NOT NULL
);

DROP TABLE Address;
CREATE TABLE Address (
    address_id INTEGER NOT NULL
    , address_country VARCHAR NOT NULL
    , address_zip   VARCHAR NOT NULL
    , address_state VARCHAR NOT NULL
    , address_city VARCHAR NOT NULL
    , address_line1 VARCHAR NOT NULL
    , address_line2 VARCHAR
);

DROP TABLE School;
CREATE TABLE School (
    school_id INTEGER NOT NULL
    , school_name VARCHAR NOT NULL
    , adress_id INTEGER NOT NULL
    );

/*
TODO:
Q: Do we need to store any other details for a student? Like Phone number.
*/
DROP TABLE Advisor;
CREATE TABLE Advisor (
    advisor_id INTEGER NOT NULL
    , advisor_name VARCHAR NOT NULL
    , emailaddress_id INTEGER NOT NULL
    , adress_password VARCHAR NOT NULL
    , school_id INTEGER NOT NULL
);

/*
TODO:
Q: Do we need to store any other details for a student? Like, DoB, Std etc?
*/
DROP TABLE Student;
CREATE TABLE Student (
    student_id INTEGER NOT NULL
    , student_name VARCHAR NOT NULL,
    , school_id INTEGER
);

/*
TODO:
Q: Do we need to store more than one advisor for a team?
*/
DROP TABLE Team (
    team_id INTEGER NOT NULL
    , team_name VARCHAR NOT NULL,
    , team_division CHAR(1) NOT NULL,
    , advisor_id INTEGER NOT NULL,
);


DROP TABLE StudentTeam (
    student_id INTEGER NOT NULL
    , team_id INTEGER NOT NULL
);