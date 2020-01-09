package main

import (
	"rpc/rpcdemo"
	"utils"
	"errors"
)


type DemoService struct {

}

func (d DemoService) Div(args *rpcdemo.Args,result * float64) error {
	//fmt.Println(args)
	if args.Y ==0{
		return errors.New("division by zero")
	}
	*result = float64(args.X)/float64(args.Y)
	return nil
}


//命令行测试：telnet localhost 5566
//{"method":"DemoService.Div","params":[{"X":3,"Y":4}],"id":1}
func main() {
	err:=utils.ServeRpc(":5566",DemoService{})
	if err != nil {
		panic(err)
	}
}
