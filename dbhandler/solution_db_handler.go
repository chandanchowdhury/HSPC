package dbhandler

import (
	"github.com/chandanchowdhury/HSPC/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	Solution_Collection_Name = "Solutions"
)

func getSolutionColl() *mgo.Collection {
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
	solution_coll_exists := false
	for _, c := range colls {
		if c == Solution_Collection_Name {
			solution_coll_exists = true
		}
	}

	solution_coll := hspc_DB.C(Solution_Collection_Name)

	if !solution_coll_exists {
		index := mgo.Index{
			Key:        []string{"solutionid"},
			Unique:     true,
			DropDups:   true,
			Background: true, // See notes.
			Sparse:     true,
		}
		err := solution_coll.EnsureIndex(index)

		if err != nil {
			log.Panic("Error creating index")
		}
	}

	return solution_coll
}

/*
Solution
*/
func SolutionCreate(solution models.Solution) int64 {
	log.Printf("Creating Solution ID = %d for Problem ID = %d", *solution.SolutionID, *solution.ProblemID)

	//make sure the Problem exists
	problem_data := ProblemRead(*solution.ProblemID)

	if problem_data.ProblemID == nil {
		log.Print("ProblemID does not exists")
		return -2
	}

	coll := getSolutionColl()

	err := coll.Insert(solution)

	if err != nil {
		log.Panic(err)
		return -1
	}

	return *solution.SolutionID
}

func SolutionRead(solution_id int64) models.Solution {
	log.Printf("Reading Solution ID = %d", solution_id)

	query := getSolutionColl().Find(bson.M{"solutionid": solution_id})

	result_count, err := query.Count()
	log.Printf("Solution Found: %d", result_count)
	if err != nil {
		log.Panic(err)
		return models.Solution{}
	}

	if result_count < 1 {
		log.Print("Solution not found")
		return models.Solution{}
	}

	if result_count > 1 {
		log.Print("Unexpected number of Solutions found")
	}

	//fetch the Solution details
	var solution models.Solution
	query.One(&solution)

	return solution
}

func SolutionUpdate(solution models.Solution) int64 {
	log.Printf("Updating Solution ID = %d", solution.SolutionID)

	//make sure the Problem exists
	problem_data := ProblemRead(*solution.ProblemID)

	if *problem_data.ProblemID == 0 {
		log.Print("Problem does not exists for the Solution")
		return -2
	}

	err := getSolutionColl().Update(bson.M{"solutionid": solution.SolutionID}, solution)

	if err != nil {
		log.Print("Failed updating Solution")
		log.Panic(err)
		return -1
	}

	return 1
}

func SolutionDelete(solution_id int64) int64 {
	log.Printf("Deleting Solution ID = %d", solution_id)

	err := getSolutionColl().Remove(bson.M{"solutionid": solution_id})

	if err != nil {
		log.Print("Failed deleting Solution")
		log.Panic(err)
		return -1
	}

	return 1
}

/**
Given an ProblemID find all related solutions.
*/
func SolutionForProblem(problem_id int64) []*models.Solution {
	log.Printf("Reading Solutions for Problem ID = %d", problem_id)

	query := getSolutionColl().Find(bson.M{"problemid": problem_id})

	result_count, err := query.Count()
	log.Printf("Solution Found: %d", result_count)
	if err != nil {
		log.Panic(err)
		return []*models.Solution{}
	}

	//create the array to hold the list of solutions
	solution_list := make([]*models.Solution, 0)

	//Get the Iterator for the results
	i := query.Iter()
	//loop through the solutions
	for !i.Done() {
		solution := new(models.Solution)

		//fetch the solution details
		i.Next(&solution)

		solution_list = append(solution_list, solution)
	}

	return solution_list

}
