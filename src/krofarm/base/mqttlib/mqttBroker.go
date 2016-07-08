package mqttpkg

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	//UUID "github.com/satori/go.uuid"
	"os"
	"krofarm/base/protobuf"
	_ "krofarm/base/const"
	"log"
	"strconv"
	"encoding/json"
	"krofarm/base/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

var messageArrived MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	log.Println("TOPIC: %s\n", msg.Topic())
	//log.Println("MSG: %s\n", msg.Payload())

	config := cfg.NewAppCfg()
	prjctNo, _  := config.ReadString("prjctNo", "")
	endpntId, _ := config.ReadString("endpntId", "")

	var pubTopic string

	// SENSORS나 ACTUATOR 데이터 인 경우...
	if msg.Topic() == "/ksn/v2/S/EVT/DATA/SENSORS" ||
	   msg.Topic() == "/ksn/v2/C/EVT/DATA/ACTUATORS" {

		dataPayload := &protobuf.DataPayload{}

		err := proto.Unmarshal(msg.Payload(), dataPayload)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		var dvcId int64
		var attrbVal string
		var attrbCd int32
		var createdAt int64
		var attrbStatCd int32
		var reqId int64
		var reqTy int32
		var curExecStep int32

		createdAt = dataPayload.GetCreatedAt()
		reqId = dataPayload.GetReqId()
		reqTy = dataPayload.GetReqTy()

		for _ , attribute := range dataPayload.GetAttribute() { // i에는 인덱스, value에는 배열 요소의 값이 들어감

			dvcId   	= attribute.GetDvcId()
			attrbCd 	= attribute.GetAttrbCd()
			attrbStatCd 	= attribute.GetAttrbStatCd()
			curExecStep	= attribute.GetCurExecStep()

			switch attribute.GetAttrbVal().GetType() {
				case protobuf.ValueType_DOUBLE :
					attrbVal = strconv.FormatFloat(attribute.GetAttrbVal().GetDoubleValue(), 'f', -1, 32)
				case protobuf.ValueType_INT32 :
					attrbVal = strconv.FormatInt(int64(attribute.GetAttrbVal().GetIntValue()), 10)
				case protobuf.ValueType_INT64 :
					attrbVal = strconv.FormatInt(attribute.GetAttrbVal().GetLongValue(), 10)
			}

			fmt.Println("dvcId : " , dvcId)
			fmt.Println("attrbVal : " , attrbVal)
			fmt.Println("attrbCd : " , attrbCd)
			fmt.Println("createdAt : " , createdAt)
			fmt.Println("attrbStatCd : " , attrbStatCd)

			fmt.Println("reqId : " , reqId)
			fmt.Println("reqTy : " , reqTy)
			fmt.Println("curExecStep : " , curExecStep)

			jsonData := make(map[string]interface{})

			jsonData["prjctNo"] = prjctNo
			jsonData["endpntId"] = endpntId
			jsonData["dvcId"] = dvcId
			jsonData["attrbVal"] = attrbVal
			jsonData["attrbCd"] = attrbCd
			jsonData["createdAt"] = createdAt
			jsonData["attrbStatCd"] = attrbStatCd
			jsonData["reqId"] = reqId
			jsonData["reqTy"] = reqTy
			jsonData["curExecStep"] = curExecStep

			jsonString, _ := json.Marshal(jsonData)

			log.Println(string(jsonString))

			client := NewPlatformClient()

			if msg.Topic() == "/ksn/v2/S/EVT/DATA/SENSORS" {
				pubTopic = "/kcsb/v2/" + prjctNo + "/" + endpntId + "/EVT/DATA/SENSORS/json"
			} else if msg.Topic() == "/ksn/v2/S/EVT/DATA/ACTUATORS" {
				pubTopic = "/kcsb/v2/" + prjctNo + "/" + endpntId + "/EVT/DATA/ACTUATORS/json"
			}

			fmt.Println("pubTopic : " + pubTopic)

			token := client.Publish(pubTopic, 0, false, jsonString)
			token.Wait()

			client.Disconnect(0)
		}
	}

	// STATUS
	if msg.Topic() == "/ksn/v2/S/EVT/DATA/STATUS" {

		statusPayload := &protobuf.StatusPayload{}

		err := proto.Unmarshal(msg.Payload(), statusPayload)
		if err != nil {
			log.Fatal("statusPayload unmarshaling error: ", err)
		}

		pubTopic = "/kcsb/v2/" + prjctNo + "/" + endpntId + "/EVT/DATA/STATUS/json"

		var createdAt int64
		var reqId int64
		var reqTy int32
		var subDvcStatChg int32
		var dvcId int64
		var commStatCd int32
		var commStatMssage string
		var fnctngStatCd int32
		var fnctngStatMssage string
		var cntrlStat int32
		var dlgatStat int32

		createdAt = statusPayload.GetCreatedAt()
		reqId = statusPayload.GetReqId()
		reqTy = statusPayload.GetReqTy()
		subDvcStatChg = statusPayload.GetSubDvcStatChg()

		fmt.Println("createdAt", createdAt)
		fmt.Println("reqId", reqId)
		fmt.Println("reqTy", reqTy)
		fmt.Println("subDvcStatChg", subDvcStatChg)

		result := make([]map[string]interface{}, len(statusPayload.GetStatDevice()), 10)

		for i , attribute := range statusPayload.GetStatDevice() {
			// i에는 인덱스, value에는 배열 요소의 값이 들어감

			dvcId = attribute.GetDvcId()
			commStatCd = attribute.GetCommStatCd()
			commStatMssage = attribute.GetCommStatMssage()
			fnctngStatCd = attribute.GetFnctngStatCd()
			fnctngStatMssage = attribute.GetFnctngStatMssage()
			cntrlStat = attribute.GetCntrlStat()
			dlgatStat = attribute.GetDlgatStat()

			jsonData := make(map[string]interface{})

			jsonData["prjctNo"] = prjctNo
			jsonData["endpntId"] = endpntId
			jsonData["dvcId"] = dvcId
			jsonData["commStatCd"] = commStatCd
			jsonData["commStatMssage"] = commStatMssage
			jsonData["commStatCreatedAt"] = createdAt
			jsonData["fnctngStatCd"] = fnctngStatCd
			jsonData["fnctngStatMssage"] = fnctngStatMssage
			jsonData["fnctngStatCreatedAt"] = createdAt
			jsonData["cntrlStat"] = cntrlStat
			jsonData["dlgatStat"] = dlgatStat


			result[i] = jsonData

		}

		fmt.Println("########################## STATUS PAYLOAD START ")

		jsonString, _ := json.Marshal(result)

		log.Println(string(jsonString))

		fmt.Println("########################## STATUS PAYLOAD END ")

		client := NewPlatformClient()

		fmt.Println("pubTopic : " + pubTopic)

		token := client.Publish(pubTopic, 0, false, jsonString)
		token.Wait()

		client.Disconnect(0)

	}

	// UDF
	if msg.Topic() == "/ksn/v2/L/CMD/POST/UDF" {

		udfPayload := &protobuf.UDFPayload{}

		err := proto.Unmarshal(msg.Payload(), udfPayload)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		fnctTyCd := udfPayload.GetFnctTyCd()
		udfNm := udfPayload.GetUdfNm()
		createdAt := udfPayload.GetCreatedAt()


		fmt.Println("fnctTyCd", fnctTyCd)
		fmt.Println("udfNm", udfNm)
		fmt.Println("createdAt", createdAt)
		fmt.Println("udfBody", createdAt)

	}

	// UDM
	if msg.Topic() == "/kgw/v2/L/EVT/MSG/UDM" {

		udmPayload := &protobuf.UDMPayload{}

		err := proto.Unmarshal(msg.Payload(), udmPayload)
		if err != nil {
			log.Fatal("udmPayload unmarshaling error: ", err)
		}

		var dvcId 	int64
		var createdAt 	int64
		var mssageTyCd 	int32
		var udmNm 	string
		var udmBody 	string
		var udmTy 	string

		dvcId 		= udmPayload.GetDvcId()
		createdAt 	= udmPayload.GetCreatedAt()
		mssageTyCd 	= udmPayload.GetMssageTyCd()
		udmNm		= udmPayload.GetUdmNm()
		udmBody		= udmPayload.GetUdmBody()
		udmTy		= udmPayload.GetUdmTy()

		jsonData := make(map[string]interface{})

		jsonData["prjctNo"] = prjctNo
		jsonData["endpntId"] = endpntId
		jsonData["dvcId"] = dvcId
		jsonData["createdAt"] = createdAt
		jsonData["mssageTyCd"] = mssageTyCd
		jsonData["udmNm"] = udmNm
		jsonData["udmBody"] = udmBody
		jsonData["udmTy"] = udmTy

		jsonString, _ := json.Marshal(jsonData)

		log.Println(string(jsonString))

		client := NewPlatformClient()

		pubTopic = "/kcsb/v2/" + prjctNo + "/" + endpntId + "/EVT/MSG/UDM/json"

		fmt.Println("pubTopic : " + pubTopic)

		token := client.Publish(pubTopic, 0, false, jsonString)
		token.Wait()

		client.Disconnect(0)
	}
}

func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "/home/krobis/logs/krofarm-broker.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     2, // days
	})
}

func StartMqttReceiver() {

	log.Println("MQTT RECEIVER START !!!!")

	client := NewLocalClient()

	for {

		if token := client.Subscribe("/#", 0, messageArrived); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}

	log.Println("MQTT RECEIVER END !!!!")
}