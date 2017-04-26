/*
Maintain the order as they follow they dependency.
*/
DROP TABLE IF EXISTS parking;
DROP TABLE IF EXISTS team_score;
DROP TABLE IF EXISTS student;
DROP TABLE IF EXISTS team;
DROP TABLE IF EXISTS school_advisor;
DROP TABLE IF EXISTS advisor;
DROP TABLE IF EXISTS school;
DROP TABLE IF EXISTS credential;
DROP TABLE IF EXISTS address;
DROP TABLE IF EXISTS session;


CREATE TABLE credential (
    credential_id SERIAL NOT NULL PRIMARY KEY
    , emailaddress VARCHAR NOT NULL
    , password_hash VARCHAR
    , credential_active BOOLEAN DEFAULT FALSE
    , CONSTRAINT EmailAddress_unique
        UNIQUE (emailaddress)
);


CREATE TABLE address (
    address_id SERIAL NOT NULL PRIMARY KEY
    , address_country VARCHAR NOT NULL
    , address_zip   VARCHAR NOT NULL
    , address_state VARCHAR NOT NULL
    , address_city VARCHAR NOT NULL
    , address_line1 VARCHAR NOT NULL
    , address_line2 VARCHAR
);


CREATE TABLE school (
    school_id SERIAL NOT NULL PRIMARY KEY
    , school_name VARCHAR NOT NULL
    , address_id INTEGER NOT NULL
    , CONSTRAINT School_FK_address_id
            FOREIGN KEY(address_id) REFERENCES address(address_id)
);


CREATE TABLE advisor (
    advisor_id SERIAL NOT NULL PRIMARY KEY
    , advisor_name VARCHAR NOT NULL
    , credential_id INTEGER NOT NULL
    , CONSTRAINT Advisor_FK_credential_id
        FOREIGN KEY(credential_id) REFERENCES credential(credential_id)
);


CREATE TABLE school_advisor(
  school_id INTEGER NOT NULL PRIMARY KEY
  , advisor_id INTEGER NOT NULL
  , school_advisor_active BOOLEAN DEFAULT FALSE
  , CONSTRAINT school_advisor_FK_school_id
      FOREIGN KEY(school_id) REFERENCES school(school_id)
  , CONSTRAINT school_advisor_FK_advisor_id
      FOREIGN KEY(advisor_id) REFERENCES advisor(advisor_id)

);


/*
Division:
    A - Advanced
    B - Beginner
*/
CREATE TABLE team (
    team_id SERIAL NOT NULL PRIMARY KEY
  , team_name VARCHAR NOT NULL
  , team_division CHAR(1) NOT NULL
  , school_id INTEGER NOT NULL
  , CONSTRAINT team_FK_school_id
      FOREIGN KEY(school_id) REFERENCES school(school_id)
);


/*
Grade:
1 - Freshmen
2 - Sophomore
3 - Junior
4 - Senior
*/
CREATE TABLE student (
    student_id SERIAL NOT NULL PRIMARY KEY
    , student_name VARCHAR NOT NULL
    , student_grade CHAR(1) NOT NULL
    , school_id INTEGER NOT NULL
    , CONSTRAINT student_FK_school_id
        FOREIGN KEY(school_id) REFERENCES school(school_id)
);

CREATE TABLE TeamStudent (
  team_id INTEGER NOT NULL
  , student_id INTEGER NOT NULL
  , CONSTRAINT teamstudent_FK_team_id
      FOREIGN KEY (team_id) REFERENCES team(team_id)
  , CONSTRAINT teamstudent_FK_student_id
      FOREIGN KEY (student_id) REFERENCES student(student_id)
);

CREATE TABLE team_score (
  team_id INTEGER NOT NULL
  , problem_id INTEGER NOT NULL
  , CONSTRAINT Team_Score_FK_team_id
      FOREIGN KEY (team_id) REFERENCES Team(team_id)
);


/*
TODO: If buses do not need parking validation, do we have to keep track of them in DB?
*/
CREATE TABLE parking (
  parking_id SERIAL NOT NULL PRIMARY KEY
  , vehicle_type CHAR(1) NOT NULL
  , vehicle_count INTEGER NOT NULL
  , validation VARCHAR NOT NULL
  , advisor_id INTEGER
  , CONSTRAINT parking_FK_advisor_id
      FOREIGN KEY(advisor_id) REFERENCES advisor(advisor_id)
);

/*
  Problem and Solution will be stored in MongoDB.
*/


CREATE TABLE Session (
  credential_id SERIAL NOT NULL PRIMARY KEY
  , session_data VARCHAR
  , CONSTRAINT Session_FK_credential_id
    FOREIGN KEY(credential_id) REFERENCES credential(credential_id)
);