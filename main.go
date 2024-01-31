package main

import (
	"fmt"
	"html/template"

	"net/http"

	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/sessions"

	_ "github.com/go-sql-driver/mysql"
)

// Add .env values
var store = sessions.NewCookieStore([]byte("temporary_secret"))

var db *sql.DB

//var tmpl *template.Template

func VideoServe(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]

	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	hlsRep()
	tmpl, err := template.ParseFiles("views/Player.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)

}

func connectDb() {
	var err error
	db, err = sql.Open("mysql", "root:pato@tcp(db:3306)/sapoing")
	for err != nil {
		time.Sleep(5 * time.Second)
		db, err = sql.Open("mysql", "root:pato@tcp(db:3306)/sapoing")
	}

}

func main() {
	go connectDb()
	//http.Handle("/vid_src/{$}", http.FileServer(http.Dir("./src")))
	http.Handle("/vid_src/", http.StripPrefix("/vid_src/", http.FileServer(http.Dir("./src"))))
	http.HandleFunc("/login", LoginPage)
	http.HandleFunc("/loginauth", LoginPageHandler)
	http.HandleFunc("/register", SignupPage)
	http.HandleFunc("/registerauth", SignupPageAuth)
	http.HandleFunc("/video", VideoServe)
	http.HandleFunc("/", LoginPage)

	http.HandleFunc("/static/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/css")
		http.ServeFile(w, r, "./static/style.css")
	})

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	defer db.Close()
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)

}

func SignupPageAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userDb string
		var hash []byte
		username := r.FormValue("username")
		password := r.FormValue("password")

		statement := "SELECT username FROM users WHERE username = ?"
		row := db.QueryRow(statement, username)
		err := row.Scan(&userDb)

		if err != sql.ErrNoRows {

			if err == nil {

				tmpl, err := template.ParseFiles("views/register.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				tmpl.Execute(w, nil)
				return

			} else {

				tmpl, err := template.ParseFiles("views/register.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				tmpl.Execute(w, nil)
				return
			}
		}

		hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
			tmpl, err := template.ParseFiles("views/register.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, nil)
			return
		}

		var stmt *sql.Stmt
		stmt, err = db.Prepare("INSERT INTO users (username, pwrd) VALUES (?,?);")
		if err != nil {
			tmpl, err := template.ParseFiles("views/register.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, nil)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(username, string(hash))

		if err != nil {
			tmpl, err := template.ParseFiles("views/register.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, nil)

			return
		}

		tmpl, err := template.ParseFiles("views/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)

		return
	}

}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userDb, hash string

		username := r.FormValue("username")
		password := r.FormValue("password")

		statement := "SELECT username, pwrd FROM users WHERE username = ?"
		row := db.QueryRow(statement, username)
		err := row.Scan(&userDb, &hash)

		if err != nil {
			tmpl, err := template.ParseFiles("views/login.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, "Revise el usuario y la clave")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err == nil {
			session, _ := store.Get(r, "session")
			session.Values["username"] = username
			session.Save(r, w)
			http.Redirect(w, r, "/video", http.StatusFound)
			return
		}

		tmpl, err := template.ParseFiles("views/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, "Revise el usuario y la clave")

		return

	}

	tmpl, err := template.ParseFiles("views/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("views/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
