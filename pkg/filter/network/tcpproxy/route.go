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

package tcpproxy

import (
	"net"

	v2 "mosn.io/mosn/pkg/config/v2"
	"mosn.io/mosn/pkg/log"
)

type route struct {
	clusterName      string
	sourceAddrs      IpRangeList
	destinationAddrs IpRangeList
	sourcePort       PortRangeList
	destinationPort  PortRangeList
}

func newRoute(routeConfig *v2.Route) *route {
	res := &route{
		clusterName:      routeConfig.Cluster,
		sourceAddrs:      IpRangeList{routeConfig.SourceAddrs},
		destinationAddrs: IpRangeList{routeConfig.DestinationAddrs},
		sourcePort:       ParsePortRangeList(routeConfig.SourcePort),
		destinationPort:  ParsePortRangeList(routeConfig.DestinationPort),
	}

	return res
}

type PortRangeList struct {
	portList []PortRange
}

func (pr *PortRangeList) Contains(address net.Addr) bool {
	_, port, ok := GetIpAndPort(address)
	if ok {
		log.DefaultLogger.Tracef("PortRangeList check port = %v , address = %v", port, address)
		for _, portRange := range pr.portList {
			log.DefaultLogger.Tracef("check port range , port range = %v , port = %v", portRange, port)
			if port >= portRange.min && port <= portRange.max {
				return true
			}
		}
	}
	return false
}

type PortRange struct {
	min int
	max int
}

type IpRangeList struct {
	cidrRanges []v2.CidrRange
}

func (ipList *IpRangeList) Contains(address net.Addr) bool {
	ip, _, ok := GetIpAndPort(address)
	log.DefaultLogger.Tracef("IpRangeList check ip = %v,address = %v", ip, address)
	if ok {
		for _, cidrRange := range ipList.cidrRanges {
			log.DefaultLogger.Tracef("check CidrRange = %v,ip = %v", cidrRange, ip)
			if cidrRange.IsInRange(ip) {
				return true
			}
		}
	}
	return false
}

func GetIpAndPort(address net.Addr) (net.IP, int, bool) {
	tcpAddr, ok := address.(*net.TCPAddr)
	if ok {
		return tcpAddr.IP, tcpAddr.Port, true
	}

	if udpAddr, ok := address.(*net.UDPAddr); ok {
		return udpAddr.IP, udpAddr.Port, true
	}

	return nil, 0, false
}
