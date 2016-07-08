package mqttpkg

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	UUID "github.com/satori/go.uuid"
	"krofarm/base/config"
	"fmt"
	"strings"
	"log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "/home/krobis/logs/krofarm-broker.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     2, // days
	})
}

func NewLocalClient() MQTT.Client {

	log.Println(">>>>>>>>>>>>>>>>>>>> NewLocalClient 111")

	config := cfg.NewAppCfg()

	log.Println(">>>>>>>>>>>>>>>>>>>> NewLocalClient 222")

	mqttUrl, _ := config.ReadString("local.mqtt.url", "")

	log.Println(">>>>>>>>>>>>>>>>>>>> mqttUrl : ", mqttUrl)

	mqttUrls := strings.Split(mqttUrl, ",")

	log.Println(">>>>>>>>>>>>>>>>>>>> mqttUrls : ", mqttUrls)

	opts := MQTT.NewClientOptions()
	for _, uri := range mqttUrls {
		opts.AddBroker(uri)
	}
	opts.SetClientID(UUID.NewV4().String())

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Println("############### client.IsConnected() : " , client.IsConnected())

	return client
}

func NewPlatformClient() MQTT.Client {

	config := cfg.NewAppCfg()
	mqttUrl, _ := config.ReadString("fms.mqtt.url", "")

	fmt.Println("MQTTURL", mqttUrl)

	opts := MQTT.NewClientOptions().AddBroker(mqttUrl)
	opts.SetClientID(UUID.NewV4().String())
	//opts.SetDefaultPublishHandler(messageArrived)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Println("############### client.IsConnected() : " , client.IsConnected())

	return client
}