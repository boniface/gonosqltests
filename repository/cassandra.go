package repository

import (
	"fmt"
	"github.com/boniface/gonosqltests/domain"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func getCluster() *gocql.Session {
	var err error
	cluster := gocql.NewCluster("localhost")
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

func update(person domain.Person) domain.Person {
	fmt.Println(" **** Updating   person ****\n", person)
	updatedPerson := domain.Person{ID: person.ID, Name: person.Name + " Updated"}
	defer getCluster().Close()
	if err := Session.Query("UPDATE persons SET name = ? WHERE empid = ?",
		updatedPerson.ID, updatedPerson.Name).Exec(); err != nil {
		fmt.Println("Error while updating Emp")
		fmt.Println(err)
	}
	return updatedPerson

}

func delete(person domain.Person) bool {
	fmt.Println(" **** Deleting Person  ****\n", person)
	defer getCluster().Close()
	if err := Session.Query("DELETE FROM persons WHERE id = ?", person.ID).Exec(); err != nil {
		fmt.Println("Error while deleting Emp")
		fmt.Println(err)
	}
	return true

}

func read(id string) domain.Person {
	fmt.Println(" **** Reading a person with ID ****\n", id)
	defer getCluster().Close()
	var person domain.Person
	m := map[string]interface{}{}
	iter := getCluster().Query("SELECT * FROM persons WHERE id=? LIMIT 1").Iter()
	for iter.MapScan(m) {
		person = domain.Person{
			ID:   m["id"].(string),
			Name: m["name"].(string),
		}
		m = map[string]interface{}{}
	}
	return person

}

func readAll() []domain.Person {
	fmt.Println(" **** Reading All Records ****\n")
	defer getCluster().Close()
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
