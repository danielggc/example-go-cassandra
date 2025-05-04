package infrastruture

import (
	"broker/cmd/api/dto"
	"fmt"

	"github.com/gocql/gocql"
	. "github.com/samber/mo"
)

type CassandraUserRepository struct {
	session *gocql.Session
}

func (t *CassandraUserRepository) GetUserById(uuid string) Result[Option[dto.User]] {
	var user dto.User
	query := `SELECT * FROM users WHERE id = ? LIMIT 1`
	iter := t.session.Query(query, uuid).Iter()
	err := iter.Scan(&user.Id, &user.Name, &user.Email)
	if !err {
		fmt.Println("hello error ->>>>>>>")
		if err := iter.Close(); err != nil {
			return Err[Option[dto.User]](err)
		}
		return Err[Option[dto.User]](fmt.Errorf("error no found  user"))
	}
	return Ok(Some(user))
}

func (t *CassandraUserRepository) SaveUser(user dto.User) Result[Option[dto.User]] {
	query := `INSERT INTO users (id, name, email) VALUES (?, ?, ?)`
	err := t.session.Query(query, user.Id, user.Name, user.Email).Exec()
	if err != nil {
		return Err[Option[dto.User]](err)
	}
	return Ok(Some(user))
}
