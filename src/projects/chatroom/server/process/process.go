package process


import (
	"net"
	"log"
	"io"
	"projects/chatroom/server/processor"
	"projects/chatroom/common/message"
	"projects/chatroom/common/utils"
	)

type Process struct {
	Conn net.Conn
}

func (this *Process) ProcessMsg(msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType :
		//处理登录登录
		//创建一个UserProcess实例

		up := &processor.UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessLogin(msg)

	case message.RegisterMsgType :
		//处理注册
		up := &processor.UserProcess{
			Conn : this.Conn,
		}
		err = up.ServerProcessRegister(msg) // type : data
	case message.SmsMsgType :
		//创建一个SmsProcess实例完成转发群聊消息.
		smsProcess := &processor.SmsProcess{}
		smsProcess.SendGroupMes(msg)
	default :
		log.Println("消息类型不存在，无法处理...")
	}

	return
}
func (this *Process) ProcessData() (err error) {
	//循环的客户端发送的信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(), 返回Message, Err
		//创建一个Transfer 实例完成读包任务
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				log.Println("客户端退出，服务器端也退出..")
				return err
			} else {
				log.Println("readPkg err=", err)
				return err
			}

		}
		err = this.ProcessMsg(&msg)
		if err != nil {
			return err
		}
	}

}