package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var dbHome *sql.DB
var err error

type User struct {
	ID          int
	Name        string
	Msisdn      string
	Email       string
	birthDate   time.Time
	createdTime time.Time
	updateTime  time.Time
	userAge     int
}

func main() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/json", handleJSON)
	http.HandleFunc("/big-project", handleBigProject)
	http.HandleFunc("/search", handleSearch)
	http.ListenAndServe(":8080", nil)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

type People struct {
	private string
	ID      int
	Name    string
	Address string
	Float   float64
	Bool    bool
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	people := People{
		"private",
		123,
		"Jeki",
		"earth",
		0.002,
		true,
	}

	byt, err := json.Marshal(people)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(byt)
}

func handleBigProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	byt, err := json.Marshal(getUser())
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(byt)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	name, _ := r.URL.Query()["name"]
	byt, err := json.Marshal(getUserByTerm(name[0]))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(byt)
}

func getUser() []User {

	dbHome, err = sql.Open("postgres", "postgres://yx180102:6R5X53xDX0xrx1@devel-postgre.tkpd/tokopedia-user?sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer dbHome.Close()
	stmt := "select user_id, full_name, msisdn, user_email, birth_date, create_time, update_time from ws_user ORDER BY user_id DESC limit 10"
	rows, err := dbHome.Query(stmt)
	userList := []User{}
	for rows.Next() {
		c := User{}
		_ = rows.Scan(&c.ID, &c.Name, &c.Msisdn, &c.Email, &c.birthDate, &c.createdTime, &c.updateTime)
		userList = append(userList, c)
	}

<<<<<<< 79d6318c389277687d4e7b40a3fb3532534f7aba
	http.HandleFunc("/hello", hwm.SayHelloWorld)
<<<<<<< HEAD
	http.HandleFunc("/go", hwm.SayIntroToGo)
=======
>>>>>>> 00f41174e6e7cd1c36a3d3ac10eea8aaea40a57c
	go logging.StatsLog()
=======
	if err != nil {
		log.Println(err)
	}
	return userList
}
>>>>>>> yulianto

func getUserByTerm(s string) []User {
	dbHome, err = sql.Open("postgres", "postgres://yx180102:6R5X53xDX0xrx1@devel-postgre.tkpd/tokopedia-user?sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	defer dbHome.Close()
	stmt := "select user_id, full_name, msisdn, user_email, birth_date, create_time, update_time from ws_user where lower(full_name) like lower('%'||'" + strings.ToLower(s) + "'||'%') ORDER BY user_id DESC limit 10"
	rows, err := dbHome.Query(stmt)
	userList := []User{}
	if rows == nil {
		return userList
	}
	for rows.Next() {
		c := User{}
		_ = rows.Scan(&c.ID, &c.Name, &c.Msisdn, &c.Email, &c.birthDate, &c.createdTime, &c.updateTime)
		userList = append(userList, c)
	}

	if err != nil {
		log.Println(err)
	}
	// fmt.Println(s)
	return userList
}
