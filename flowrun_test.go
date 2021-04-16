package flowrun_test

import (
	"testing"

	"github.com/yametech/go-flowrun"
)

func TestFlowRun(t *testing.T) {
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
	println(flowrun.Generate())
}
