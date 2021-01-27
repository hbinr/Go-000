package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// UserClient 客户端
type UserClient struct {
	Name string
	Addr string
	C    chan string
}

// 存储在线用户
var onlineMap map[string]*UserClient

// 创建chan, 传递用户消息
var mesChan = make(chan string)

// HandleConnect 创建用户结构体对象，存入onlineMap；
// 发送用户的登录广播，聊天信息；
// 处理用户查询在线用户，改名，下线，超时退出
func HandleConnect(conn net.Conn) {
	// 断开连接
	defer conn.Close()
	// 创建一个channel，标识用户是否活跃
	isAive := make(chan bool)
	// 创建一个chan，用来标识用户是否退出，true标识用户退出
	isQuit := make(chan bool)
	// 存连接用户信息
	netAddr := conn.RemoteAddr().String()
	userClt := UserClient{
		Name: netAddr,
		Addr: netAddr,
		C:    make(chan string),
	}
	onlineMap[netAddr] = &userClt

	// 创建专门用来给当前用户发送消息的Go程
	go WriteMsgToClt(userClt, conn)

	// 发送用户登录消息到 message 全局通道中。
	mesChan <- MakeMsg(&userClt, " login")

	// 创建一个Go程，专门用来处理用户发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			// 由于用户退出会返回错误，所以要先处理用户退出的情况，后处理错误
			if n == 0 {
				isQuit <- true
				fmt.Printf("检测到客户端[%s]退出\n", userClt.Name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read err:", err)
				return
			}

			msg := string(buf[:n-1]) // 去除最后一位的空格
			// 提取在线用户列表，只输入who命令，则输出当前用户列表。该信息只有当前用户可看
			if msg == "who" && len(msg) == 3 {

				conn.Write([]byte(""))
				for _, user := range onlineMap {
					userInfo := MakeMsg(user, "\n")
					conn.Write([]byte(userInfo))
				}
				// 改名，只输入rename命令
			} else if len(msg) >= 8 && msg[:6] == "rename" { // eg: rename|newname，不允许单独一个空格作为新名字
				newName := strings.Split(msg, "|")[1]
				userClt.Name = newName // 修改用户名
				userInfo := MakeMsg(&userClt, "\n")
				conn.Write([]byte(userInfo))
			} else {
				// 将读取到用户消息广播给所有用户
				mesChan <- MakeMsg(&userClt, msg)
			}
			isAive <- true
		}
	}()

	//	保证不退出
	for {
		// 监听channel上的数据流动
		select {
		case <-isQuit:
			// 关掉子go程
			close(userClt.C)
			// 用户退出，要从onlineMap中删掉此用户。先退出用户，再广播消息
			delete(onlineMap, userClt.Addr)
			// 广播用户退出的消息,即将信息写入到全局chan
			mesChan <- MakeMsg(&userClt, " logout")
			// 当前用户退出，关闭当前用户的Go程
			return

		case <-isAive:
			// 什么都不做，目的就是重置下面的case计时器。因为触发case逻辑后，会返回起始点，重新循环

		// After等待持续时间过去，然后在返回的通道上发送当前时间
		// 超时机制，当前用户超过20s未做任何操作，则强制用户退出
		case <-time.After(time.Second * 20):
			close(userClt.C)
			delete(onlineMap, userClt.Addr)
			mesChan <- MakeMsg(&userClt, " logout")
			return
		}
	}
}

// MakeMsg 构造消息，优化字符串连接
func MakeMsg(clt *UserClient, msg string) string {
	var buf strings.Builder
	buf.WriteString("[")
	buf.WriteString(clt.Addr)
	buf.WriteString("] ")
	buf.WriteString(clt.Name)
	buf.WriteString(": ")
	buf.WriteString(msg)
	return buf.String()
}

// WriteMsgToClt 读取每个用户自带的 channle 上的信息（有manage发送该消息），回写给当前用户
func WriteMsgToClt(clt UserClient, conn net.Conn) {
	// 监听用户的消息通道是否阻塞，不阻塞就说明消息通道中有值了，将该值内容发送给客户端
	for msg := range clt.C {
		conn.Write([]byte(msg + "\n")) // 得加 换行符，缓冲区不刷新，消息会阻塞
		//conn.Write([]byte(msg))
	}

}

// Manager 监听全局的channel message，将读到的信息广播给onlineMap中的所有用户
func Manager() {
	// 初始化用户
	onlineMap = make(map[string]*UserClient)
	// 监听全局message 通道，如果有数据存储到 msg，无则阻塞
	for {
		msg := <-mesChan
		// 遍历map,向每个用户的信息通道发送信息
		for _, v := range onlineMap {
			v.C <- msg
		}
	}
}
func main() {
	// 主Go程，创建连接socket
	listen, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// 记得关闭
	defer listen.Close()

	// 创建管理者Go程，管理全局map和chan
	go Manager()

	//	循环监听客户端连接请求
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
			return
		}
		go HandleConnect(conn)
	}
}
