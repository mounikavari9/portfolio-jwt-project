package main

import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/mounikavari9/portfolio-jwt-project/middleware"
	"github.com/mounikavari9/portfolio-jwt-project/auth"  //import custom Claims struct
	"github.com/mounikavari9/portfolio-jwt-project/config"
	"github.com/mounikavari9/portfolio-jwt-project/handler"
)


//this function generally read username/password, check db and verify password
//for this demo we directly generate a JWT 
func loginHandler(w http.ResponseWriter, r *http.Request){
	//generate JWT for a demo user
	token, err := auth.GenerateJWT("USER123", "USER")
	if err!= nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return 
	}

	//tell client the repose is JSON
	w.Header().Set("Content-Type", "application/json")

	//send the JWT token back to the client
	json.NewEncoder(w).Encode(map[string]string{"token": token})

}

func main(){
	
	//load environment variables from .env file
	config.LoadConfig()

	//Register Login API and for GET/login calls loginHandler()
	http.HandleFunc("/login", loginHandler)

	//convert GetPortfolioSummary function into a http handler
	portfolioHandler := http.HandlerFunc(
		handler.GetPortfolioSummary,
	)

	//Register protected portfolio API 
	http.Handle("/portfolio/summary", middleware.JWTMiddleware(portfolioHandler))

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}