package dbhandler

import (
	"log"

	"github.com/chandanchowdhury/HSPC/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

//TODO: Read from config file
const (
	MONGO_DB_HOST           = "localhost"
	MONGO_DB_USER           = "hspc"
	MONGO_AUTH_DB           = "admin"
	MONGO_DB_PASSWORD       = "HSPC-Password"
	Problem_Collection_Name = "Problems"
)

func getProblemColl() *mgo.Collection {
	session, err := mgo.Dial(MONGO_DB_HOST)

	if err != nil {
		log.Fatalf("Error Connecting to MongoDB: %s", err.Error())
	}

	admindb := session.DB(MONGO_AUTH_DB)

	err = admindb.Login(MONGO_DB_USER, MONGO_DB_PASSWORD)

	if err != nil {
		log.Panic(err)
	}

	hspc_DB := session.DB("HSPC")

	//if the collection not already exits
	colls, err := hspc_DB.CollectionNames()
	problem_coll_exists := false
	for _, c := range colls {
		if c == Problem_Collection_Name {
			problem_coll_exists = true
		}
	}

	problem_coll := hspc_DB.C(Problem_Collection_Name)

	if !problem_coll_exists {
		index := mgo.Index{
			Key:        []string{"problemid"},
			Unique:     true,
			DropDups:   true,
			Background: true, // See notes.
			Sparse:     true,
		}
		err := problem_coll.EnsureIndex(index)

		if err != nil {
			log.Panic("Error creating index")
		}
	}

	return problem_coll
}

/*
Problem

For experiment purpose, create return an model.Error instead of error code. We want to see if that helps
in reducing the logic at requesthandler side.
*/
func ProblemCreate(problem models.Problem) models.Error {
	log.Printf("Creating Problem ID = %d", *problem.ProblemID)

	error := models.Error{}

	coll := getProblemColl()

	err := coll.Insert(problem)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			error.Message = "Duplicate ProblemID"
			return error
		}

		log.Panic(err)
		error.Message = err.Error()
		return error
	}

	error.Code = *problem.ProblemID

	return error
}

func ProblemRead(problem_id int64) models.Problem {
	log.Printf("Reading Problem ID = %d", problem_id)

	query := getProblemColl().Find(bson.M{"problemid": problem_id})

	result_count, err := query.Count()
	log.Printf("Problem Found: %d", result_count)
	if err != nil {
		log.Panic(err)
	}

	if result_count < 1 {
		log.Print("Problem not found")
		return models.Problem{}
	}

	if result_count > 1 {
		log.Print("Unexpected number of Problems found")
	}

	//fetch the Problem details
	var problem models.Problem
	query.One(&problem)

	return problem
}

func ProblemUpdate(problem models.Problem) int64 {
	log.Printf("Updating Problem ID = %d", *problem.ProblemID)

	err := getProblemColl().Update(bson.M{"problemid": problem.ProblemID}, problem)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return -2
		}

		log.Print("Failed updating Problem")
		log.Panic(err)
		return -1
	}

	return 1
}

func ProblemDelete(problem_id int64) int64 {
	log.Printf("Deleting Problem ID = %d", problem_id)

	//make sure no Solution exists
	solution_list := SolutionForProblem(problem_id)

	if len(solution_list) > 0 {
		return -3
	}

	err := getProblemColl().Remove(bson.M{"problemid": problem_id})

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return -2
		}

		log.Print("Failed deleting Problem")
		log.Panic(err)
		return -1
	}

	return 1
}

func ProblemReadList() []*models.Problem {
	log.Print("Reading Problem List")

	query := getProblemColl().Find(bson.M{})

	result_count, err := query.Count()
	log.Printf("Problem Found: %d", result_count)
	if err != nil {
		log.Panic(err)
		return []*models.Problem{}
	}

	if result_count < 1 {
		log.Print("Empty Problem List")
		return []*models.Problem{}
	}

	//create the array to hold the list of Problems
	problem_list := make([]*models.Problem, 0)

	//Get the Iterator for the results
	i := query.Iter()
	//loop through the Problems
	for !i.Done() {
		problem := new(models.Problem)

		//fetch the Problem details
		i.Next(&problem)

		problem_list = append(problem_list, problem)
	}

	return problem_list
}
