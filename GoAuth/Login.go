package GoAuth

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(mongoClient mongo.Client,DatabaseName string,collectionName string,data map[string] any,key string) (error) {

	// check if the collection name is valid or not
	if collectionName==""{
		return errors.New(fmt.Sprint(collectionName,"is not a valid collection name"));
	}

	// check if the given property is valid or not
	if _,exist:=data[key];!exist{
		return errors.New(fmt.Sprint(key," is not a valid key"));
	} 

	var result bson.M

	// create a query out of user data
	user:=bson.D{primitive.E{Key:key,Value:data[key]}}

	// find the user in database
	if err:=mongoClient.Database(DatabaseName).Collection(collectionName).FindOne(context.TODO(),user).Decode(&result); err!=nil{
		// return error if there is no data in database
		return errors.New("user does not exist");
	}

	// return error if the user does not exist
	if(len(result)==0){
		return errors.New("user does not exist")
	}
	return nil
}
