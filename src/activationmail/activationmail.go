package activationmail
import(
	//"net/http"
	b64"encoding/base64"
    "net/smtp"
	"log"
	"fmt"
)
/*var err error
func Send_mail(to [] string, uid string){
	auth:=smtp.PlainAuth("","godummytest@gmail.com","gomailtest","smtp.gmail.com")
	if err != nil{
			log.Fatal(err)
			fmt.Println(err)
	}
	
	//Code to send the activation link to the user via email
	activation_link := "http://localhost:8080/activate-account?registration-id="+uid;
	msg:=[]byte("Registration sucessfull !!\r\n"+
				"Please click on the link to activate your account\r\n"+activation_link)
	err3:=smtp.SendMail("smtp.gmail.com:587",auth,"godummytest@gmail.com",to,msg)
	if err3!=nil{
			log.Fatal(err3)
			fmt.Println("smtp error: %s", err3)
	return
	}
}*/

func Send_mail(uid [] string, to [] string){
   
    codeid := b64.StdEncoding.EncodeToString([]byte(uid[0]))
    fmt.Println("encoded id:", codeid)

/*text := []byte(uid)
    key := []byte("the-key-has-to-be-32-bytes-long!")
    ciphertext, err := encrypt(text, key)

    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s => %x\n", text, ciphertext)
    //fmt.Println(ciphertext)
   // fmt.Println("inside send mail, emailid is",to)

   plaintext, err := decrypt(ciphertext, key)
    if err != nil {
        // TODO: Properly handle error
        log.Fatal(err)
    }
    fmt.Printf("%x => %s\n", ciphertext, plaintext)*/

    auth:=smtp.PlainAuth("","godummytest@gmail.com","gomailtest","smtp.gmail.com")
  
    //Code to send the activation link to the user via email
    activation_link := "http://localhost:8080/activate-account?registration-id="+codeid;
    msg:=[]byte("Registration sucessfull !!\r\n"+
                "Please click on the link to activate your account\r\n"+ activation_link)
    err3:=smtp.SendMail("smtp.gmail.com:587",auth,"godummytest@gmail.com",to,msg)
    if err3!=nil{
            log.Fatal(err3)
            fmt.Println("smtp error: %s", err3)
    }else{
        fmt.Println("mail sent sucessfully")
    }
    
}

/*func encrypt(plaintext []byte, key []byte) ([]byte, error) {
    c, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}
*/