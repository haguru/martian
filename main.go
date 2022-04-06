package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

type DBObj struct {
	Id   gocql.UUID
	Text string
}

func main() {
	// fmt.Println("WIP")
	// TODO: TAKE A LOOK AT GOCQLX FOR STRACT TAGS
	// api needs to wait a bit for cassandra to come up thus, retries need to implemented
	cluster := gocql.NewCluster("172.17.0.3")
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	// create keyspace
	if err := session.Query("create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }").Exec(); err != nil {
		fmt.Println("Error while inserting example")
		fmt.Println(err)
	}
	if err := session.Query("create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id))").Exec(); err != nil {
		fmt.Println("Error while creating table example.tweet")
		fmt.Println(err)
	}
	if err := session.Query("create index on example.tweet(timeline)").Exec(); err != nil {
		fmt.Println("Error while creating index")
		fmt.Println(err)
	}

	// create index on example.tweet(timeline);
	session.Close()
	// SELECT THE CREATED KEYSPACE
	cluster.Keyspace = "example"

	cluster.Consistency = gocql.Quorum
	session2, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session2.Close()

	if err = session2.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
		log.Fatal(err)
	}

	var id gocql.UUID
	var text string
	// testVar := &DBObj{}

	if err := session2.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
		"me").Consistency(gocql.One).Scan(&id, &text); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tweet:", id, text)

	iter := session2.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
	for iter.Scan(&id, &text) {
		fmt.Println("Tweet:", id, text)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
