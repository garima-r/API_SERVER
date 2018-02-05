package profileupdate
import(
	//"database/sql"
	_"database/sql/driver/mysql"
	"net/http"
	"log"
	"fmt"
	"dbconnection"
)

func ProfileSaveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		cookie, err := r.Cookie("loggedin")
		fmt.Println("cokiiiiii",cookie.Value)
		if err != nil {
				// User is not logged in and don't have any cookie
				fmt.Println("err", err)
				http.Redirect(w,r,"/login.html",301)
				return
		}

		fmt.Println("yes working!!")

		Pfname:= r.FormValue("fn")
		Plname:=  r.FormValue("ln")
		Pdob:= r.FormValue("dob")

		fmt.Println("FNAME:",Pfname)
		fmt.Println("LNAME:",Plname)
		fmt.Println("DATE:",Pdob)

		_, err= dbconnection.Db.Exec("update user set fname='"+Pfname+"', lname='"+Plname+"', dob='"+Pdob+"' where email_id='"+string(cookie.Value)+"'")
	     
       if err != nil{
		log.Fatal("update:error",err)
		fmt.Println(err)
		return
		}
		http.Redirect(w,r,"/profile.html", 301)
	}
}
