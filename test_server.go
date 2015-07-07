package main

import (
	"mail-test/server"
	"mail-test/serverutil"
	"fmt"
	"crypto/tls"
	"net"
)
const password = "password"
type InsecureAuth struct{}
func (i InsecureAuth) Authenticate (u, p string) (server.AllowAddrFunc,error) {
	return func (a string) (bool) {
		return a == u && password == p
	},nil
}
func main () {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		fmt.Println(err)
		return
	}
	config := tls.Config{Certificates:[]tls.Certificate{cert},InsecureSkipVerify:true}
	
	listen,err := net.Listen("tcp","localhost:2500")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	s := server.NewServer("localhost:2500")
	s.TLSconfig = &config
	s.Auth = InsecureAuth{}
	serverutil.Serve(s,listen)
}