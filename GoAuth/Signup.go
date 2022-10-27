package GoAuth

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Signup(mongoClient mongo.Client,DatabaseName string,CollectionName string,Data map[string]any,key string) (any,error) {
	// check if the collection name is valid or not
	if CollectionName==""{
		return nil,errors.New(fmt.Sprint(CollectionName,"is not a valid collection name"));
	}
 
	// check if the given property is valid or not
	if _,exist:=Data[key];!exist{
		return nil,errors.New(fmt.Sprint(key," is not a valid key"));
	}

	var result bson.M
	query:=bson.D{primitive.E{Key: key,Value: Data[key]}}
	mongoClient.Database(DatabaseName).Collection(CollectionName).FindOne(context.TODO(),query).Decode(&result)

	if len(result)!=0{
		return nil,errors.New("user already exist")
	}

	input :=bson.D{}

	for k:=range Data{
		input = append(input, primitive.E{Key: k,Value: Data[k]})
	}

	response,err:=mongoClient.Database(DatabaseName).Collection(CollectionName).InsertOne(context.TODO(),input)
	if err!=nil{
		return nil,errors.New("error occured while signup")
	}

	return response,nil
}