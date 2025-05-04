package cassandra

import (
	"github.com/gocql/gocql"
	. "github.com/samber/mo"
)

func ConnectDB(url string, keyspace string) Result[*gocql.Session] {
	cluster := gocql.NewCluster(url)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	cluster.Port = 9042
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandra",
	}
	maybeSession := Try(func() (*gocql.Session, error) {
		return cluster.CreateSession()
	})
	return maybeSession

}
