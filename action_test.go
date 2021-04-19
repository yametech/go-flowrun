package flowrun_test

import (
	"log"
	"testing"

	"github.com/yametech/go-flowrun"
)

func TestActionCreate(t *testing.T) {
	action := &flowrun.Action{
		EchoerUrl:        "http://10.200.100.200:8080",
		Name:             "test",
		InterfaceUrl:     "http://hello.com",
		ReturnStatusList: []string{"Y", "N"},
	}
	action.AddParmas(map[string]interface{}{
		"project": "int",
		"value":   "str",
		"work":    "str",
	})
	fsl := action.Generate()
	log.Println(fsl)
	action.Create(fsl)
}

func TestActionAll(t *testing.T) {
	action := &flowrun.Action{
		EchoerUrl: "http://10.200.100.200:8080",
	}
	log.Println(action.All())
}

func TestActionOne(t *testing.T) {
	action := &flowrun.Action{
		EchoerUrl: "http://10.200.100.200:8080",
	}
	log.Println(action.One("test"))
}

func TestActionDelete(t *testing.T) {
	action := &flowrun.Action{
		EchoerUrl: "http://10.200.100.200:8080",
	}
	action.Delete("test", "zJFP6xmfi1rnwfrhDi1")
}
