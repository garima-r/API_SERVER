package main
import(
	"fmt"
	_"database/sql/driver/mysql"
	"activationmail"
	"log"
	"cron"
   "sync"
	"dbconnection"
)

var id int
var email string

func main(){ 

 	dbconnection.Init()
	
 	wg := &sync.WaitGroup{}
    c := cron.New()
    fmt.Println("cron job inialize")
    wg.Add(1)
    c.AddFunc(" * * * * * *",Scheduler)
    fmt.Println("cron 1")
    c.Start()
    fmt.Println("cron job ")
    defer c.Stop() 
     fmt.Println("cron job 121")
    wg.Wait()
}

func Scheduler(){
fmt.Println("4")
rows, err := dbconnection.Db.Query("select reg_id, email_id from user where datediff(curdate(), reg_date)="+"7")
fmt.Println("5")
//fmt.Println("ans is",rows)
if err != nil {
	log.Fatal(err)
	fmt.Println(err)
}
//defer rows.Close()
for rows.Next() {
	//fmt.Println("enter for")
	rows.Scan(&id,&email)
	inac_email :=[]string{email}
	reg_id:=string(id)
	inac_id :=[]string{reg_id}


	activationmail.Send_mail(inac_id, inac_email )



	fmt.Println(inac_email)
	fmt.Println(inac_id)
	fmt.Println("6")
}
	fmt.Println("7")
}