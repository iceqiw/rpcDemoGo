package zk

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

const (
	timeOut = 20
)

var hosts []string = []string{"192.168.99.100:32773"} // the zk server list

func GetConnect() (conn *zk.Conn, err error) {
	conn, _, err = zk.Connect(hosts, timeOut*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func RegistServer(conn *zk.Conn, host string) (err error) {
	_, err = conn.Create("/rpcservers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/rpcservers")
	return
}
