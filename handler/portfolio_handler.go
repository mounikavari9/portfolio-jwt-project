package handler

import(
	"encoding/json"
	"net/http"
	""
)

type PortfolioSummary struct {
	UserID 	string 		`json:"user_id"`
	Role 	string		`json:"role"`
	TotalCash	float64 	`json:"total_cash"`
	Message 	string 		`json:"message"`
}

//this function is called only after JWT middleware successfully validates the user's token
func GetPortfolioSummary(w http.ResponseWriter, r *http.Request){
	//Read UserID from request context
	//JWT middleware stored it here after validating the JWT token
	userID := r.Context().Value(middleware.UserIDKey).(string)

	//read role from the request context
	role := r.Context().Value(middleware.RoleKey).(string)

	response := PortfolioSummary{
		UserID : userID,
		Role: role,
		TotalCash: 2500.50,
		Message: "Portfolio summary accessed successfully",
	}

	//tell client that response formate is JSON
	w.Header().Set("Content-Type", "application/json")

	//convert response struct into JSON and send back to the client
	json.NewEncoder(w).Encode(response)

}