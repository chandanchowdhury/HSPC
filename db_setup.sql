/*
Maintain the order as they follow they dependency.
*/
DROP TABLE IF EXISTS parking;
DROP TABLE IF EXISTS team_score;
DROP TABLE IF EXISTS student_team;
DROP TABLE IF EXISTS student;
DROP TABLE IF EXISTS team;
DROP TABLE IF EXISTS school_advisor;
DROP TABLE IF EXISTS advisor;
DROP TABLE IF EXISTS credential;
DROP TABLE IF EXISTS school;
DROP TABLE IF EXISTS address;
DROP TABLE IF EXISTS session;

/**
  Address will hold the address of a School.
 */
CREATE TABLE address (
    address_id SERIAL NOT NULL PRIMARY KEY
    , address_country VARCHAR NOT NULL
    , address_zip   VARCHAR NOT NULL
    , address_state VARCHAR NOT NULL
    , address_city VARCHAR NOT NULL
    , address_line1 VARCHAR NOT NULL
    , address_line2 VARCHAR
    , CONSTRAINT address_unique_all
        UNIQUE (address_country, address_zip, address_state, address_city
                ,address_line1 ,address_line2 )
);

/**
  School details.
 */
CREATE TABLE school (
    school_id SERIAL NOT NULL PRIMARY KEY
    , school_name VARCHAR NOT NULL
    , address_id INTEGER NOT NULL
    , CONSTRAINT School_FK_address_id
            FOREIGN KEY(address_id) REFERENCES address(address_id)
    , CONSTRAINT school_unique_school_name_address
        UNIQUE (school_name, address_id)
);

/**
  A credential is used for authentication purpose.
 */
CREATE TABLE credential (
    credential_id SERIAL NOT NULL PRIMARY KEY
  , emailaddress VARCHAR NOT NULL
  , password_hash VARCHAR
  , credential_active BOOLEAN DEFAULT FALSE
  , CONSTRAINT credential_EmailAddress_unique
      UNIQUE (emailaddress)
);


/**
  An advisor is connected to one or more school and credential which allows
  them to login and make changes to data.
 */
CREATE TABLE advisor (
    advisor_id SERIAL NOT NULL PRIMARY KEY
    , advisor_name VARCHAR NOT NULL
    , credential_id INTEGER NOT NULL
    , CONSTRAINT Advisor_FK_credential_id
        FOREIGN KEY(credential_id) REFERENCES credential(credential_id)
    , CONSTRAINT advisor_unique_name_credential
        UNIQUE(advisor_name, credential_id)
);

/**
  * Connect an advisor to one or more school.
  * One School can have only one advisor.
  * We could have made Advisor part of School, however, keeping them separate is
  easier to maintain.
  * An Advisor can only see and update details of the School she
  is approved as representing. She will not be able to see or update details
  of another school.
 */
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
  A Team represents a School.
  There are limit on how many Teams a School can have but the check done in
  application.

  Division:
    A - Advanced
    B - Beginner
*/
CREATE TABLE team (
    team_id SERIAL NOT NULL PRIMARY KEY
  , team_name VARCHAR NOT NULL
  , team_division VARCHAR NOT NULL
  , school_id INTEGER NOT NULL
  , CONSTRAINT team_FK_school_id
      FOREIGN KEY(school_id) REFERENCES school(school_id)
  , CONSTRAINT team_unique_name_school
      UNIQUE (team_name, school_id)
);


/*
  A Student is part of a School.

  Grade:
    1 - Freshmen
    2 - Sophomore
    3 - Junior
    4 - Senior
*/
CREATE TABLE student (
    student_id SERIAL NOT NULL PRIMARY KEY
    , student_name VARCHAR NOT NULL
    , student_grade VARCHAR NOT NULL
    , school_id INTEGER NOT NULL
    , CONSTRAINT student_FK_school_id
        FOREIGN KEY(school_id) REFERENCES school(school_id)
    , CONSTRAINT student_unique_name_grade_school
        UNIQUE (student_name, student_grade, school_id)
);

/**
  * One team can have one or more Students.
  * However, one Student is part of only one Team.
  * Maintenance is easier if we store Student-Team relation separately.
  * BUT in application logic, we have to make sure the school of the student and
  school of the team are same. We cannot allow a student from one school
  to be added to a team from another school.
  * We could have created a trigger, but application logic check are easier to
  code and debug.
 */
CREATE TABLE student_team (
  -- Student can be part of only one Team
  student_id INTEGER PRIMARY KEY
  , team_id INTEGER NOT NULL
  , CONSTRAINT teamstudent_FK_student_id
      FOREIGN KEY (student_id) REFERENCES student(student_id)
  , CONSTRAINT teamstudent_FK_team_id
      FOREIGN KEY (team_id) REFERENCES team(team_id)
);

/**
  Record when a Team has successfully solved a Problem.
  The Problems are stored in MongoDB, so the check is done in application logic.
 */
CREATE TABLE team_score (
  team_id INTEGER NOT NULL
  , problem_id INTEGER NOT NULL
  , score_earned INTEGER NOT NULL
  , submit_ts TIMESTAMP NOT NULL
  , CONSTRAINT Team_Score_FK_team_id
      FOREIGN KEY (team_id) REFERENCES Team(team_id)
  -- A Problem can be solved by a team only once
  , UNIQUE (team_id, problem_id)
);


/**
  Parking permits to allow vehicals to be parked in the Garage.
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
