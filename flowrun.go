package flowrun

import (
	"fmt"
	"log"
	"net/url"
	"path"
	"reflect"
	"strings"

	requests "github.com/levigross/grequests"
)

type Step struct {
	Name          string
	Action        string
	FlowDirection map[string]string
	ActionParams  map[string]interface{}
	flowFmt       string
	argsFmt       string
}

func (s *Step) flowDirectionFmt() {
	s.flowFmt = ""
	for key, value := range s.FlowDirection {
		if s.flowFmt != "" {
			s.flowFmt += " | "
		}
		s.flowFmt += fmt.Sprintf("%s->%s", key, value)
	}
}

func (s *Step) actionParamsFmt() {
	s.argsFmt = ""
	for key, value := range s.ActionParams {
		if s.argsFmt != "" {
			s.argsFmt += ", "
		}
		switch reflect.TypeOf(value).Kind() {
		case reflect.Map:
			value = fmt.Sprintf("`%s`", fmt.Sprint(value))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = fmt.Sprintf(`%d`, value)
		case reflect.Float32, reflect.Float64:
			if float64(int(value.(float64))) == float64(value.(float64)) {
				value = fmt.Sprintf(`%d`, int(value.(float64)))
			} else {
				value = fmt.Sprintf(`"%s"`, fmt.Sprint(value))
			}
		case reflect.String:
			if strings.Contains(value.(string), `"`) {
				value = fmt.Sprintf("`%s`", value)
			} else {
				value = fmt.Sprintf(`"%s"`, value)
			}
		default:
			value = fmt.Sprintf(`"%s"`, value)
		}
		s.argsFmt += fmt.Sprintf("%s=%s", key, value)
	}
}

func (s *Step) generate() string {
	s.flowDirectionFmt()
	s.actionParamsFmt()
	result := fmt.Sprintf(`step %s => (%s) {action="%s"; args=(%s);};`, s.Name, s.flowFmt, s.Action, s.argsFmt)
	return result
}

type FlowRun struct {
	EchoerUrl string
	Name      string
	steps     []Step
}

func (f *FlowRun) AddStep(flowRunName string, flowDirection map[string]string, actionName string, actionParams map[string]interface{}) {
	f.steps = append(f.steps, Step{
		Name:          flowRunName,
		Action:        actionName,
		FlowDirection: flowDirection,
		ActionParams:  actionParams,
	})
}

func (f *FlowRun) Generate() string {
	r := ""
	for _, step := range f.steps {
		p := step.generate()
		r += fmt.Sprintf("  %s\n", p)
	}
	return fmt.Sprintf("flow_run  %s\n%sflow_run_end", f.Name, r)
}

func (f *FlowRun) Create(fsl string) bool {
	ro := &requests.RequestOptions{
		JSON: map[string]string{
			"data": fsl,
		},
	}
	// >> request url
	u, _ := url.Parse(f.EchoerUrl)
	u.Path = path.Join(u.Path, "flowrun")
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

func (f *FlowRun) Delete(name string, uuid string) bool {
	// >> request url
	u, _ := url.Parse(f.EchoerUrl)
	u.Path = path.Join(u.Path, "flowrun", name, uuid)
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

func (f *FlowRun) All() interface{} {
	// >> request url
	u, _ := url.Parse(f.EchoerUrl)
	u.Path = path.Join(u.Path, "flowrun")
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

func (f *FlowRun) One(name string) interface{} {
	// >> request url
	u, _ := url.Parse(f.EchoerUrl)
	u.Path = path.Join(u.Path, "flowrun", name)
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
