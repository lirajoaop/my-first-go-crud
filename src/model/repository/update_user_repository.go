package repository

import (
	"context"
	"os"

	"github.com/lirajoaop/my-first-go-crud/src/configuration/logger"
	"github.com/lirajoaop/my-first-go-crud/src/configuration/rest_err"
	"github.com/lirajoaop/my-first-go-crud/src/model"
	"github.com/lirajoaop/my-first-go-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init updateUser repository",
		zap.String("journey", "updateUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	filter := bson.M{"_id": userId}

	_, err := collection.UpdateOne(context.Background(), filter, value)
	if err != nil {
		logger.Error("Error trying to update user",
			err,
			zap.String("journey", "updateUser"),
		)
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"updateUser repository executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "updateUser"),
	)

	return nil

}
