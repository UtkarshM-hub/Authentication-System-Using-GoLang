package GoAuth

import (
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func Connect() (*mongo.Client,error) {
	godotenv.Load()
	
	client,err:=mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	// Showing Error if occured
	if err!=nil{
		return nil,err
	}

	// Setting a time limit so that the operation will not take more time than mentioned
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)

	// Trying to connect
	err=client.Connect(ctx);

	// cancel releases the memory taken by temporary values
	// and defer runs it after the surrounding function returns
	defer cancel()

	if err!=nil {
		return nil,err
	}

	return client,nil
}