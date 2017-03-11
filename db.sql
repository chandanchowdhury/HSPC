
/*
DROP TABLE IF EXISTS EmailAddress;
DROP TABLE IF EXISTS Address;
DROP TABLE IF EXISTS School;
DROP TABLE IF EXISTS Advisor;
DROP TABLE IF EXISTS Student;
DROP TABLE IF EXISTS Team;
DROP TABLE IF EXISTS StudentTeam;
 */

CREATE TABLE EmailAddress (
    emailaddress_id INTEGER NOT NULL PRIMARY KEY
    , emailaddress VARCHAR NOT NULL
);

CREATE TABLE Address (
    address_id INTEGER NOT NULL PRIMARY KEY
    , address_country VARCHAR NOT NULL
    , address_zip   VARCHAR NOT NULL
    , address_state VARCHAR NOT NULL
    , address_city VARCHAR NOT NULL
    , address_line1 VARCHAR NOT NULL
    , address_line2 VARCHAR
);

CREATE TABLE School (
    school_id INTEGER NOT NULL PRIMARY KEY
    , school_name VARCHAR NOT NULL
    , address_id INTEGER NOT NULL
    , CONSTRAINT School_FK_address_id
            FOREIGN KEY(address_id) REFERENCES Address(address_id)
    );

/*
TODO:
    * Q: Do we need to store any other details for a student? Like Phone number.
*/
CREATE TABLE Advisor (
    advisor_id INTEGER NOT NULL PRIMARY KEY
    , advisor_name VARCHAR NOT NULL
    , emailaddress_id INTEGER NOT NULL
    , advisor_password VARCHAR NOT NULL
    , school_id INTEGER NOT NULL
    , CONSTRAINT Student_FK_school_id
        FOREIGN KEY(school_id) REFERENCES School(school_id)
    , CONSTRAINT Advisor_FK_emailaddress_id
        FOREIGN KEY(emailaddress_id) REFERENCES EmailAddress(emailaddress_id)
);

/*
TODO:
    * Q: Do we need to store any other details for a student? Like, DoB, Std etc?
*/
CREATE TABLE Student (
    student_id INTEGER NOT NULL PRIMARY KEY
    , student_name VARCHAR NOT NULL
    , school_id INTEGER
    , CONSTRAINT Student_FK_school_id
        FOREIGN KEY(school_id) REFERENCES School(school_id)
);

/*
TODO:
    * Q: Do we need to store more than one advisor for a team?
*/
CREATE TABLE Team (
    team_id INTEGER NOT NULL PRIMARY KEY
    , team_name VARCHAR NOT NULL
    , team_division CHAR(1) NOT NULL
    , advisor_id INTEGER NOT NULL
    , CONSTRAINT Team_FK_advisor_id
        FOREIGN KEY(advisor_id) REFERENCES Advisor(advisor_id)
);

CREATE TABLE StudentTeam (
    student_id INTEGER NOT NULL
    , team_id INTEGER NOT NULL
    , CONSTRAINT StudentTeam_FK_student_id
        FOREIGN KEY(student_id) REFERENCES Student(student_id)
    , CONSTRAINT StudentTeam_FK_team_id
        FOREIGN KEY(team_id) REFERENCES Team(team_id)
    , CONSTRAINT StudentTeam_PK PRIMARY KEY (student_id, team_id)
);