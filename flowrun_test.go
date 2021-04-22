package flowrun_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/yametech/go-flowrun"
)

func TestFlowRunCreate(t *testing.T) {
	flowrun := &flowrun.FlowRun{
		EchoerUrl: "http://10.200.65.192:8080",
		Name:      "pipeline_1618890450",
	}
	steps := map[string]string{
		"SUCCESS": "done", "FAIL": "done",
	}
	args := map[string]interface{}{
		"codeType":    "web",
		"output":      "registry-d.ym/fengdr",
		"serviceName": "devopsui",
		"projectFile": "",
		"projectPath": "",
		"retryCount":  15.0,
		"branch":      "master",
		"commitId":    "devopsui-2fcba0622839119ef5d84472eab353e68e8705dd",
		"gitUrl":      "http://git.ym/fengdr/Devops-prototype.git",
	}
	flowrun.AddStep("pipeline_y1ui6Orfi6rvauYFNZ1", steps, "artifactoryCI", args)
	fsl := flowrun.Generate()
	fmt.Println(fsl)
	// log.Println(flowrun.Create(fsl))
}

func TestFlowRunAll(t *testing.T) {
	flowrun := &flowrun.FlowRun{
		EchoerUrl: "http://localhost:8080",
	}
	log.Println(flowrun.All())
}

func TestFlowRunOne(t *testing.T) {
	flowrun := &flowrun.FlowRun{
		EchoerUrl: "http://localhost:8080",
	}
	log.Println(flowrun.One("apollo_run_APYCpQEd"))
}

func TestFlowRunDelete(t *testing.T) {
	flowrun := &flowrun.FlowRun{
		EchoerUrl: "http://localhost:8080",
	}
	log.Println(flowrun.Delete("apollo_run_APYCpQEd", "uuid:NcfT73nfiLtnwfrhDi1"))
}
