# go-flowrun

Finite-state machine [echoer](https://github.com/yametech/echoer "Markdown") Go  SDK
---------

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
flowrun.Generate()
```

Result
```
flow_run  test
  step step_name1 => (SUCCESS->a | FAIL->done) {action="action_name1"; args=(project="https://github.com/yametech/compass.git", version=3);};
flow_run_end
```
