package adapayCore

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
)

type Message struct {
	opts                  *MQTT.ClientOptions
	client                MQTT.Client
	topic                 string
	DefaultPublishHandler MessageHandler
}

type MesBody struct {
	Topic     string
	MessageID uint16
	Payload   []byte
}

type MessageHandler func(MesBody)

func Creat(mig *MsgConfig) (*Message, error) {

	u1 := uuid.Must(uuid.NewV4()).String()
	clientStr := mig.Api_key + u1

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(clientStr))
	cipherStr := md5Ctx.Sum(nil)


	clientId := mig.Group_id + "@@@" + hex.EncodeToString(cipherStr)


	username := "Token|" + mig.Access_key + "|" + mig.Instance_id


	reqUrl := mig.Base_url + "/v1/token/apply"
	tokenData, apiError, err := requestToken(reqUrl, mig.MerchSysConfig)


	if err != nil || apiError != nil {
		return nil, errors.New("====== 密码获取失败 ======")
	}


	password := "R|" + tokenData["token"].(string)


	brokerUrl := "tcp://" + mig.Broker_url + ":1883"


	opts := MQTT.NewClientOptions().AddBroker(brokerUrl)

	topic := "topic_crhs_sender/" + mig.Api_key

	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)

	opts.SetAutoReconnect(true)

	opts.SetResumeSubs(true)
	opts.SetKeepAlive(60 * 2 * time.Second)

	m := &Message{
		opts:  opts,
		topic: topic,
	}
	return m, nil
}

func Connect(msg *Message) error {

	var MQMessHandler MQTT.MessageHandler = func(mqClient MQTT.Client, mqMsg MQTT.Message) {
		body := MesBody{
			Topic:     mqMsg.Topic(),
			MessageID: mqMsg.MessageID(),
			Payload:   mqMsg.Payload(),
		}

		msg.DefaultPublishHandler(body)
	}

	msg.opts.SetDefaultPublishHandler(MQMessHandler)

	client := MQTT.NewClient(msg.opts)
	msg.client = client

	if token := msg.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func Subscribe(msg *Message) error {
	if token := msg.client.Subscribe(msg.topic, 0, nil); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (msg *Message) SetDefaultPublishHandler(defaultHandler MessageHandler) *Message {
	msg.DefaultPublishHandler = defaultHandler
	return msg
}

func requestToken(reqUrl string, msc *MerchSysConfig) (map[string]interface{}, *ApiError, error) {
	parma := make(map[string]interface{})

	parma["expire_time"] = 30000000000

	return RequestAdaPay(reqUrl, POST, parma, msc)
}
