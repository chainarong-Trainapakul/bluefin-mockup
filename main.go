package main
import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)
type Response struct {
	ID uint64 `json:"ID"`
	ResultCode	int `json:"ResultCode"`
	Description string `json:"Description"`
}

var ID uint64

func WriteResponse(res http.ResponseWriter, req *http.Request){
	ID++
	resp := Response{
	ID : ID,
	ResultCode: 200,
	Description: "ok",
	}
	json.NewEncoder(res).Encode(resp)
	fmt.Println("HIT WriteResponse")
}

func TestDeploy(res http.ResponseWriter, req *http.Request){
	ID++
	resp := Response{
		ID : ID,
		ResultCode: 200,
		Description: "ok test deploy",
		}
	json.NewEncoder(res).Encode(resp)
	fmt.Println("TestDeploy")
}

func homePage(res http.ResponseWriter, req *http.Request){
	fmt.Println("-------------------------------------------------------------------")
	fmt.Fprintf(res, "endpint:")
	fmt.Println("method    :", req.Method)
	fmt.Println("URL       :", req.URL)
	fmt.Println("Proto     :", req.Proto)
	fmt.Println("ProtoMajor:", req.ProtoMajor)
	fmt.Println("ProtoMinor:", req.ProtoMinor)
	fmt.Println("Header    :", req.Header)
	fmt.Println("Body      :", req.Body)
	fmt.Println("ContentLen:", req.ContentLength)
	fmt.Println("TransferEncoding:", req.TransferEncoding)
	fmt.Println("Close     :", req.Close)
	fmt.Println("Host      :", req.Host)
	fmt.Println("Form      :", req.Form)
	fmt.Println("PostForm  :", req.PostForm)
	fmt.Println("MultipartForm :", req.MultipartForm)
	fmt.Println("Trailer   :", req.Trailer)
	fmt.Println("RemoteAddr:", req.RemoteAddr)
	fmt.Println("RequestURI:", req.RequestURI)
	fmt.Println("TLS       :", req.TLS)
	fmt.Println("Cancel    :", req.Cancel)
}

func handleRequest() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/WriteResponse", WriteResponse).Methods("GET")
	router.HandleFunc("/testDeploy", TestDeploy).Methods("GET")
	fmt.Println("start")
	port := os.Getenv("PORT")
	if port == ""{
		port = "8081"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main(){
	handleRequest()
}