package dbhandler

type credential_struct struct {
    credential_id int64
    emailaddress string
    password_hash string
}

type address_struct struct {
    address_id int64
    country string
    zipcode string
    state string
    city string
    line1 string
    line2 string
}

type school_struct struct {
    school_id int64
    school_name string
    address_id int64
}

type advisor_struct struct {
    advisor_id int64
    advisor_name string
    credential_id int32
}

type school_advisor_struct struct {
    advisor_id int64
    school_id int32
}

type team_struct struct {
    team_id int64
    team_name string
    team_division string
    advisor_id int64
}

type student_struct struct {
    student_id int64
    student_name string
    student_grade string
    school_id int64
    team_id int64
}

type team_score_struct struct {
    team_id int64
    problem_id int64
}

type parking struct {
    parking_id int64
    vehicle_count int64
    validation string
    advisor_id int64
}

type problem_struct struct {
    problem_id int64
    round int64
    division int64
    problem_desc string
}

type solution_struct struct {
    solution_id int64
    solution_desc string
}

type problem_solution_struct struct {
    problem_id int64
    solution_id int64
}