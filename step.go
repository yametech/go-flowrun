package flowrun

import (
	"log"
	"net/url"
	"path"

	requests "github.com/levigross/grequests"
)

type ResponseStep struct {
	EchoerUrl string
}

func (r *ResponseStep) Response(flowId interface{}, stepName string, ackState string, uuid string, done bool, globalVariables interface{}) bool {
	// >> json data
	data := map[string]interface{}{
		"flowId":   flowId,
		"stepName": stepName,
		"ackState": ackState,
		"uuid":     uuid,
		"done":     done,
	}
	if globalVariables != nil {
		data["globalVariables"] = globalVariables
	}
	ro := &requests.RequestOptions{
		JSON: data,
	}

	// >> request url
	u, _ := url.Parse(r.EchoerUrl)
	u.Path = path.Join(u.Path, "step")
	resp, err := requests.Post(u.String(), ro)

	// >> response
	if err != nil {
		log.Fatalln("发送echoer错误: ", err.Error())
		return false
	}
	if resp.StatusCode != 200 {
		log.Println("echoer : ", resp.String())
		return false
	}
	return true
}
