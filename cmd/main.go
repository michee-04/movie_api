package main

import (
	"fmt"
	// "log"
	"movie-card/routes"
	// "net/http"
)

const port = ":8000"

// type Movie struct {
// 	Id       string    `json:"id"`
// 	Isbn     string    `json:"isbn"`
// 	Title    string    `json:"title"`
// 	Director *Director `json:"director"`
// }

// type Director struct {
// 	FirstName string `json:"firstname"`
// 	LastName  string `json:"lastname"`
// }


// func Film(){
// 	var movies []Movie

// 	movies = append(movies, Movie{Id: "1", Isbn: "Cool", Title: "Titre", Director: &Director{FirstName: "Mic", LastName: "Michee"}})
// 	movies = append(movies, Movie{Id: "2", Isbn: "fun", Title: "algo", Director: &Director{FirstName: "Mac", LastName: "MacGyver"}})
// }

func main() {

	fmt.Printf("le serveur fonctionne sur http://localhost%s", port)

	routes.Web()


}