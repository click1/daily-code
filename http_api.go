package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var http_data string

func JsonServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	io.WriteString(w, strings.Replace(http_data, "\n", "", -1))
}

func main() {

	filepath := flag.String("file", "", "输入加载文件完整路径")
	api := flag.String("api", "", "输入产生api的接口地址")

	flag.Parse()

	if (*filepath == "") || (*api == "") {
		fmt.Println("示例: ./http_api -file=/tmp/test.json -api=/test.json")
	}

	data, err := ioutil.ReadFile(*filepath)
	if err != nil {
		fmt.Println("读取文件失败", *filepath)
		os.Exit(1)
	}

	http_data = string(data)
	http.HandleFunc(*api, JsonServer)

	fmt.Println("http请求创建")
	fmt.Println("访问:http://127.0.0.1:9898" + *api)

	err = http.ListenAndServe(":9898", nil)

	if err != nil {
		fmt.Println("监听http端口失败")
		os.Exit(2)
	}

}
