package GoAuth

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func IfExist(mongoClient mongo.Client,DatabaseName string,CollectionName string,key string,value any)(error){

	// Check whether the input key is empty or not
	if len(key)==0{
		return errors.New("key or the value provided is invalid")
	}
	
	//query to find whether the object already exist in the database or not 	
	var result bson.M
	query:=bson.D{primitive.E{Key: key,Value: value}}
	mongoClient.Database(DatabaseName).Collection(CollectionName).FindOne(context.TODO(),query).Decode(&result)

	// return if there are no objects in database
	if len(result)!=0{
		return errors.New("user already exist")
	}

	return nil
}
