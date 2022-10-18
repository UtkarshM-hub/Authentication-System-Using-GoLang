package GoAuth

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(mongoClient mongo.Client,DatabaseName string,collectionName string,data map[string] any,key string) (error) {
	if collectionName==""{
		return errors.New(fmt.Sprint(collectionName,"is not a valid collection name"));
	}
	if _,exist:=data[key];!exist{
		return errors.New(fmt.Sprint(key," is not a valid key"));
	} 
	user:=bson.D{{key,data[key]}}
	result:=mongoClient.Database(DatabaseName).Collection(collectionName).FindOne(context.TODO(),user)
	fmt.Println(result)
	return nil
}
