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
	"encoding/json"
	"fmt"

	"mosn.io/api"

	v2 "mosn.io/mosn/pkg/config/v2"
)

func init() {
	api.RegisterNetwork(v2.UDP_PROXY, CreateUDPProxyFactory)
}
// TODO(move to somewhere)
type udpProxyFilterConfigFactory struct {
	proxy *v2.UDPProxy
}

func (f *udpProxyFilterConfigFactory) CreateFilterChain(context context.Context, callbacks api.NetWorkFilterChainFactoryCallbacks) {
	rf := NewUDPProxy(context, f.proxy)
	callbacks.AddReadFilter(rf)
}

func CreateUDPProxyFactory(conf map[string]interface{}) (api.NetworkFilterChainFactory, error) {
	p, err := ParseUDPProxy(conf)
	if err != nil {
		return nil, err
	}
	return &udpProxyFilterConfigFactory{
		proxy: p,
	}, nil
}

// ParseUDPProxy convert the map to v2.UDPProxy
func ParseUDPProxy(cfg map[string]interface{}) (*v2.UDPProxy, error) {
	proxy := &v2.UDPProxy{}
	if data, err := json.Marshal(cfg); err == nil {
		json.Unmarshal(data, proxy)
	} else {
		return nil, fmt.Errorf("[config] config is not a udp proxy config: %v", err)
	}
	return proxy, nil
}
