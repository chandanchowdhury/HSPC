package dbhandler

import (
	"github.com/chandanchowdhury/HSPC/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//TODO: Read from config file
const (
	MONGO_DB_HOST           = "localhost"
	MONGO_DB_USER           = "hspc"
	MONGO_DB_PASSWORD       = "HSPC-Password"
	Problem_Collection_Name = "Problems"
)

func getProblemColl() *mgo.Collection {
	session, err := mgo.Dial(MONGO_DB_HOST)

	admindb := session.DB("admin")

	err = admindb.Login(MONGO_DB_USER, MONGO_DB_PASSWORD)

	if err != nil {
		log.Panic(err)
	}

	hspc_DB := session.DB("HSPC")

	return hspc_DB.C(Problem_Collection_Name)
}

/*
Problem
*/
func ProblemCreate(problem models.Problem) int64 {
	log.Print("# Creating Problem")
	log.Printf("Problem ID = %d", problem.ProblemID)

	coll := getProblemColl()

	//TODO: How to automatically set _id which is different from problemid
	//problem_json, err := bson.Marshal(problem)
	//log.Print(string(problem_json))

	err := coll.Insert(problem)

	if err != nil {
		log.Print(err)
		return -1
	}

	//TODO: return the ID of the newly inserted Problem
	return problem.ProblemID
}

func ProblemRead(problem_id int64) models.Problem {
	log.Print("# Reading Problem")
	log.Printf("Problem ID = %d", problem_id)

	//query := getProblemColl().FindId(problem_id)
	query := getProblemColl().Find(bson.M{"problemid": problem_id})

	result_count, err := query.Count()
	log.Printf("Problem Found: %d", result_count)
	if err != nil {
		log.Print(err)
		return models.Problem{}
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

func ProblemUpdate(problem models.Problem) bool {
	log.Print("# Updating Problem")
	log.Printf("Problem ID = %d", problem.ProblemID)

	err := getProblemColl().Update(bson.M{"problemid": problem.ProblemID}, problem)

	if err != nil {
		log.Print("Failed updating Problem")
		log.Print(err)
		return false
	}

	return true
}

func ProblemDelete(problem_id int64) bool {
	log.Print("# Deleting Problem")
	log.Printf("Problem ID = %d", problem_id)

	err := getProblemColl().Remove(bson.M{"problemid": problem_id})

	if err != nil {
		log.Print("Failed deleting Problem")
		log.Print(err)
		return false
	}

	return true
}

func ProblemReadList() []*models.Problem {
	log.Print("# Reading Problem List")

	query := getProblemColl().Find(bson.M{})

	result_count, err := query.Count()
	log.Printf("Problem Found: %d", result_count)
	if err != nil {
		log.Print(err)
		return []*models.Problem{}
	}

	if result_count < 1 {
		log.Print("Problem not found")
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
