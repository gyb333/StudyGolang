package job

//待执行的工作
type Job struct {
	Payload Payload
}

type Payload struct {
	Num int
}

type PayloadCollection struct {
	WindowsVersion  string    `json:"version"`
	Token           string    `json:"token"`
	Payloads        []Payload `json:"data"`
}

//任务channal
var JobQueue chan Job



