package signup

import(
	//"database/sql"
	_"database/sql/driver/mysql"
	"net/http"
	"log"
	"fmt"
	"time"
	"golang.org/x/crypto/bcrypt"
	//"net/smtp"
	"activationmail"
	"dbconnection"

)

//var db *sql.DB 
var err error

func SignupHandler(w http.ResponseWriter, r *http.Request){
	if r.Method !="POST"{
		http.ServeFile(w,r,"signup.html")
		return
	}
//insertion of signup form values to the database table
	Fname:= r.FormValue("fname")
	Lname:=  r.FormValue("lname")
	Gender:=   r.FormValue("gender")
	Dob:= r.FormValue("birthdate")
	Email:=    r.FormValue("email_id")
	Password:= r.FormValue("password")

	//password encryption
	hash,_:= bcrypt.GenerateFromPassword([]byte(Password), 13)
	fmt.Println("asd:" +  string(hash) )
	
	/*_, err := dbconnection.Db.Exec("insert into inactive_user(fname,email_id)VALUES(?,?)",Fname,Email)
	if err != nil{
		log.Fatal(err)
		fmt.Println("error in insert for inactiveuser is:",err)
	}*/
	
	_, err1 := dbconnection.Db.Exec("insert into user(fname,lname,gender,dob,email_id,password,reg_date,status)VALUES(?,?,?,?,?,?,?,?)",Fname,Lname,Gender,Dob,Email,string(hash),time.Now(),"inactive")
	if err1 != nil{
		log.Fatal(err1)
		fmt.Println("error in insert for user is:",err1)
		fmt.Println("<script> alert('User already exist !'); window.location.href='/signup.html';</script>")
		fmt.Fprintf(w,"<b>User Already Exist !! </b>")
	}else{		
	
	/*if dbconnection.err 
		1fmt.Println("<script> alert('User already exist !'); window.location.href='/signup.html';</script>")
		fmt.Fprintf(w,"<b>User Already Exist !! </b>")
		//http.Redirect(w,r,"/signup.html",301)
	}*/
	
	w.Write([]byte("<script> alert('Signup successfull ! Please Activate Your Account via the link send to your registered Email-Id'); window.location.href='/login';</script>"))
	}//Encryption of Registration_id of the user to send as a token in activation link
	var  id int
	err = dbconnection.Db.QueryRow("select reg_id from user where email_id = ?", Email).Scan(&id)
	if err != nil {
			log.Fatal(err)
			return
	}

	fmt.Println("id is:",id)
	a:=string(id)
	uid := [] string{a}
	fmt.Println(Email)
	eid:=[] string{Email}
	//fmt.Println("email id is:",eid)
	activationmail.Send_mail(uid, eid)
	fmt.Println("sdf")
	//http.Redirect(w,r,"/login.html",303)
	http.ServeFile(w,r,"login.html")
	
}