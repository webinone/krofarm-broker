package main

import (
	"net/http"
        "github.com/gorilla/mux"
	mqttlib "krofarm/base/mqttlib"
	//MQTT "github.com/eclipse/paho.mqtt.golang"
	"krofarm/base/protobuf"
	"github.com/golang/protobuf/proto"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
	"time"
	"unsafe"
	"reflect"
	"gopkg.in/natefinch/lumberjack.v2"
	"strconv"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("Hello, Go Lang Krofarm MQTT Message broker")) // 웹 브라우저에 응답

}

// UdfHandler
func UdfHandler (res http.ResponseWriter, req *http.Request) {


	fmt.Println("#################### UdfHandler !!!!!!!")
	log.Println("#################### UdfHandler !!!!!!!")

	vars := mux.Vars(req)
	endpntId := vars["endpntId"]
	accessToken := req.FormValue("accessToken")
	apiKey := req.FormValue("apiKey")
	timestamp := req.FormValue("timestamp")

	fmt.Println("endpntId", endpntId)
	fmt.Println("accessToken", accessToken)
	fmt.Println("apiKey", apiKey)
	fmt.Println("timestamp", timestamp)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	var jsonData map[string]interface{}

	json.Unmarshal([]byte(body), &jsonData)

	fmt.Println("udfNm", jsonData["udfNm"])
	fmt.Println("udfTy", jsonData["udfTy"])
	fmt.Println("udfBody", jsonData["udfBody"])

	// MQTT 데이터 전송
	//--------------------------------------------------

	udfPayload := &protobuf.UDFPayload{
		FnctTyCd: proto.Int32(10160001),
		UdfTy: proto.String(jsonData["udfTy"].(string)),
		UdfNm: proto.String(jsonData["udfNm"].(string)),
		UdfBody:proto.String(jsonData["udfBody"].(string)),
		CreatedAt:proto.Int64(time.Now().UnixNano()),
	}

	fmt.Println("############# UDF BODY START ")

	fmt.Println(udfPayload.GetUdfBody())

	fmt.Println("############# UDF BODY END ")


	client := mqttlib.NewLocalClient()

	pubTopic := "/ksn/v2/L/CMD/POST/UDF"

	message, err := proto.Marshal(udfPayload)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	token := client.Publish(pubTopic, 0, false, message)
	token.Wait()

	fmt.Println("############# UDF PUBLISH !!! ")

	client.Disconnect(0)
	//--------------------------------------------------
	resultJson := `
	{
	    "header": {
		"resultCode": 201,
		"resultMessage": "Created"
	    },
	    "body": {
		"item": {
		    "reqId": 3661303731663466,
		    "sndngCd": 10150002
		}
	    }
	}
	`
	res.Header().Add("Content-Type", "Application/json")
	res.Write([]byte(resultJson)) // 웹 브라우저에 응답

}

type ConfDevice struct {
	DvcId 			int64 `json:"dvcId"`
	FnctngModeCd 		int32 `json:"fnctngModeCd"`
	TotalOpenExecTime 	int32 `json:"totalOpenExecTime"`
	TotalCloseExecTime 	int32 `json:"totalCloseExecTime"`
	ExecOffsetTime 		int32 `json:"execOffsetTime "`
}

// ConfHandler
func ConfHandler (res http.ResponseWriter, req *http.Request) {

	fmt.Println("#################### ConfHandler !!!!!!!")

	///ksn/v2/[L|C|S]/CMD/POST/CONF
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	var jsonData []ConfDevice = make([]ConfDevice, 10)

	json.Unmarshal([]byte(body), &jsonData)

	fmt.Println("jsonData", jsonData)

	reqId := time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))

	var confDevices []*protobuf.ConfDevice = make([]*protobuf.ConfDevice, len(jsonData))

	configPayload := &protobuf.ConfigPayload{}

	configPayload.ReqId = proto.Int64(reqId)
	configPayload.CreatedAt = proto.Int64(reqId)
	// 수동제어
	configPayload.ReqTy = proto.Int32(1)

	for i , json := range jsonData {

		confDevice := &protobuf.ConfDevice{
			DvcId:proto.Int64(json.DvcId),
			FnctngModeCd:proto.Int32(json.FnctngModeCd),
			TotalOpenExecTime:proto.Int32(json.TotalOpenExecTime),
			TotalCloseExecTime:proto.Int32(json.TotalCloseExecTime),
			ExecOffsetTime:proto.Int32(json.ExecOffsetTime),
		}

		fmt.Println(confDevice)
		//
		confDevices[i] = confDevice
	}

	configPayload.ConfDevice = confDevices

	fmt.Println("#################### confPayload : ", configPayload)

	client := mqttlib.NewLocalClient()

	message, err := proto.Marshal(configPayload)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	client.Publish("/ksn/v2/L/CMD/POST/CONF", 0, false, message)
	client.Publish("/ksn/v2/C/CMD/POST/CONF", 0, false, message)
	client.Publish("/ksn/v2/S/CMD/POST/CONF", 0, false, message)

	fmt.Println("############# CONF PUBLISH !!! ")

	client.Disconnect(0)

	resultJson := `
	{
	    "header": {
		"resultCode": 201,
		"resultMessage": "Created"
	    },
	    "body": {
		"item": {
		    "reqId": 3661303731663466,
		    "sndngCd": 10150002
		}
	    }
	}
	`
	res.Header().Add("Content-Type", "Application/json")
	res.Write([]byte(resultJson)) // 웹 브라우저에 응답


}

// ExecHandler
//------------------------------------------------------------------------------

type ExecRoot struct {
	DvcId 	int64 `json:"dvcId"`
	AttrbCd int32 `json:"attrbCd"`
	AttrbVal int32 `json:"attrbVal"`
	AttrbExecs []ExecAttrbExecs `json:"attrbExecs"`
}

type ExecAttrbExecs struct {
	StepSeq int32 `json:"stepSeq"`
	StepDelay int32 `json:"stepDelay"`
	StepFactor int32 `json:"stepFactor"`
}

func ExecHandler (res http.ResponseWriter, req *http.Request) {

	fmt.Println("#################### ExecHandler !!!!!!!")
	log.Println("#################### ExecHandler !!!!!!!")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	var jsonData []ExecRoot

	json.Unmarshal([]byte(body), &jsonData)

	fmt.Println("#################### JSON DATA LENGTH : ", len(jsonData))

	reqId := time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))

	dataPayload := &protobuf.DataPayload{}

	dataPayload.ReqId = proto.Int64(reqId)
	dataPayload.CreatedAt = proto.Int64(reqId)
	// 수동제어
	dataPayload.ReqTy = proto.Int32(1)

	var attributes []*protobuf.Attribute = make([]*protobuf.Attribute, len(jsonData))

	int32ValueType  := protobuf.ValueType_INT32

	for i , execRoot := range jsonData {

		fmt.Println("DVC ID : ", execRoot.DvcId)
		fmt.Println("i : ", i)

		fmt.Println("execRoot.AttrbExecs length : ", len(execRoot.AttrbExecs))

		var execAttributes []*protobuf.ExecAttribute = make([]*protobuf.ExecAttribute, len(execRoot.AttrbExecs))
		////
		fmt.Println("execAttributes : " ,execAttributes)
		//
		for x, attrbExecs := range execRoot.AttrbExecs {
		//
			fmt.Println("############# x : ", x)
			fmt.Println("StepSeq : ", attrbExecs.StepSeq)
			fmt.Println("StepDelay : ", attrbExecs.StepDelay)
			fmt.Println("StepFactor : ",attrbExecs.StepFactor)

			fmt.Println("ssibal 1")

			execAttribute := &protobuf.ExecAttribute{
				StepSeq : proto.Int32(attrbExecs.StepSeq),
				StepDelay : proto.Int32(attrbExecs.StepDelay),
				StepFactor :  &protobuf.AttributeValue {
					Type: &int32ValueType,
					IntValue:proto.Int32(attrbExecs.StepFactor),
				},
			}

			execAttributes[x] = execAttribute
		}

		fmt.Println("execAttributes : " ,execAttributes)

		attribute := &protobuf.Attribute{
			DvcId:proto.Int64(execRoot.DvcId),
			AttrbCd : proto.Int32(execRoot.AttrbCd),
			AttrbVal:&protobuf.AttributeValue{
				Type: &int32ValueType,
				IntValue:proto.Int32(execRoot.AttrbVal),
			},
			ExecAttribute: execAttributes,
		}
		//
		fmt.Println(attribute)

		attributes[i] = attribute
	}

	fmt.Println("attributes : ", attributes)

	dataPayload.Attribute = attributes

	fmt.Println("dataPayload : ", dataPayload)

	client := mqttlib.NewLocalClient()
	pubTopic := "/ksn/v2/C/CMD/EXEC/ACTUATORS"

	message, err := proto.Marshal(dataPayload)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	client.Publish(pubTopic, 0, false, message)

	fmt.Println("############# EXEC PUBLISH !!! ")

	client.Disconnect(0)

	resultJson := `
	{
	    "header": {
		"resultCode": 201,
		"resultMessage": "Created"
	    },
	    "body": {
		"item": {
		    "reqId": ` + strconv.FormatInt(reqId, 10) + `,
		    "sndngCd": 10150002
		}
	    }
	}
	`
	res.Header().Add("Content-Type", "Application/json")
	res.Write([]byte(resultJson)) // 웹 브라우저에 응답

}
//------------------------------------------------------------------------------

// Reset Handler
type ResetStruct struct {
	DvcId 	int64 `json:"dvcId"`
}

func ResetHandler (res http.ResponseWriter, req *http.Request) {

	fmt.Println("#################### ResetHandler !!!!!!!")

	reqId := time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	//var jsonData map[string]interface{}

	var jsonData []ResetStruct = make([]ResetStruct, 10)

	json.Unmarshal([]byte(body), &jsonData)

	fmt.Println("JSON DATA : ", jsonData)

	var attributes []*protobuf.Attribute = make([]*protobuf.Attribute, len(jsonData))

	dataPayload := &protobuf.DataPayload{}

	dataPayload.ReqId = proto.Int64(reqId)
	dataPayload.CreatedAt = proto.Int64(reqId)
	// 수동제어
	dataPayload.ReqTy = proto.Int32(1)

	for i , resetDvcId := range jsonData {

		fmt.Println("dvcId : ", resetDvcId)

		attribute := &protobuf.Attribute{
			DvcId:proto.Int64(resetDvcId.DvcId),
		}

		fmt.Println(attribute)
		//
		attributes[i] = attribute
	}

	fmt.Println("attributes : ", attributes)

	dataPayload.Attribute = attributes

	fmt.Println("dataPayload : ", dataPayload)

	client := mqttlib.NewLocalClient()
	pubTopic := "/ksn/v2/C/CMD/RESET/ACTUATORS"

	message, err := proto.Marshal(dataPayload)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	client.Publish(pubTopic, 0, false, message)

	fmt.Println("############# RESET PUBLISH !!! ")

	client.Disconnect(0)

	resultJson := `
	{
	    "header": {
		"resultCode": 201,
		"resultMessage": "Created"
	    },
	    "body": {
		"item": {
		    "reqId": ` + strconv.FormatInt(reqId, 10) + `,
		    "sndngCd": 10150002
		}
	    }
	}
	`
	res.Header().Add("Content-Type", "Application/json")
	res.Write([]byte(resultJson)) // 웹 브라우저에 응답

}

//------------------------------------------------------------------------------


func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
// DvcListHandler
func DvcListHandler (res http.ResponseWriter, req *http.Request) {

	fmt.Println("#################### DvcListHandler !!!!!!!")

	log.Println("#################### DvcListHandler !!!!!!!")

	b, err := ioutil.ReadFile("./src/krofarm/resources/datas/dvcList.json") // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		log.Fatal(err)
		return
	}

	resultJson := BytesToString(b)

	fmt.Println(resultJson)

	res.Header().Add("Content-Type", "Application/json")
	res.Write([]byte(resultJson)) // 웹 브라우저에 응답
}



// 초기화 메서드
func init() {

	log.SetOutput(&lumberjack.Logger{
		Filename:   "/home/krobis/logs/krofarm-broker.log",
		//Filename:   "d:/logs/krofarm-broker.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     2, // days
	})
}


func main() {


	log.Println("###### Krofarm Broker starting !!!")

	// 채널 생성
	go mqttlib.StartMqttReceiver()

	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")

	// udf 요청
	router.HandleFunc("/csb/iotservice/v2/iottransfer/infocmmnd/endpnts/{endpntId}/post/udf", UdfHandler).Methods("GET", "POST")

	// dvc list 요청
	router.HandleFunc("/csb/iotservice/v2/iotdvc/endpnts/{endpntId}/dvcs/list", DvcListHandler).Methods("GET")

	// 원격제어
	router.HandleFunc("/csb/iotservice/v2/cntrlcmmnd/endpnts/{endpntId}/dvcs/list/exec/actuators", ExecHandler).Methods("GET", "POST")

	// conf 설정
	router.HandleFunc("/csb/iotservice/v2/iottransfer/infocmmnd/endpnts/{endpntId}/dvcs/list/post/conf", ConfHandler).Methods("GET", "POST")

	// reset 제어
	router.HandleFunc("/csb/iotservice/v2/cntrlcmmnd/endpnts/{endpntId}/dvcs/list/reset/actuators", ResetHandler).Methods("GET", "POST")

	http.Handle("/", router)
	http.ListenAndServe(":8888", nil)

	log.Println("###### Krofarm Broker started !!!")



}
