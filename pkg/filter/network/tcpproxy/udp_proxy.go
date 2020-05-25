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
	"context"

	v2 "mosn.io/mosn/pkg/config/v2"
	"mosn.io/mosn/pkg/log"
	"mosn.io/mosn/pkg/types"
)

type udpProxy struct {
	proxy
}

func NewUDPProxy(ctx context.Context, config *v2.UDPProxy) Proxy {
	return &udpProxy{proxy: *newProxy(ctx, newUdpProxyConfig(config))}
}

func (p *udpProxy) initConnForCluster(ctx types.LoadBalancerContext, clusterSnapshot types.ClusterSnapshot) types.CreateConnectionData {
	return p.clusterManager.UDPConnForCluster(ctx, clusterSnapshot)
}

func newUdpProxyConfig(config *v2.UDPProxy) ProxyConfig {

	var routes []*route

	log.DefaultLogger.Tracef("Udp Proxy :: New Proxy Config = %v", config)
	for _, routeConfig := range config.Routes {
		route := newRoute((*v2.Route)(routeConfig))
		log.DefaultLogger.Tracef("Udp Proxy add one route : %v", route)
		routes = append(routes, route)
	}

	return &udpProxyConfig{
		proxyConfig: proxyConfig{
			statPrefix: config.StatPrefix,
			cluster:    config.Cluster,
			routes:     routes,
		},
	}
}

type udpProxyConfig struct {
	proxyConfig
}
