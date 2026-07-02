//Middleware validate token, check expiry, extract UserID, extract Role
//Middleware runs before the actual API handler

package middleware

import (
	"net/http"
	"context" //used to store values (UserID, Role) insode the request context
	"strings" //used to remove "Bearer" from Authorization header
	"github.com/golang-jwt/jwt/v5"
	"github.com/mounikavari9/portfolio-jwt-project/auth"  //import custom Claims struct
	"github.com/mounikavari9/portfolio-jwt-project/config" //used to read JWT secret from .env
)

//Create a custom type for context keys
//Using our own Type prevents key collisions with other packages
type contextKey string 

//keys used to store values inside request context
const UserIDKey contextKey = "userID"
const RoleKey contextKey = "role"

//Every request passes through this function before reaching the handler
//takes input paraneter as next and type http.Handler and returns new handler with authentication
func JWTMiddleware(next http.Handler) http.Handler {

	//return another http handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		//read authorization header
		authHeader := r.Header.Get("Authorization")

		//if authorization header is missing, user is not authenticated
		if authHeader == ""{
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return 
		}

		//remove Bearer from the authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer")

		//create an empty Claims object
		//why empty Claims object: Because the JWT library needs a place to put the data after it reads the token.
		claims := &auth.Claims{}

		//parse and validate the JWT
		token, err := jwt.ParseWithClaims(

			//jwt string
			tokenString,
			//store decoded values inside claims
			claims,

			//provide secret key used for verification
			func(token *jwt.Token) (interface{}, error){
				return config.GetJWTSecret(), nil 
			},
		)

		//Token invalid?, signature wrong?, expired?
		if err != nil || !token.Valid{
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return 
		}

		//store authenticated user's ID inside request context
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)

		//store authenticated user's role inside request context
		ctx = context.WithValue(ctx, RoleKey, claims.Role)

		//continue to next handler
		//the handler can now read UserID and Role from the request context
		next.ServeHTTP(w, r.WithContext(ctx))


	})


}
