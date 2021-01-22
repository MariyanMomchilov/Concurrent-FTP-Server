package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
)

type cmdRequest interface {
	apply(net.Conn)
}

type getFileRequest struct {
	filename string
}

type listRequest struct {
	flags string
}

func (r *getFileRequest) apply(conn net.Conn) {
	file, err := os.Open(r.filename)

	if err != nil { // error is of type path error only
		// consider a change about how messages are passed between Server and Client
		io.WriteString(conn, fmt.Sprintf("error %s", err))
		return
	}
	io.Copy(conn, file)
}

func (r *listRequest) apply(conn net.Conn) {

	// to do
	if r.flags != "" {
		return
	}

	filesInfo, err := ioutil.ReadDir(".")
	if err != nil {
		io.WriteString(conn, fmt.Sprintf("error %s", err))
	}

	for i := range filesInfo {
		//fmt.Println(filesInfo[i].Name())
		if i == len(filesInfo)-1 {
			io.WriteString(conn, filesInfo[i].Name()+"\n")
			break
		}

		io.WriteString(conn, filesInfo[i].Name()+" ")
	}
}
