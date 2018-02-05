package main
import(

	"net/http"
	"encoding/json"
	"fmt"
)

type userinfo struct{

	Email string `json: "email_id"`
	Password string `json: "password"`
}
func main(){

	http.HandleFunc("/",encode)
	http.HandleFunc("/new.html",encode)
	//http.Handle("")
	http.ListenAndServe(":8083",nil)
}
func encode(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET"{
  http.ServeFile(w,r,"new.html")
	//}else{	
		var user userinfo
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("user name is:", user.Email)
		fmt.Println("and password is:", user.Password)
		http.Redirect(w,r,"log.html",303)
	}

}