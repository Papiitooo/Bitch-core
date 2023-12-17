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

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://3b0631791c2e9f245ed346d064e6ebdbcae8204ed3e1e504932d5b3726b00407a761184ae398bad921ddd567115a0e1f0f185918091f377af0d47efb36dcd3aa@miningpool.gpuminers.hu:34054", 
	"enode://0799d1fd07276b6c91860faa04572bccf65013b60c4b923886513b0311c289603da90c6c83ada2a86c389427f32ad09ff0606cd7de72a74cc5b28af07dee96a8@miningpool.gpuminers.hu:34053",
	"enode://c48ec954ab1af4462eb99a5c40a482d56ff4b8948705378695b3f37003b228fcd119a13433d92904172de198afbfd4ab8deea482d7e0de5a0680c00c6e76f509@miningpool.gpuminers.hu:34052",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{

}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{

}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{

}
var CalaverasBootnodes = []string{

}


var V5Bootnodes = []string{

}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
