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

func TestDeploy2(res http.ResponseWriter, req *http.Request){
	ID++
	resp := Response{
		ID : ID,
		ResultCode: 200,
		Description: "ok test deploy2",
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

func ListProcess(res http.ResponseWriter, req *http.Request){
	json.NewEncoder(res).Encode(processListResponseSuccess)
}

func handleRequest() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/WriteResponse", WriteResponse).Methods("GET")
	router.HandleFunc("/testDeploy", TestDeploy).Methods("GET")
	router.HandleFunc("/testDeploy2", TestDeploy2).Methods("GET")
	router.HandleFunc("/v1/api/processes", ListProcess).Methods("GET")
	fmt.Println("start")
	port := os.Getenv("PORT")
	if port == ""{
		port = "8081"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}

type DataResponse struct {
	id						uint64 			`json:"id"`
	url						string 			`json:"url"`
	businessKey				string 			`json:"businessKey"`
	suspended				bool 			`json:"suspended"`
	ended					bool 			`json:"ended"`
	processDefinitionId		string			`json:"processDefinitionId"`
	processDefinitionUrl	string			`json:"processDefinitionUrl"`
	processDefinitionKey	string			`json:"processDefinitionKey"`
	activityId				string			`json:"activityId"`
	//variables must be object
	variables				string 			`json:"variables"`
	tenantId				string			`json:"tenantId"`
	name					string			`json:"name"`
	activeActivity			[]string		`json:"activeActivity"`
	completed				string			`json:"completed"`
}

type ProcessResponse struct {
	resultCode			string			`json:"resultCode"`
	resultDescription	string			`json:"resultDescription"`
	developMessage		string			`json:"delelopMessage"`
	data				DataResponse	`json:"data"`
	start				int				`json:"start"`
	size				int				`json:"size"`
	sort				string			`json:"sort"`
	order				string			`json:"order"`
	total				int				`json:"total"`
}

func V1ApiProcesses(res http.ResponseWriter, req *http.Request) {
	data := DataResponse {
		id : id,
		
	}
}

func main(){
	// handle http request
	handleRequest()
}

testData := DataResponse{
	id:						"2501",
	url:					null,
	businessKey:			null,
	suspended:				false,
	ended:					false,
	processDefinitionId: 	"sampleProcess:1:4",
	processDefinitionUrl: 	null,
	processDefinitionKey: 	"sampleProcess",
	activityId:				null,
	variables:				"null object"
	tenantId:				"",
	name:					null,
	completed:				false,
}

processListResponseSuccess := ProcessResponse{
	resultCode: 			"20000",
	resultDescription:		"success",
	data:					testData,
	start:					0,
	size:					1,
	sort:					"id",
	order:					"asc",
	total:					"1",
}

processListResponseError := ProcessResponse{
	resultCode:				"50000",
	resultDescription:		"System error",
	data:					DataResponse{},
}

