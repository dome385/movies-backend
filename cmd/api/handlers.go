package main

import (
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Filme läuft",
		Version: "1.0.0",
	}

	//dieser helfer funktion ersetzt das untere
	_ = app.writeJSON(w, http.StatusOK, payload)

	//davor
	//out, err := json.Marshal(payload)
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	//diese helfer funktion ersetzt das untere
	_ = app.writeJSON(w, http.StatusOK, movies)

	/*davor:
	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)*/
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read json payload

	// validate user against database

	// check password

	// create a jwt user
	u := jwtUser{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
	}

	// generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	log.Println(tokens.Token)

	w.Write([]byte(tokens.Token))

}
