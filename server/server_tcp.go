package main

import (
    "fmt"
    "net"
    "net/rpc"
    "./zk"
    "os"
)

var S string

type MyRPC int

func (r *MyRPC) HelloRPC(S string, reply *string) error {
    fmt.Println(S)
    *reply = "This Server. Hello Client RPC."
    return nil
}

func starServer(port string){
    registServer(port)
    
    r := new(MyRPC)
    rpc.Register(r)

    tcpAddr, err := net.ResolveTCPAddr("tcp", port)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    fmt.Printf(" star server : %s \n", port)

    registServer(port)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        rpc.ServeConn(conn)
    }
}

func registServer(addr string){
    //注册zk节点q
	zkconn, zkerr := zk.GetConnect()
	if zkerr != nil {
		fmt.Printf(" connect zk error: %s ", zkerr)
	}	
	zkerr = zk.RegistServer(zkconn, addr)
	if zkerr != nil {
		fmt.Printf(" regist node error: %s ", zkerr)
    }
    // defer zkconn.Close()
}

func main() {   
   go starServer("127.0.0.1:8897")
   go starServer("127.0.0.1:8898")
   go starServer("127.0.0.1:8899")
   a := make(chan bool, 1)
   <-a

}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}