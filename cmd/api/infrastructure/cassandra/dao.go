package cassandra

import (
	"broker/cmd/api/dto"

	. "github.com/samber/mo"
)

type UserRepository interface {
	SaveUser(dto.User) Result[Option[dto.User]]
	GetUserById(uuid string) Result[Option[dto.User]]
}
