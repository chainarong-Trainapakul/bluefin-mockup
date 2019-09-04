package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	//"io/ioutil"

	"github.com/gorilla/mux"
)

type Response struct {
	ID          uint64 `json:"ID"`
	ResultCode  int    `json:"ResultCode"`
	Description string `json:"Description"`
}

var ID uint64

func WriteResponse(res http.ResponseWriter, req *http.Request) {
	ID++
	resp := Response{
		ID:          ID,
		ResultCode:  200,
		Description: "ok",
	}
	json.NewEncoder(res).Encode(resp)
	fmt.Println("HIT WriteResponse")
}

func TestDeploy(res http.ResponseWriter, req *http.Request) {
	ID++
	resp := Response{
		ID:          ID,
		ResultCode:  200,
		Description: "ok test deploy",
	}
	json.NewEncoder(res).Encode(resp)
	fmt.Println("TestDeploy")
}

func homePage(res http.ResponseWriter, req *http.Request) {
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

func ListProcess(res http.ResponseWriter, req *http.Request) {
	tmp := processListResponseSuccess
	json.NewEncoder(res).Encode(tmp)
}

func testListProcess(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.FormValue("name"))
	resp := Response{
		ID:          ID,
		ResultCode:  200,
		Description: "ok test deploy",
	}
	json.NewEncoder(res).Encode(resp)
}

type Variables struct {
	Name      string `json:"Name"`
	LeaveType string `json:"LeaveType"`
	From      string `json:"From"`
	To        string `json:"To"`
	LeaveDesc string `json:"LeaveDesc"`
	Field	  []Field	 `json:"Field"`
}

type Process struct {
	ProcessDefinitionKey string    `json:"processDefinitionKey"`
	Initiator            string    `json:"initiator"`
	Variables            Variables `json:"variables"`
}
type Field struct {
	FieldName	string		`json:"fieldName"`
	SelectType	string		`json:"selectType"`
	Value		string		`json:"value"`
}

func Processes(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body)
	switch req.Method {
	case "GET":
		if req.FormValue("initiator") != "" { 
			response := ProcessResponse{
				ResultCode:        "20000",
				ResultDescription: "success",
				Data: []DataResponse{
					DataResponse{
						ID:                   "P001",
						Ended:                false,
						ProcessDefinitionKey: "leave",
						Variables: Variables{
							Name:      req.FormValue("initiator"),
							LeaveType: "sick",
							From:      "2019-06-03T17:26:59.344Z",
							To:        "2019-06-03T17:26:59.344Z",
							LeaveDesc: "Asc",
						},
						Completed: "false",
					},
					DataResponse{
						ID:                   "P002",
						Ended:                false,
						ProcessDefinitionKey: "leave",
						Variables: Variables{
							Name:      req.FormValue("initiator"),
							LeaveType: "vacation",
							From:      "2019-06-03T17:26:59.344Z",
							To:        "2019-06-03T17:26:59.344Z",
							LeaveDesc: "Desc",
						},
						Completed: "false",
					},
					DataResponse{
						ID:                   "P003",
						Ended:                true,
						ProcessDefinitionKey: "leave",
						Variables: Variables{
							Name:      req.FormValue("initiator"),
							LeaveType: "vacation",
							From:      "2019-06-03T17:26:59.346Z",
							To:        "2019-06-03T17:26:59.346Z",
							LeaveDesc: "Desc",
						},
						Completed: "true",
					},
					
				},
			}
			
			json.NewEncoder(res).Encode(response)
			return
		}
		var reqJSON Process
		decoder := json.NewDecoder(req.Body)
		decoder.Decode(&reqJSON)
		response := ProcessResponse{
			ResultCode:        "20000",
			ResultDescription: "success",
			Data: []DataResponse{
				DataResponse{
					ID:                   "5001",
					Ended:                false,
					ProcessDefinitionKey: reqJSON.ProcessDefinitionKey,
					Variables:            reqJSON.Variables,
					Name:                 "",
					Completed:            "false",
					Initiator:            reqJSON.Initiator,
				},
			},
		}
		json.NewEncoder(res).Encode(response)
	case "DELETE":
		response := SimpleResponse{
			ResultCode: "20000",
		}
		//mux.Vars(req)
		if req.FormValue("processInstanceId") != "" {
			response.ResultDescription = "success"
		} else {
			response.ResultDescription = "success"

		}
		json.NewEncoder(res).Encode(response)
		return
	default:
	}

}

type ApproveOrReject struct {
	Action string `json:"action"`
	//Assignee string `json:"assignee"`
}

func Tasks(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if req.FormValue("involvedUser") != "" {
			response := TaskResponse{
				ResultCode:        "20000",
				ResultDescription: "success",
				Data: DataTaskResponse{
					ID:         "T001",
					Assignee:   req.FormValue("involvedUser"),
					CreateTime: "2019-06-03T10:26:05.407Z",
					DueDate:    "",
					Name:       "leaveapproveform",
					Variables:	Variables{
						Field:	[]Field{
							Field{	
								FieldName:	"name",
								SelectType:	"text",
								Value:		"Employee_A",
							},
							Field{
								FieldName:	"lastname",
								SelectType:	"text",
								Value:		"lastnameE",
							},
							Field{
								FieldName:	"email",
								SelectType:	"text",
								Value:		"employee@gmail.com",
							},
							Field{
								FieldName:	"phone",
								SelectType:	"number",
								Value:		"0855555555",
							},
						},
					},
				},
			}
			json.NewEncoder(res).Encode(response)
			return
		}
	case "PUT":
		fmt.Println("active", req.FormValue("active"))
		fmt.Println("includeProcessVariables", req.FormValue("includeProcessVariables"))
		fmt.Println("processInstanceId", req.FormValue("processInstanceId"))
		if req.FormValue("active") != "" && req.FormValue("includeProcessVariables") != "" && req.FormValue("processInstanceId") != "" {
			var reqJSON ApproveOrReject
			decoder := json.NewDecoder(req.Body)
			decoder.Decode(&reqJSON)
			fmt.Println("reqJson:", reqJSON)
			response := SimpleResponse{
				ResultCode: "20000",
			}
			if strings.EqualFold(reqJSON.Action, "Approve") {
				response.ResultDescription = "Approve Success"
			} else if strings.EqualFold(reqJSON.Action, "Reject") {
				response.ResultDescription = "Reject Success"
			}
			json.NewEncoder(res).Encode(response)
			return
		}
	default:
		response := `{"resultCode":"50000","resultDescription":"error"}`
		json.NewEncoder(res).Encode(response)
	}
}

func Test(res http.ResponseWriter, req *http.Request) {
	var a Variables
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&a)
	fmt.Println(a)
}

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/v1/api/processes", Processes).Methods("GET")
	router.HandleFunc("/v1/api/processes/{processInstanceId}", Processes).Methods("DELETE")
	router.HandleFunc("/v1/api/tasks", Tasks).Methods("GET")
	router.HandleFunc("/v1/api/tasks", Tasks).Methods("PUT")
	router.HandleFunc("/test", Test).Methods("GET")

	fmt.Println("start")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}

type DataResponse struct {
	ID                   string    `json:"id"`
	Url                  string    `json:"url"`
	BusinessKey          string    `json:"businessKey"`
	Suspended            bool      `json:"suspended"`
	Ended                bool      `json:"ended"`
	ProcessDefinitionId  string    `json:"processDefinitionId"`
	ProcessDefinitionUrl string    `json:"processDefinitionUrl"`
	ProcessDefinitionKey string    `json:"processDefinitionKey"`
	ActivityId           string    `json:"activityId"`
	Variables            Variables `json:"variables"`
	TenantId             string    `json:"tenantId"`
	Name                 string    `json:"name"`
	ActiveActivity       []string  `json:"activeActivity"`
	Completed            string    `json:"completed"`
	Initiator            string    `json:"initiator"`
}

type ProcessResponse struct {
	ResultCode        string       `json:"resultCode"`
	ResultDescription string       `json:"resultDescription"`
	DevelopMessage    string       `json:"delelopMessage"`
	Data              []DataResponse `json:"data"`
	Start             int          `json:"start"`
	Size              int          `json:"size"`
	Sort              string       `json:"sort"`
	Order             string       `json:"order"`
	Total             int          `json:"total"`
}

type DataTaskResponse struct {
	ID         string    `json:"id"`
	Assignee   string    `json:"assignee"`
	Name       string    `json:"name"`
	CreateTime string    `json:"createTime"`
	DueDate    string    `json:"dueDate"`
	Variables  Variables `json:"variables"`
}

type TaskResponse struct {
	ResultCode        string           `json:"resultCode"`
	ResultDescription string           `json:"resultDescription"`
	Data              DataTaskResponse `json:"data"`
}

type SimpleResponse struct {
	ResultCode        string `json:"resultCode"`
	ResultDescription string `json:"resultDescription"`
}

func main() {
	handleRequest()
}

var (
	testData = []DataResponse{
		DataResponse{
			ID:                   "2501",
			Url:                  "null",
			BusinessKey:          "null",
			Suspended:            false,
			Ended:                false,
			ProcessDefinitionId:  "sampleProcess:1:4",
			ProcessDefinitionUrl: "null",
			ProcessDefinitionKey: "sampleProcess",
			ActivityId:           "null",
			Variables:            Variables{},
			TenantId:             "",
			Name:                 "null",
			Completed:            "false",
		},
	}

	processListResponseSuccess = ProcessResponse{
		ResultCode:        "20000",
		ResultDescription: "success",
		Data:              testData,
		Start:             0,
		Size:              1,
		Sort:              "id",
		Order:             "asc",
		Total:             1,
	}

	processListResponseError = ProcessResponse{
		ResultCode:        "50000",
		ResultDescription: "System error",
		Data:              []DataResponse{},
	}
)
