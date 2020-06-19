package main

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

// Target is a ping
type Target struct {
	Host string
	Port int

	Counter  int
	Interval time.Duration
	Timeout  time.Duration
}

func (target Target) String() string {
	return fmt.Sprintf("%s:%d", target.Host, target.Port)
}

// Pinger is a ping interface
type Pinger interface {
	Start() <-chan struct{}
	Stop()
	Result() *Result
	SetTarget(target *Target)
}

// Ping is a ping interface
type Ping interface {
	Start() <-chan struct{}

	Host() string
	Port() int
	Counter() int

	Stop()

	Result() Result
}

// Result ...
type Result struct {
	Counter        int
	SuccessCounter int
	Target         *Target

	MinDuration   time.Duration
	MaxDuration   time.Duration
	TotalDuration time.Duration
}

// Avg return the average time of ping
func (result Result) Avg() time.Duration {
	if result.SuccessCounter == 0 {
		return 0
	}
	return result.TotalDuration / time.Duration(result.SuccessCounter)
}

// Failed return failed counter
func (result Result) Failed() int {
	return result.Counter - result.SuccessCounter
}

func (result Result) String() string {
	const resultTpl = `
Ping statistics {{.Target}}
	{{.Counter}} probes sent.
	{{.SuccessCounter}} successful, {{.Failed}} failed.
Approximate trip times:
	Minimum = {{.MinDuration}}, Maximum = {{.MaxDuration}}, Average = {{.Avg}}`
	t := template.Must(template.New("result").Parse(resultTpl))
	res := bytes.NewBufferString("")
	t.Execute(res, result)
	return res.String()
}
