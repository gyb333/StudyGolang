package millionlevel

import (
	"net/http"
	"encoding/json"
	"io"
)

const MaxLength  = 1024*1024
const MAX_QUEUE  = 1024*1024

var Queue chan Payload

func init() {
	Queue = make(chan Payload, MAX_QUEUE)
}


func payloadHandler1(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read the body into a string for json decoding
	var content = &PayloadCollection{}
	err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range content.Payloads {
		//1.只是试图将作业处理并行化为一个简单的goroutine
		//无法控制我们产生的go例程。当然，由于我们每分钟收到一百万个POST请求，因此该代码崩溃并很快烧毁。
		go payload.UploadToS3()   // 对于中等负载，这可能适合大多数人，但很快就证明了这种方法在大规模情况下效果不佳。

		//2.只将作业缓存在通道队列中
		Queue <- payload

	}

	w.WriteHeader(http.StatusOK)
}


func StartProcessor() {
	for {
		select {
		case job := <-Queue:
			//这种方法并没有给我们带来任何好处，我们已经将有缺陷的并发与缓冲队列进行了交换，而这只是在推迟问题。
			//因此我们的缓冲通道迅速达到了其极限并阻塞了请求处理程序的能力排队更多物品。
			job.UploadToS3()  // <-- STILL NOT GOOD

		}
	}
}
