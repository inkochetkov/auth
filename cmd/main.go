package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"www/iternal/rands"

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
	Id       string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	HASH     string `db:"hash"`
}

type Auth struct {
	Id int
}

//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

//type handler func(w http.ResponseWriter, r *http.Request, s *sessions.Session)
const (
	COOKIE_NAME = "sessionId"
)

func index(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(COOKIE_NAME)
	var auth Auth
	if COOKIE_NAME == "" {
		auth.Id = 0
	} else {
		auth.Id = 1
	}
	log.Println(cookie)

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

		HASH := rands.GenerateId()

		sql2 := "INSERT INTO `user` (`username`, `password`, `HASH`) VALUES('" + username + "', '" + password + "', '" + HASH + "')"
		sql1, _ := db.Prepare(sql2)
		tx, _ := db.Begin()
		_, _ = tx.Stmt(sql1).Exec()
		tx.Commit()
		defer sql1.Close()

		cookie := &http.Cookie{
			Name: COOKIE_NAME,
			//	Domain: "localhost",
			//  Path: "/",
			//  HttpOnly: false,
			//  Secure: true,
			//  MaxAge: 86400,
			Expires: time.Now().AddDate(1, 0, 0),
		}
		cookie.Value = HASH
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
func auth(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/templates/auth.html", "assets/templates/header.html", "assets/templates/footer.html")

	t.ExecuteTemplate(w, "auth", nil)
}
func valid(w http.ResponseWriter, r *http.Request) {
	U := r.FormValue("username")
	P := r.FormValue("password")
	if U == "" || P == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {
		db, _ := sql.Open("sqlite3", "init/bd.sq3")
		defer db.Close()

		sql3 := fmt.Sprintf("SELECT * FROM `user` WHERE  `username` = %s and `password`=%s", U, P)
		res, err := db.Query(sql3)
		checkError(err)
		for res.Next() {
			var user User
			err = res.Scan(&user.Id, &user.Username, &user.Password, &user.HASH)
			checkError(err)
			//	fmt.Println(user.Id, user.Username, user.Password, user.HASH)

			cookie := &http.Cookie{
				Name: COOKIE_NAME,
				//	Domain: "localhost",
				//  Path: "/",
				//  HttpOnly: false,
				//  Secure: true,
				//  MaxAge: 86400,
				Expires: time.Now().AddDate(1, 0, 0),
			}
			cookie.Value = user.HASH
			http.SetCookie(w, cookie)

		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name: COOKIE_NAME,
		//	Domain: "localhost",
		//  Path: "/",
		//  HttpOnly: false,
		//  Secure: true,
		MaxAge:  0,
		Expires: time.Now().AddDate(0, 0, 0),
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/auth", auth).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_user", save_user).Methods("POST")
	rtr.HandleFunc("/valid", valid).Methods("POST")
	rtr.HandleFunc("/logout", logout).Methods("GET")
	http.Handle("/", rtr)

	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./style/"))))
	http.ListenAndServe(":8080", nil)
}
func main() {
	handleFunc()
	//	inMemorySession = session.NewSession()
}
