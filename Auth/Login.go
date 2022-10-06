package auth

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Login() string {
	godotenv.Load()
	fmt.Println(os.Getenv("JUST"))
	return "you have called login function"
}
