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

	admindb := session.DB("admin")

	err = admindb.Login(MONGO_DB_USER, MONGO_DB_PASSWORD)

	if err != nil {
		log.Panic(err)
	}

	hspc_DB := session.DB("HSPC")

	return hspc_DB.C(Solution_Collection_Name)
}

/*
Solution
*/
func SolutionCreate(solution models.Solution) int64 {
	log.Print("# Creating Solution")
	log.Printf("Solution ID = %d", solution.SolutionID)

	coll := getSolutionColl()

	//TODO: How to automatically set _id which is different from solutionid

	err := coll.Insert(solution)

	if err != nil {
		log.Print(err)
		return -1
	}

	//TODO: return the ID of the newly inserted Solution
	return solution.SolutionID
}

func SolutionRead(solution_id int64) models.Solution {
	log.Print("# Reading Solution")
	log.Printf("Solution ID = %d", solution_id)

	//query := getSolutionColl().FindId(solution_id)
	query := getSolutionColl().Find(bson.M{"solutionid": solution_id})

	result_count, err := query.Count()
	log.Printf("Solution Found: %d", result_count)
	if err != nil {
		log.Print(err)
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

func SolutionUpdate(solution models.Solution) bool {
	log.Print("# Updating Solution")
	log.Printf("Solution ID = %d", solution.SolutionID)

	err := getSolutionColl().Update(bson.M{"solutionid": solution.SolutionID}, solution)

	if err != nil {
		log.Print("Failed updating Solution")
		log.Print(err)
		return false
	}

	return true
}

func SolutionDelete(solution_id int64) bool {
	log.Print("# Deleting Solution")
	log.Printf("Solution ID = %d", solution_id)

	err := getSolutionColl().Remove(bson.M{"solutionid": solution_id})

	if err != nil {
		log.Print("Failed deleting Solution")
		log.Print(err)
		return false
	}

	return true
}
