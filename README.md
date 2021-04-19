# go-flowrun

Finite-state machine [echoer](https://github.com/yametech/echoer "Markdown") Go  SDK
---------

Action Example
---------

```go
action := &flowrun.Action{
  EchoerUrl:        "http://localhost:8080",
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
```

```go
action := &flowrun.Action{
  EchoerUrl: "http://localhost:8080",
}
log.Println(action.All())
```

```go
action := &flowrun.Action{
  EchoerUrl: "http://localhost:8080",
}
log.Println(action.One("test"))
```

```go
action := &flowrun.Action{
  EchoerUrl: "http://localhost:8080",
}
action.Delete("test", "zJFP6xmfi1rnwfrhDi1")
```

FlowRun Example
---------

```go
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
flowrun.Create(fsl)
```

Result
```
flow_run  test
  step step_name1 => (SUCCESS->a | FAIL->done) {action="action_name1"; args=(project="https://github.com/yametech/compass.git", version=3);};
flow_run_end
```

```go
flowrun := &flowrun.FlowRun{
  EchoerUrl: "http://localhost:8080",
}
log.Println(flowrun.All())
```

```go
flowrun := &flowrun.FlowRun{
  EchoerUrl: "http://localhost:8080",
}
log.Println(flowrun.One("apollo_run_APYCpQEd"))
```

```go
flowrun := &flowrun.FlowRun{
  EchoerUrl: "http://localhost:8080",
}
log.Println(flowrun.Delete("apollo_run_APYCpQEd", "uuid:NcfT73nfiLtnwfrhDi1"))
```