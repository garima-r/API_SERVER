package logout
import(
	"net/http"
	"fmt"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		cookieexpire := http.Cookie{Name: "loggedin", Value:" ", Expires: time.Now()}
		http.SetCookie(w, &cookieexpire)
		fmt.Println("cookievalue", cookieexpire.Value)

		http.Redirect(w,r,"/login.html",301)
	}
}
