package main

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
)

const
(
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)
var cookieHandler *securecookie.SecureCookie

func init() {
	// we need hashKey [] byte,blockKey [] byte
	cookieHandler = securecookie.New(securecookie.
	GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))
}

func createCookie(w http.ResponseWriter, r *http.Request) {
	//create map as key and value
	// key is username and value is foo
	value := map[string]string{
		"username": "Foo",
	}
	//encode by sring and value(key and value)
	base64Encoded, err := cookieHandler.Encode("key", value)
	if err == nil  {
		cookie := &http.Cookie {
		Name: "first-cookie",
		Value: base64Encoded,
		Path: "/",
	}
		//Set cookie to writer
	http.SetCookie(w, cookie)
	}
w.Write([]byte(fmt.Sprintf("Cookie created.")))
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	//write a message
	log.Printf("Reading Cookie..")
	//read cookie from Request
	cookie, err := r.Cookie("first-cookie")
	if cookie != nil && err == nil {
		//create a map with key and value
		value := make(map[string]string)
		//check if error hanppend
		if err = cookieHandler.Decode("key", cookie.Value, &value);
			err == nil{
			w.Write([]byte(fmt.Sprintf("Hello %v \n",
				value["username"])))
		}
	} else{
		log.Printf("Cookie not found..")
		w.Write([]byte(fmt.Sprint("Hello")))
	}
}

func main() {
	http.HandleFunc("/create", createCookie)
	http.HandleFunc("/read", readCookie)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}