package main

/*
有大规模请求（十万或百万级qps）,我们处理的请求可能不需要立马知道结果,例如数据的打点,文件的上传等等.这时候我们需要异步化处理。
goroutine协同带缓存的管道:这种方法使用了缓冲队列一定程度上了提高了并发,但也是治标不治本,
大规模并发只是推迟了问题的发生时间.当请求速度远大于队列的处理速度时,缓冲区很快被打满,后面的请求一样被堵塞了
job队列+工作池(线程池)
 */




type Job interface {
	Do() error
}



func main() {

}
