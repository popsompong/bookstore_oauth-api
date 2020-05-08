package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// Connect to Cassandra cluster:
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra_db", "init")
}

func GetSession() *gocql.Session {
	return session
}
