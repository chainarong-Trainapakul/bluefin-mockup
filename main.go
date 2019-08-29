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
	//fmt.Fprintf("Listprocess")
	//fmt.Fprintf(processListResponseSuccess)
	//fmt.Println(processListResponseSuccess)
	tmp := processListResponseSuccess
	fmt.Println(tmp)
	json.NewEncoder(res).Encode(tmp)
}

func testListProcess(res http.ResponseWriter, req *http.Request){
	resp := Response{
		ID : ID,
		ResultCode: 200,
		Description: "ok test deploy",
		}
	//json.NewEncoder(res).Encode(processListResponseSuccess)
	json.NewEncoder(res).Encode(resp)
}

func handleRequest() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/WriteResponse", WriteResponse).Methods("GET")
	router.HandleFunc("/testDeploy", TestDeploy).Methods("GET")
	router.HandleFunc("/testDeploy2", TestDeploy2).Methods("GET")
	router.HandleFunc("/v1/api/processes", ListProcess).Methods("GET")
	router.HandleFunc("/processes", testListProcess).Methods("GET")
	fmt.Println("start")
	port := os.Getenv("PORT")
	if port == ""{
		port = "8081"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}

type DataResponse struct {
	ID						string 			`json:"id"`
	Url						string 			`json:"url"`
	BusinessKey				string 			`json:"businessKey"`
	Suspended				bool 			`json:"suspended"`
	Ended					bool 			`json:"ended"`
	ProcessDefinitionId		string			`json:"processDefinitionId"`
	ProcessDefinitionUrl	string			`json:"processDefinitionUrl"`
	ProcessDefinitionKey	string			`json:"processDefinitionKey"`
	ActivityId				string			`json:"activityId"`
	//variables must be object
	Variables				string 			`json:"variables"`
	TenantId				string			`json:"tenantId"`
	Name					string			`json:"name"`
	ActiveActivity			[]string		`json:"activeActivity"`
	Completed				string			`json:"completed"`
}

type ProcessResponse struct {
	ResultCode			string			`json:"resultCode"`
	ResultDescription	string			`json:"resultDescription"`
	DevelopMessage		string			`json:"delelopMessage"`
	Data				DataResponse	`json:"data"`
	Start				int				`json:"start"`
	Size				int				`json:"size"`
	Sort				string			`json:"sort"`
	Order				string			`json:"order"`
	Total				int				`json:"total"`
}

func V1ApiProcesses(res http.ResponseWriter, req *http.Request) {
	
}

func main(){
	// handle http request
	handleRequest()
}

var (
	testData = DataResponse{
		ID:						"2501",
		Url:					"null",
		BusinessKey:			"null",
		Suspended:				false,
		Ended:					false,
		ProcessDefinitionId: 	"sampleProcess:1:4",
		ProcessDefinitionUrl: 	"null",
		ProcessDefinitionKey: 	"sampleProcess",
		ActivityId:				"null",
		Variables:				"null object",
		TenantId:				"",
		Name:					"null",
		Completed:				"false",
	}

	processListResponseSuccess = ProcessResponse{
		ResultCode: 			"20000",
		ResultDescription:		"success",
		Data:					testData,
		Start:					0,
		Size:					1,
		Sort:					"id",
		Order:					"asc",
		Total:					1,
	}

	processListResponseError = ProcessResponse{
		ResultCode:				"50000",
		ResultDescription:		"System error",
		Data:					DataResponse{},
	}
)
