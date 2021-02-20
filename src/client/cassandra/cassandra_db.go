package cassandra

import (
	"github.com/gocql/gocql"
)
var (
	cluster *gocql.ClusterConfig
	session *gocql.Session
)
func init(){
	// connect to the cassandra cluster
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}
func GetSession() *gocql.Session{
	return session
}
