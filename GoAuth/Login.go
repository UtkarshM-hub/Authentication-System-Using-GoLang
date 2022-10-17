package GoAuth

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func Login[T map[string] any](mongoClient mongo.Client,collectionName string,data T,key string) (error) {
	if collectionName==""{
		return errors.New(fmt.Sprint(collectionName,"is not a valid collection name"))
	}
	if _,exist:=data[key];!exist{
		return errors.New(fmt.Sprint(key,"is not a valid key"))
	} else{
		fmt.Println(data[key]);
	}
	return nil
}
