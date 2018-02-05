package main

import(
    //"database/sql"
    //_"database/sql/driver/mysql"
    "net/http"
    //"golang.org/x/crypto/bcrypt"
     "login"
    "signup"
    "dbconnection"
   "useractivation"
   "profileinfo"
   "profileupdate"
   "logout"
)
//var db *sql.DB 
//var err error

func main(){

  dbconnection.Init()
    //http.HandleFunc("/", login.LoginHandler)
    http.HandleFunc("/login.html", login.LoginHandler)
    http.HandleFunc("/signup.html", signup.SignupHandler)
    http.HandleFunc("/activate-account", useractivation.ActivationHandler)
    http.HandleFunc("/profile.html", profileinfo.ProfileHandler)
    http.HandleFunc("/profile-save", profileupdate.ProfileSaveHandler)
    http.HandleFunc("/logout", logout.LogoutHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080",nil)
}