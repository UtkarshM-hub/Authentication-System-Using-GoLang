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

	// Check if the document already exist in database or not
	err:=IfExist(mongoClient,DatabaseName,CollectionName,key,Data[key]);
	if err!=nil{
		return nil,err
	}

	// get input data in bson format
	input :=bson.D{}
	for k:=range Data{
		input = append(input, primitive.E{Key: k,Value: Data[k]})
	}

	// Insert the data
	response,err:=mongoClient.Database(DatabaseName).Collection(CollectionName).InsertOne(context.TODO(),input)
	if err!=nil{
		return nil,errors.New("error occured while signup")
	}

	return response,nil
}