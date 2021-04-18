package flowrun

import (
	"fmt"
	"reflect"
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
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = fmt.Sprintf(`%d`, value)
		case reflect.Float32, reflect.Float64:
			value = fmt.Sprintf(`%f`, value)
		default:
			value = fmt.Sprintf(`"%s"`, value.(string))
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
	Name  string
	steps []Step
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
