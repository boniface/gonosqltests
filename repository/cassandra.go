package repository

import (
	"fmt"
	"github.com/boniface/gonosqltests/domain"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func getCluster() *gocql.Session {
	var err error
	cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.3")
	cluster.Keyspace = "nosqltest"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return Session
}

func create(person domain.Person) domain.Person {
	fmt.Println(" **** Creating new person ****\n", person)
	defer getCluster().Close()

	if err := getCluster().Query("INSERT INTO persons(id, name) VALUES(?, ?)",
		person.ID, person.Name).Exec(); err != nil {
		fmt.Println("Error while inserting Emp")
		fmt.Println(err)
	}
	return person
}

func upodate(person domain.Person) domain.Person {
	fmt.Println(" **** Creating new person ****\n", person)

}

func delete(person domain.Person) bool {

}

func read(id string) domain.Person {

}

func readAll() []domain.Person {
	var persons []domain.Person
	m := map[string]interface{}{}
	defer getCluster().Close()
	iter := getCluster().Query("SELECT * FROM persons").Iter()
	for iter.MapScan(m) {
		persons = append(persons, domain.Person{
			ID:   m["id"].(string),
			Name: m["name"].(string),
		})
		m = map[string]interface{}{}
	}
	return persons
}
