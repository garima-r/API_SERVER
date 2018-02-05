package useractivation
import(
	_"database/sql/driver/mysql"
	"net/http"
	"log"
	"fmt"
  //  "errors"
	//"crypto/aes"
    //"crypto/cipher"
    // "crypto/rand"
    b64"encoding/base64"
    "net/url"
    "strconv"
    "dbconnection"
)

func ActivationHandler(w http.ResponseWriter, r *http.Request){
   
    fmt.Println("entered inside")
	
    activationurl := r.URL.RequestURI()
	u, err:= url.Parse(activationurl)
	if err != nil {
		log.Fatal(err)
	}
	mail,_ :=url.ParseQuery(u.RawQuery)
	uid:= mail["registration-id"][0]

    id, _ := b64.URLEncoding.DecodeString(uid)

	fmt.Println("decoded id is:",int(id[0]))
    a:=strconv.Itoa(int(id[0]))
    fmt.Println("Value of a: ",a)

    //regid, _ := strconv.Atoi(string(id))
    //fmt.Println("check:",regid)


	_,err=dbconnection.Db.Exec("update user set status='active' where reg_id="+a)
    if err != nil {
        log.Fatal(err)
    }else{
        fmt.Fprintf(w,"<script>alert('Account is ready to use!!');</script>")
        http.Redirect(w,r,"/login.html",303)
    }


}




    /*key := []byte("the-key-has-to-be-32-bytes-long!")

	id, err := decrypt([]byte(uid), key)
    if err != nil {
        // TODO: Properly handle error
        log.Fatal(err)
    }
    fmt.Printf("%x => %s\n", , plaintext)




    //change this query!!!!


	stmt, err := dbconnection.Db.Prepare("update user set status=? where reg_id=?")
        if err != nil{
				log.Fatal(err)
				fmt.Println(err)
		}else{
            fmt.Println("send")
        }

  	//Estring := strings.Join(to, "")
  	//fmt.Println("js" , Estring)
	_, err = stmt.Exec("active", string(id))
	if err != nil{
			log.Fatal(err)
			fmt.Println(err)
	}
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
    c, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        return nil, err
    }

    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, errors.New("ciphertext too short")
    }

    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    return gcm.Open(nil, nonce, ciphertext, nil)
}*/