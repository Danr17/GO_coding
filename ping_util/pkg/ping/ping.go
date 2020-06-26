package ping

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

// Target is a ping
type Target struct {
	Host     string
	HostName string
	Proto    string
	Port     int

	Counter    int
	Interval   time.Duration
	Timeout    time.Duration
	Privileged bool
}

func (target Target) String() string {
	return fmt.Sprintf("%s:%d", target.Host, target.Port)
}

// Result holds ping results
type Result struct {
	Counter        int
	SuccessCounter int
	Status         bool
	LastSeen       string
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
