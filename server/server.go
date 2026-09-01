package server

import (
	"fmt"
	"net/http"
)
type PlayerStore interface {
	GetPlayerScore(name string) int
}

var data = map[string]int{
	"Mahir":200,	
	"Neel":100,	
	"Anita":50,	
	"Dhansukh":0,	

} 
func GetPlayerScore(name string) int  {
	return data[name]
}
func ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	name := r.PathValue("name") 
	score := GetPlayerScore(name)
	fmt.Fprint(w, score)   
}