package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	mqttlib "krofarm/base/mqttlib"
//MQTT "github.com/eclipse/paho.mqtt.golang"
	"krofarm/base/protobuf"
	"krofarm/base/const"
	"github.com/golang/protobuf/proto"
	//"time"
	"time"
	"log"
)

// 내부 센서 정보 보내기...
func internalSensorTask() {

	fmt.Println("######################### internalSensorTask.")

	doubleValueType := protobuf.ValueType_DOUBLE
	int32ValueType  := protobuf.ValueType_INT32

	var dvcId int64
	// DEVICE ID
	dvcId = 3735353762356533

	// 내부온도
	attribute01 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.IN_TEMPERATURE),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &doubleValueType,
			DoubleValue:proto.Float64(14),
		},
	}

	// 내부 습도
	attribute02 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.IN_HUMIDITY),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &doubleValueType,
			DoubleValue:proto.Float64(0),
		},
	}

	// 내부 CO2
	attribute03 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.CO2),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &int32ValueType,
			IntValue:proto.Int32(120),
		},
	}

	// dataPayLoad
	dataPayload := &protobuf.DataPayload{
		CreatedAt:proto.Int64(time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))),
		Attribute: []*protobuf.Attribute {
			attribute01,
			attribute02,
			attribute03,
		},
	}

	fmt.Println("dataPayload", dataPayload)

	client := mqttlib.NewLocalClient()

	pubTopic := "/kgw/v2/S/EVT/DATA/SENSORS"

	message, err := proto.Marshal(dataPayload)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	token := client.Publish(pubTopic, 0, false, message)
	token.Wait()

	client.Disconnect(0)
}

// 배관온도 센서 정보
func heatPipeSensorTask() {

	fmt.Println("######################### heatPipeSensorTask.")

	doubleValueType := protobuf.ValueType_DOUBLE
	//int32ValueType  := protobuf.ValueType_INT32

	client := mqttlib.NewLocalClient()
	pubTopic := "/kgw/v2/S/EVT/DATA/SENSORS"

	var dvcId int64
	// DEVICE ID
	dvcId = 6662396530643836

	// 주배관 온도
	attribute01 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.HEAT_PIPE),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &doubleValueType,
			DoubleValue:proto.Float64(45),
		},
	}

	// dataPayLoad
	dataPayload01 := &protobuf.DataPayload{
		CreatedAt:proto.Int64(time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))),
		Attribute: []*protobuf.Attribute {
			attribute01,
		},
	}


	fmt.Println("dataPayload", dataPayload01)

	message, err := proto.Marshal(dataPayload01)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	token := client.Publish(pubTopic, 0, false, message)
	token.Wait()


	dvcId = 6565393539653961

	// 회수온도
	attribute02 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.HEAT_PIPE),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &doubleValueType,
			DoubleValue:proto.Float64(18),
		},
	}

	// dataPayLoad
	dataPayload02 := &protobuf.DataPayload{
		CreatedAt:proto.Int64(time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))),
		Attribute: []*protobuf.Attribute {
			attribute02,
		},
	}

	message02, err02 := proto.Marshal(dataPayload02)
	if err02 != nil {
		log.Fatal("marshaling error: ", err)
	}

	token02 := client.Publish(pubTopic, 0, false, message02)
	token02.Wait()

	dvcId = 6666346664663763

	// 보일러온도
	attribute03 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.HEAT_PIPE),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &doubleValueType,
			DoubleValue:proto.Float64(80),
		},
	}

	// dataPayLoad
	dataPayload03 := &protobuf.DataPayload{
		CreatedAt:proto.Int64(time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))),
		Attribute: []*protobuf.Attribute {
			attribute03,
		},
	}

	message03, err03 := proto.Marshal(dataPayload03)
	if err03 != nil {
		log.Fatal("marshaling error: ", err)
	}

	token03 := client.Publish(pubTopic, 0, false, message03)
	token03.Wait()

	client.Disconnect(0)
}



// 외부 센서 정보 Task
func externalSensorTask() {

	fmt.Println("######################### externalSensorTask.")

	doubleValueType := protobuf.ValueType_DOUBLE
	int32ValueType  := protobuf.ValueType_INT32

	var dvcId int64
	// DEVICE ID
	dvcId = 3239643463376332

	// 일사량
	attribute01 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.OUT_RADIATION),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &int32ValueType,
			IntValue:proto.Int32(0),
		},
	}

	// 외부온도
	attribute02 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.OUT_TEMPERATURE),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &doubleValueType,
			DoubleValue:proto.Float64(12),
		},
	}

	// 외부습도
	attribute03 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.OUT_HUMIDITY),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &doubleValueType,
			DoubleValue:proto.Float64(70),
		},
	}

	// 풍향
	attribute04 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.WIND_DIRECTION),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &doubleValueType,
			DoubleValue:proto.Float64(270),
		},
	}

	// 풍속
	attribute05 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.WIND_SPEED),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &doubleValueType,
			DoubleValue:proto.Float64(0),
		},
	}

	// 감우
	attribute06 := &protobuf.Attribute{
		DvcId:proto.Int64(dvcId),
		AttrbCd:proto.Int32(constants.AttrbCd.RAIN_LEVEL),
		AttrbStatCd:proto.Int32(constants.AttrbStatCd.SENSOR_NORMAL),
		AttrbVal:&protobuf.AttributeValue{
			Type: &int32ValueType,
			IntValue:proto.Int32(0),
		},
	}

	// dataPayLoad
	dataPayload := &protobuf.DataPayload{
		CreatedAt:proto.Int64(time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))),
		Attribute: []*protobuf.Attribute {
			attribute01,
			attribute02,
			attribute03,
			attribute04,
			attribute05,
			attribute06,
		},
	}

	fmt.Println("dataPayload", dataPayload)

	client := mqttlib.NewLocalClient()

	pubTopic := "/kgw/v2/S/EVT/DATA/SENSORS"

	message, err := proto.Marshal(dataPayload)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	token := client.Publish(pubTopic, 0, false, message)
	token.Wait()

	client.Disconnect(0)
}

// STATUS 센서 정보 Task
func statusTask() {

	fmt.Println("######################### statusTask.")

	// 웨더 스테이션
	statDevice01 := &protobuf.StatDevice{

		DvcId:proto.Int64(3239643463376332),
		CommStatCd:proto.Int32(constants.CommStatCd.NORMAL),
		FnctngStatCd:proto.Int32(constants.FnctngStatCd.NORMAL),
		CntrlStat:proto.Int32(1),
	}

	// 내부 센서
	statDevice02 := &protobuf.StatDevice{
		DvcId:proto.Int64(3234333038616463 ),
		CommStatCd:proto.Int32(constants.CommStatCd.NORMAL),
		FnctngStatCd:proto.Int32(constants.FnctngStatCd.NORMAL),
		CntrlStat:proto.Int32(1),
	}

	// 배관온도1
	statDevice03 := &protobuf.StatDevice{
		DvcId:proto.Int64(6662396530643836),
		CommStatCd:proto.Int32(constants.CommStatCd.NORMAL),
		FnctngStatCd:proto.Int32(constants.FnctngStatCd.NORMAL),
		CntrlStat:proto.Int32(1),
	}

	// 배관온도2
	statDevice04 := &protobuf.StatDevice{
		DvcId:proto.Int64(6565393539653961),
		CommStatCd:proto.Int32(constants.CommStatCd.NORMAL),
		FnctngStatCd:proto.Int32(constants.FnctngStatCd.NORMAL),
		CntrlStat:proto.Int32(1),
	}

	// 배관온도3
	statDevice05 := &protobuf.StatDevice{
		DvcId:proto.Int64(6666346664663763),
		CommStatCd:proto.Int32(constants.CommStatCd.NORMAL),
		FnctngStatCd:proto.Int32(constants.FnctngStatCd.NORMAL),
		CntrlStat:proto.Int32(1),
	}

	statusPayLoad := &protobuf.StatusPayload{
		CreatedAt:proto.Int64(time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))),
		StatDevice:[]*protobuf.StatDevice {
			statDevice01,
			statDevice02,
			statDevice03,
			statDevice04,
			statDevice05,
		},
	}

	fmt.Println("statusPayLoad", statusPayLoad)

	client := mqttlib.NewLocalClient()

	pubTopic := "/kgw/v2/S/EVT/DATA/STATUS"

	message, err := proto.Marshal(statusPayLoad)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	token := client.Publish(pubTopic, 0, false, message)
	token.Wait()

	client.Disconnect(0)
}

func init() {

	fmt.Println("FUCK YOU INIT !!!")
}

func main() {
	// Do jobs with params
	//gocron.Every(1).Second().Do(taskWithParams, 1, "hello")

	// Do jobs without params
	gocron.Every(1).Seconds().Do(internalSensorTask)
	gocron.Every(1).Seconds().Do(externalSensorTask)
	//gocron.Every(30).Seconds().Do(statusTask)
	//gocron.Every(30).Seconds().Do(heatPipeSensorTask)
/*	gocron.Every(2).Seconds().Do(task)
	gocron.Every(1).Minute().Do(task)
	gocron.Every(2).Minutes().Do(task)
	gocron.Every(1).Hour().Do(task)
	gocron.Every(2).Hours().Do(task)
	gocron.Every(1).Day().Do(task)
	gocron.Every(2).Days().Do(task)

	// Do jobs on specific weekday
	gocron.Every(1).Monday().Do(task)
	gocron.Every(1).Thursday().Do(task)

	// function At() take a string like 'hour:min'
	gocron.Every(1).Day().At("10:30").Do(task)
	gocron.Every(1).Monday().At("18:30").Do(task)

	// remove, clear and next_run
	_, time := gocron.NextRun()
	fmt.Println(time)

	gocron.Remove(task)
	gocron.Clear()

	// function Start start all the pending jobs
	<- gocron.Start()

	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(task)
	<- s.Start()*/

	//gocron.Clear()

	<- gocron.Start()

}
