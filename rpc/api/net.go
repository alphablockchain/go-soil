// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package api

import (
	"github.com/soilcurrency/go-soil/eth"
	"github.com/soilcurrency/go-soil/rpc/codec"
	"github.com/soilcurrency/go-soil/rpc/shared"
	"github.com/soilcurrency/go-soil/xeth"
)

const (
	NetApiVersion = "1.0"
)

var (
	// mapping between methods and handlers
	netMapping = map[string]nethandler{
		"net_peerCount": (*netApi).PeerCount,
		"net_listening": (*netApi).IsListening,
		"net_version":   (*netApi).Version,
	}
)

// net callback handler
type nethandler func(*netApi, *shared.Request) (interface{}, error)

// net api provider
type netApi struct {
	xeth     *xeth.XEth
	ethereum *eth.Ethereum
	methods  map[string]nethandler
	codec    codec.ApiCoder
}

// create a new net api instance
func NewNetApi(xeth *xeth.XEth, eth *eth.Ethereum, coder codec.Codec) *netApi {
	return &netApi{
		xeth:     xeth,
		ethereum: eth,
		methods:  netMapping,
		codec:    coder.New(nil),
	}
}

// collection with supported methods
func (self *netApi) Methods() []string {
	methods := make([]string, len(self.methods))
	i := 0
	for k := range self.methods {
		methods[i] = k
		i++
	}
	return methods
}

// Execute given request
func (self *netApi) Execute(req *shared.Request) (interface{}, error) {
	if callback, ok := self.methods[req.Method]; ok {
		return callback(self, req)
	}

	return nil, shared.NewNotImplementedError(req.Method)
}

func (self *netApi) Name() string {
	return shared.NetApiName
}

func (self *netApi) ApiVersion() string {
	return NetApiVersion
}

// Number of connected peers
func (self *netApi) PeerCount(req *shared.Request) (interface{}, error) {
	return newHexNum(self.xeth.PeerCount()), nil
}

func (self *netApi) IsListening(req *shared.Request) (interface{}, error) {
	return self.xeth.IsListening(), nil
}

func (self *netApi) Version(req *shared.Request) (interface{}, error) {
	return self.xeth.NetworkVersion(), nil
}
