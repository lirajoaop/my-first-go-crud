package converter

import (
	"github.com/lirajoaop/my-first-go-crud/src/model"
	"github.com/lirajoaop/my-first-go-crud/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)
	domain.SetID(entity.ID.Hex())
	return domain
}
