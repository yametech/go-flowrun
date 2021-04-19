package flowrun_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/yametech/go-flowrun"
)

func TestFlowRunCreate(t *testing.T) {
	flowrun := &flowrun.FlowRun{
		Name: "test",
	}
	steps := map[string]string{
		"SUCCESS": "a", "FAIL": "done",
	}
	args := map[string]interface{}{
		"project": "https://github.com/yametech/compass.git", "version": 3,
	}
	flowrun.AddStep("step_name1", steps, "action_name1", args)
	fsl := flowrun.Generate()
	fmt.Println(fsl)
	flowrun.Create(fsl)
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
