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
	if collectionName==""{
		return errors.New(fmt.Sprint(collectionName,"is not a valid collection name"));
	}
	if _,exist:=data[key];!exist{
		return errors.New(fmt.Sprint(key," is not a valid key"));
	} 
	var result=struct{
		name string
	}{}
	// user:=bson.D{primitive.E{Key:key,Value:data[key]}}
	mongoClient.Database("GoAuth").Collection("user").FindOne(context.TODO(),bson.D{primitive.E{Key: "name",Value: "UtkarshM-hub"}}).Decode(&result)
	// mongoClient.Database(DatabaseName).Collection(collectionName).FindOne(context.TODO(),user).Decode(&result)
	fmt.Println(result)
	return nil
}
