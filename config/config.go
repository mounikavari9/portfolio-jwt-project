//Because every other file needs the JWT secret.

//This file loads .env and reads JWT_SECRET


package config

import(
	"log"
	"os" //used to read environment variables
	"github.com/joho/godotenv" //third-party package that reads the .env file and loads its values into environment variables
)


//This function loads values from the .env file into the apllication's environment variable
//this function is called once when the application starts
func LoadConfig(){

	//to load the .env file
	err := godotenv.Load()
	if err!= nil{
		log.Println("No .env file found, using system environment variables")
	}	
}

//Every time we generate or validate a jwt,
//we call this function to get the secret 
func GetJWTSecret() []byte {

	//reads JW_SECRET from the environment file
	secret := os.Getenv("JWT_SECRET")
	if secret == ""{
		log.Fatal("JWT_SECRET is not set")
	}

	//convert the strig into byte slice 
	//The JWT library expects the signing key as []byte not as a string
	return []byte(secret)
}


