// this file generates JWT token after user successfully logs in
//It doesnot validate username/password, receive HTTP requests, Verify JWT tokens, it ONLY creates JWT token

package auth

import(
	"github.com/golang-jwt/jwt/v5"
	"time"

	//use it to get the JWT secret from .env file
	"github.com/mounikavari9/portfolio-jwt-project/config"
)

//claims is the information that will be stored inside the JWT
//Think claims as the payload of the token
type Claims struct{
	//store the logged in user's ID
	UserID 		string 		`json:"user_id"`

	//store the logged in user's role
	Role 		string		`json:"role"`

	//embedded struct that provides standard JWT fields such as
	//Expiry time, Issued Time, Not before, Issuer 
	jwt.RegisteredClaims 
}

//Creating GenerateJWT function take input as logged in user's userID, role and returns JWT token as string and error if any
func GenerateJWT(userID string, role string)(string, error) {

	claims := Claims{
		UserID: userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExperiesAt: jwt.NewNumericDate(
				time.Now().Add(1*time.Hour),
			),
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}

	//create a new JWT token
	//HS256 is the signing algorithm
	//the claims object becomes the payload of the jwt
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	//sign the token using secret key
	//the secret key comes from the config package, which reads JWT_SECRET from .env file
	//finally return the signed JWT token
	return token.SignedString(
		config.GetJWTSecret(),
	)

}



