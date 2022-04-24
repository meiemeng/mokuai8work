package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	c, python, java := true, false, "no!"
	//定义healthz用于返回200
	http.HandleFunc("/healthz", healthz)
	fmt.Println(c, python, java)
	err := http.ListenAndServe(":80", nil)
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", rootHandler)
	// mux.HandleFunc("/healthz", healthz)
	// mux.HandleFunc("/debug/pprof/", pprof.Index)
	// mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	//问题4当访问 localhost/healthz 时，应返回 200
	//调用healthz函数，并返回200返回值
	io.WriteString(w, "200\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	//问题1：接收客户端 request，并将 request 中带的 header 写入 response header
	//这里已经实现了将request的r.Header写入进 responsewriter中了
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	//问题2: 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	//问题解析1.环境变量中是否存在VERSION环境变量需要做判断如果存在直接get不存在要走set这一步
	// LookupEnv为检索key这个健对应的环境变量的值 如果环境变量存在，则返回对应的值，并且布尔值为true
	//反之则为false 返回值为空
	_, ok := os.LookupEnv("VERSION")
	if !ok {
		os.Setenv("VERSION", "golandv1.18.1")
	}
	version := os.Getenv("VERSION")
	//写入respose Header
	io.WriteString(w, version)

	//问题3：Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	//这里调用另外的函数来获取到 IP  HTTP返回码
	addr, statuscode := GetipStatus(w, r)
	//打印输出addr status并记录进日志
	fmt.Printf("addr%s : status is %d", addr, statuscode)
	glog.V(2).Infof("addr is %s: status is %d", addr, statuscode)

}

func GetipStatus(w http.ResponseWriter, r *http.Request) (addr string, statuscode int) {
	//为了防止中间存在代理模式所以要定义RemoteHeader,用于获取客户端真实IP
	const remoteAddrHeader = "X-AppEngine-Remote-Addr"
	if addr = r.Header.Get(remoteAddrHeader); addr != "" {
		//获取客户端ip,并赋予变量
		addr = r.RemoteAddr
		//删除header中代理相关配置，防止攻击
		r.Header.Del(remoteAddrHeader)
	} else {
		r.RemoteAddr = "127.0.0.1"
		addr = r.RemoteAddr
	}
	//fg := r.Response.StatusCode
	//fmt.Printf("%T", fg)
	//fmt.Println(fg)
	//w.WriteHeader()
	statuscode = http.StatusOK
	w.WriteHeader(statuscode)
	return addr, statuscode

}
