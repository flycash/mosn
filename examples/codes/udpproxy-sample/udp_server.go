/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"net"
	"strings"
)

// 读取消息
func handleConnection(udpConn *net.UDPConn) {

	buf := make([]byte, 1024)
	len, udpAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil{
		return
	}
	logContent := strings.Replace(string(buf),"\n","",1)
	fmt.Println("server read len:", len)
	fmt.Println("server read data:", logContent)

	// 发送数据
	len, err = udpConn.WriteToUDP([]byte("ok\r\n"), udpAddr)
	if err != nil{
		return
	}

	fmt.Println("server write len:", len, "\n")
}

func main() {
	udpAddr, _ := net.ResolveUDPAddr("udp4", "127.0.0.1:9998")

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer udpConn.Close()

	fmt.Println("udp listening ... ")

	for {
		handleConnection(udpConn)
	}
}
