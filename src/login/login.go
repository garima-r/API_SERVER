
package login
import(
	//"database/sql"
	_"database/sql/driver/mysql"
	"net/http"
	"log"
	"time"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"dbconnection"
)

//var db *sql.DB 
//var err error

func LoginHandler(w http.ResponseWriter, r *http.Request){

if r.Method !="POST"{
		http.ServeFile(w,r,"login.html")
		return
}

Email:= r.FormValue("email_id") 
Password:= r.FormValue("password")

//Validating the credentials of the user
var pass string
var user_status string
row, err := dbconnection.Db.Query("select password, status from user where email_id ='"+Email+"'")
fmt.Println("afetr query")
if err != nil {
	log.Fatal(err)
}
defer row.Close()
for row.Next() {
	 
    err := row.Scan(&pass, &user_status);
	if err != nil {
			log.Fatal(err)
		}
	log.Println(pass)
	log.Println(user_status)

}

err1 := bcrypt.CompareHashAndPassword([]byte(pass), []byte(Password))
		
if (user_status == "active"){

		fmt.Println("User_status: active and password matched")


		cookie_exp :=time.Now().Add(365*24*time.Hour)
		user_cookie := http.Cookie{ Name: "loggedin", Value: Email, Expires: cookie_exp }
		http.SetCookie(w, &user_cookie)
		fmt.Println("cookievalue", user_cookie.Value)

		
		if err1 != nil {
			fmt.Println("Password not matched")
		//log.Fatal(err1)
			fmt.Fprintf(w,"<script> alert('Incorrect Email-Id or Password');</script>")//;window.location.href='/signup'
		//http.ServeFile(w,r,"/login.html")   not wotking
			http.Redirect(w,r,"/login.html",301)  //not wotking
		}else{
			http.Redirect(w,r,"/profile.html",301)
		}

} else{

			fmt.Fprintf(w,"<script> alert('Please activate your account');</script>")
			http.Redirect(w,r,"/login",301) //not working
}


}
/*if err != nil {
	fmt.Println("Password not matched")
	log.Fatal(err1)
	fmt.Fprintf(w,"<script> alert('Incorrect Email-Id or Password');</script>")//;window.location.href='/signup'
	http.ServeFile(w,r,"/login.html")
	//http.Redirect(w,r,"/login.html",301)

else if  (user_status == "active"){

		fmt.Println("User_status: active and password matched")


		cookie_exp :=time.Now().Add(365*24*time.Hour)
		user_cookie := http.Cookie{ Name: "loggedin", Value: Email, Expires: cookie_exp }
		http.SetCookie(w, &user_cookie)
		fmt.Println("cookievalue", user_cookie.Value)

		http.Redirect(w,r,"/profile.html",301)

	}  else{

		fmt.Fprintf(w,"<script> alert('Please activate your account');</script>")
		http.Redirect(w,r,"/login",301)
	}*/
