package GoAuth

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Signup(mongoClient mongo.Client,DatabaseName string,CollectionName string,Data map[string]any,key string) error {
	// check if the collection name is valid or not
	if CollectionName==""{
		return errors.New(fmt.Sprint(CollectionName,"is not a valid collection name"));
	}
 
	// check if the given property is valid or not
	if _,exist:=Data[key];!exist{
		return errors.New(fmt.Sprint(key," is not a valid key"));
	}

	err:=IfExist(mongoClient,DatabaseName,CollectionName,key,Data[key]);
	if err!=nil{
		return err
	}

	input :=bson.D{}

	for k:=range Data{
		input = append(input, primitive.E{Key: k,Value: Data[k]})
	}
	fmt.Println(input)
	response,err:=mongoClient.Database(DatabaseName).Collection(CollectionName).InsertOne(context.TODO(),input)
	if err!=nil{
		return errors.New("error occured while signup")
	}

	fmt.Println(response)

	return nil
}