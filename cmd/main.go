package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"
	"www/iternal/session"

	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id                   int
	Name, Password, HASH string
}

//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

//type handler func(w http.ResponseWriter, r *http.Request, s *sessions.Session)
const (
	COOKIE_NAME = "sessionId"
)

var inMemorySession *session.Session

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/templates/index.html", "assets/templates/header.html", "assets/templates/footer.html")

	db, err := sql.Open("sqlite3", "init/bd.sq3")
	checkError(err)
	defer db.Close()

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/templates/create.html", "assets/templates/header.html", "assets/templates/footer.html")

	t.ExecuteTemplate(w, "create", nil)
}
func save_user(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {

		db, err := sql.Open("sqlite3", "init/bd.sq3")
		checkError(err)
		defer db.Close()

		sql2 := "INSERT INTO `user` (`username`, `password`) VALUES('" + username + "', '" + password + "')"
		sql1, _ := db.Prepare(sql2)
		tx, _ := db.Begin()
		_, _ = tx.Stmt(sql1).Exec()
		tx.Commit()
		defer sql1.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
func auth(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/templates/auth.html", "assets/templates/header.html", "assets/templates/footer.html")

	t.ExecuteTemplate(w, "auth", nil)
}
func valid(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {
		db, _ := sql.Open("sqlite3", "init/bd.sq3")
		defer db.Close()

		sql3 := "SELECT * FROM `user` WHERE  username = " + username + " and password = " + password + ""
		res, _ := db.Query(sql3)
		if res == nil {
			fmt.Fprintf(w, "Неверные данные")
		} else {

			sessionId := inMemorySession.Init(username)
			cookie := &http.Cookie{
				Name:    COOKIE_NAME,
				Value:   sessionId,
				Expires: time.Now().Add(5 * time.Minute),
			}
			http.SetCookie(w, cookie)

			//		http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		defer res.Close()
	}
}
func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/auth", auth).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_user", save_user).Methods("POST")
	rtr.HandleFunc("/valid", valid).Methods("POST")
	http.Handle("/", rtr)

	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./style/"))))
	http.ListenAndServe(":8080", nil)
}
func main() {
	handleFunc()
	inMemorySession = session.NewSession()
}
