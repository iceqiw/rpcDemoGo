package main

import (
    "fmt"
    "net/rpc"
    "./zk"
    "log"
    "errors"
    "math/rand"
	"time"
)

var S string

func startClient(){

    serverHost, err := getServerHost()
    fmt.Printf("get server host : %s \n", serverHost)
	if err != nil {
		fmt.Printf("get server host fail: %s \n", err)
		return
    }
    
    client, err := rpc.Dial("tcp", serverHost)
    if err != nil {
        log.Fatal("dialhttp: ", err)
    }

    var reply *string
    S = "This Client. Hello Server RPC."
    err = client.Call("MyRPC.HelloRPC", S, &reply)
    if err != nil {
        log.Fatal("call hellorpc: ", err)
    }

    fmt.Println(*reply)
}

func getServerHost() (host string, err error) {
	conn, err := zk.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s \n ", err)
		return
	}
	defer conn.Close()
	serverList, err := zk.GetServerList(conn)
	if err != nil {
		fmt.Printf(" get server list error: %s \n", err)
		return
	}

	count := len(serverList)
	if count == 0 {
		err = errors.New("server list is empty \n")
		return
	}

	//随机选中一个返回
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = serverList[r.Intn(3)]
	return
}

func main() {
    startClient()
}