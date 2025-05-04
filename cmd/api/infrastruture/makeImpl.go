package infrastruture

import (
	"broker/cmd/api/infrastruture/cassandra"

	"github.com/gocql/gocql"
)

func CreateRepository(session *gocql.Session) cassandra.UserRepository {
	return &CassandraUserRepository{session: session}
}
