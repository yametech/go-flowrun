package flowrun

import (
	"fmt"
	"log"
	"net/url"
	"path"

	requests "github.com/levigross/grequests"
)

type Action struct {
	EchoerUrl        string
	Name             string
	InterfaceUrl     string
	ReturnStatusList []string
	params           string
	returnStatusFmt  string
}

func (a *Action) AddParmas(data map[string]interface{}) {
	a.params = "("
	lens := len(data)
	index := 0
	for key, value := range data {
		switch value {
		case "int":
			a.params += fmt.Sprintf(`%s %s`, "int", key)
		default:
			a.params += fmt.Sprintf(`%s %s`, "str", key)
		}
		index += 1
		if index < lens {
			a.params += ", "
		}
	}
	a.params += ")"
}

func (a *Action) actionReturnStatusFmt() {
	a.returnStatusFmt = ""
	if len(a.ReturnStatusList) == 0 {
		a.ReturnStatusList = []string{"SUCCESS", "FAIL"}
	}
	lens := len(a.ReturnStatusList)
	index := 0
	for _, s := range a.ReturnStatusList {
		a.returnStatusFmt += s
		index += 1
		if index < lens {
			a.returnStatusFmt += " | "
		}
	}
}

func (a *Action) Generate() string {
	a.actionReturnStatusFmt()
	data := fmt.Sprintf(`
action %s 
  addr = "%s";
  method = http;
  args = %s;
  return = (%s);
action_end`, a.Name, a.InterfaceUrl, a.params, a.returnStatusFmt)
	return data
}

func (a *Action) Create(fsl string) bool {
	ro := &requests.RequestOptions{
		JSON: map[string]string{
			"data": fsl,
		},
	}
	// >> request url
	u, _ := url.Parse(a.EchoerUrl)
	u.Path = path.Join(u.Path, "action")
	resp, err := requests.Post(u.String(), ro)

	// >> response
	if err != nil {
		log.Println("发送echoer错误: ", err.Error())
		return false
	}
	if resp.StatusCode != 200 {
		log.Println("echoer: ", resp.String())
		return false
	}
	return true
}

func (a *Action) Delete(name string, uuid string) bool {
	// >> request url
	u, _ := url.Parse(a.EchoerUrl)
	u.Path = path.Join(u.Path, "action", name, uuid)
	resp, err := requests.Delete(u.String(), nil)

	// >> response
	if err != nil {
		log.Println("发送echoer错误: ", err.Error())
		return false
	}
	if resp.StatusCode != 200 {
		log.Println("echoer : ", resp.String())
		return false
	}
	return true
}

func (a *Action) All() interface{} {
	// >> request url
	u, _ := url.Parse(a.EchoerUrl)
	u.Path = path.Join(u.Path, "action")
	resp, err := requests.Get(u.String(), nil)

	// >> response
	if err != nil {
		log.Println("发送echoer错误: ", err.Error())
		return nil
	}
	if resp.StatusCode != 200 {
		log.Println("echoer: ", resp.String())
		return nil
	}

	// >> decode
	var result interface{}
	if err := resp.JSON(&result); err != nil {
		log.Println("echoer 序列化错误: ", err.Error())
		return nil
	}
	return result
}

func (a *Action) One(name string) interface{} {
	// >> request url
	u, _ := url.Parse(a.EchoerUrl)
	u.Path = path.Join(u.Path, "action", name)
	resp, err := requests.Get(u.String(), nil)

	// >> response
	if err != nil {
		log.Println("发送echoer错误: ", err.Error())
		return nil
	}
	if resp.StatusCode != 200 {
		log.Println("echoer : ", resp.String())
		return false
	}

	// >> decode
	var result interface{}
	if err := resp.JSON(&result); err != nil {
		log.Println("echoer 序列化错误: ", err.Error())
		return nil
	}
	return result
}
