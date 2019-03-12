package tcp

import (
	"ScreenHttpServer/models"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func Tcp_main() *models.ScreenData {
	var (
		host   = "127.0.0.1"
		port   = "16385"
		remote = host + ":" + port
		msg    = "@$_$len=00000024{\"CID\":\"90000\",\"PL\":\"cq\"}#$_$\r\n"
		data   = make([]byte, 4096)
		count  = 0
	)
	ScreenDat := &models.ScreenData{}
	conn, _ := net.Dial("tcp", remote)
	io.WriteString(conn, msg)
	for {
		count, _ = conn.Read(data)
		datstr := string(data[:count])
		fmt.Println(datstr)
		err := json.Unmarshal(data[:count], &ScreenDat)
		if err != nil {
			fmt.Println(err)
		}
		break
	}
	conn.Close()
	return ScreenDat
}
