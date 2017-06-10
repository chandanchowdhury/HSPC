package dbhandler

import (
	"github.com/chandanchowdhury/HSPC/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"
)

const (
	Session_Collection_Name = "Session"
)

func getSessionColl() *mgo.Collection {
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
	session_coll_exists := false
	for _, c := range colls {
		if c == Session_Collection_Name {
			session_coll_exists = true
		}
	}

	session_coll := hspc_DB.C(Session_Collection_Name)

	if !session_coll_exists {
		index := mgo.Index{
			Key:        []string{"email"},
			Unique:     true,
			DropDups:   true,
			Background: true, // See notes.
			Sparse:     true,
		}
		err := session_coll.EnsureIndex(index)

		if err != nil {
			log.Panic("Error creating index")
		}
	}

	return session_coll
}

func SessionCreate(principal *models.Principal) models.Error {
	log.Printf("Creating Session for Email = %s", principal.Email)

	error := models.Error{}

	coll := getSessionColl()

	err := coll.Insert(principal)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			error.Message = "Duplicate Session"
			return error
		}

		log.Print(err)
		error.Message = err.Error()
		return error
	}

	error.Code = 0

	return error
}

func SessionRead(email string) *models.Principal {
	log.Printf("Reading Session for email = %s", email)

	query := getSessionColl().Find(bson.M{"email": email})

	result_count, err := query.Count()
	log.Printf("Session Found: %d", result_count)
	if err != nil {
		log.Panic(err)
	}

	if result_count < 1 {
		log.Print("Problem not found")
		return &models.Principal{}
	}

	if result_count > 1 {
		log.Print("Unexpected number of Problems found")
	}

	//fetch the data
	var principal = models.Principal{}
	query.One(&principal)

	return &principal
}

func sessionUpdate(principal models.Principal) int64 {
	//TODO: complete the logic

	return 0
}

func SessionDelete(email string) int64 {
	log.Printf("Deleting Session for Email = %s", email)

	err := getSessionColl().Remove(bson.M{"email": email})

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Print("Session not found")
			return -2
		}

		log.Print("Failed deleting Session")
		log.Panic(err)
		return -1
	}

	log.Print("Session deleted")
	return 1
}
