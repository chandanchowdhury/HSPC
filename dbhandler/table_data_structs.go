package dbhandler

type credential_struct struct {
    credential_id uint32
    emailaddress string
    password_hash string
}

type address_struct struct {
    address_id uint32
    country string
    zipcode string
    state string
    city string
    line1 string
    line2 string
}

type school_struct struct {
    school_id uint32
    school_name string
    address_id uint32
}

type advisor_struct struct {
    advisor_id uint32
    advisor_name string
    credential_id int32
}

type school_advisor_struct struct {
    advisor_id uint32
    school_id int32
}

type team_struct struct {
    team_id uint32
    team_name string
    team_division string
    advisor_id uint32
}

type student_struct struct {
    student_id uint32
    student_name string
    student_grade string
    school_id uint32
    team_id uint32
}

type team_score_struct struct {
    team_id uint32
    problem_id uint32
}

type parking struct {
    parking_id uint32

}

type problem_struct struct {
    problem_id uint32
    round uint32
    division uint32
    problem_desc string
}

type solution_struct struct {
    solution_id uint32
    solution_desc string
}

type problem_solution struct {
    problem_id uint32
    solution_id uint32
}