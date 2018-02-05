package profileinfo
import(
	_"database/sql/driver/mysql"
	"net/http"
	"log"
	"fmt"
	"html/template"
	"path"
	"os"
	"dbconnection"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
    
	if r.Method!= "POST"{
			cookie, err := r.Cookie("loggedin")
			fmt.Println("cokiiiiii",cookie.Value)
			if err != nil {
				// User is not logged in and don't have any cookie
				fmt.Println("err", err)
				http.Redirect(w,r,"/login.html",301)
				return
			}

			fmt.Println("cookievalue", cookie)
			lp := path.Join("", "profile.html")


    // Return a 404 if the template doesn't exist
    	info, err := os.Stat(lp)
    	if err != nil {
        	if os.IsNotExist(err) {
            		http.NotFound(w, r)
            	return
        	}
   		  }

   
    	if info.IsDir() {
        		http.NotFound(w, r)
        		return
   		 }

   		templates, err := template.ParseFiles(lp)
    	if err != nil {
       			fmt.Println(err)
        		http.Error(w, "500 Internal Server Error", 500)
        return
    	}

	//Code to display the profile information of the user	
		type UserProfile struct{
	 			FirstName string
	 			LastName  string
	 			Dob string
	 			Gender string
	 			Email_Id string
		}
		data := UserProfile{} 
		
		rows, err := dbconnection.Db.Query("select fname,lname,dob,gender,email_id from user where email_id = ?",string(cookie.Value))//userck.Value)//user_ck.Value)
		if err != nil {
				log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
				err := rows.Scan(&data.FirstName, &data.LastName, &data.Dob, &data.Gender, &data.Email_Id);
				if err != nil {
						log.Fatal(err)
				}
				log.Println(data.FirstName)
				log.Println(data.Gender)
		}
		w.Header().Set("Content-Type", "text/html")
    	templates.ExecuteTemplate(w, "profile.html",data)
    }
	
}
