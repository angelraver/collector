package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func UserForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// fmt.Fprint(w, r.URL.Query()["gg"])
	// fmt.Fprint(w, r.URL.Query().Get("gg"))
	// q := r.URL.Query()
	// name := q.Get("name")
	// fmt.Fprint(w, name)
	fmt.Fprint(w, r.FormValue("nombre"))
}

func UserJson(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Name string
	}
	var user User
	// body, err := ioutil.ReadAll(r.Body)
	// err = json.Unmarshal(body, &user)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}
	fmt.Println(user.Name)
}

type User struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

func GetUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	results, err := db.Query("SELECT id, name FROM user")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var usuarios []User

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.id, &user.name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		usuarios = append(usuarios, user)
	}
	// data := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, usuarios)
	// w.Write(data)
}

// perform a db.Query insert
// insert, err := db.Query("CALL userCreate(?,?)", "angel", "123456")
// if err != nil {
//   panic(err.Error())
// }
// be careful deferring Queries if you are using transactions
// defer insert.Close()
