// Copyright 2015 The go-ethereum Authors
// Copyright 2017 The cypherBFT Authors
// This file is part of the cypherBFT library.
//
// The cypherBFT library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cypherBFT library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cypherBFT library. If not, see <http://www.gnu.org/licenses/>.

// Contains all the wrappers from the accounts package to support client side cnode
// management on mobile platforms.

package cypher

import (
	"errors"

	"github.com/cypherium/cypherBFT/p2p/discv5"
)

// Enode represents a host on the network.
type Enode struct {
	node *discv5.Node
}

// NewEnode parses a node designator.
//
// There are two basic forms of node designators
//   - incomplete nodes, which only have the public key (node ID)
//   - complete nodes, which contain the public key and IP/Port information
//
// For incomplete nodes, the designator must look like one of these
//
//    cnode://<hex node id>
//    <hex node id>
//
// For complete nodes, the node ID is encoded in the username portion
// of the URL, separated from the host by an @ sign. The hostname can
// only be given as an IP address, DNS domain names are not allowed.
// The port in the host name section is the TCP listening port. If the
// TCP and UDP (discovery) ports differ, the UDP port is specified as
// query parameter "discport".
//
// In the following example, the node URL describes
// a node with IP address 10.3.58.6, TCP listening port 30303
// and UDP discovery port 30301.
//
//    cnode://<hex node id>@10.3.58.6:30303?discport=30301
func NewEnode(rawurl string) (cnode *Enode, _ error) {
	node, err := discv5.ParseNode(rawurl)
	if err != nil {
		return nil, err
	}
	return &Enode{node}, nil
}

// Enodes represents a slice of accounts.
type Enodes struct{ nodes []*discv5.Node }

// NewEnodes creates a slice of uninitialized enodes.
func NewEnodes(size int) *Enodes {
	return &Enodes{
		nodes: make([]*discv5.Node, size),
	}
}

// NewEnodesEmpty creates an empty slice of Enode values.
func NewEnodesEmpty() *Enodes {
	return NewEnodes(0)
}

// Size returns the number of enodes in the slice.
func (e *Enodes) Size() int {
	return len(e.nodes)
}

// Get returns the cnode at the given index from the slice.
func (e *Enodes) Get(index int) (cnode *Enode, _ error) {
	if index < 0 || index >= len(e.nodes) {
		return nil, errors.New("index out of bounds")
	}
	return &Enode{e.nodes[index]}, nil
}

// Set sets the cnode at the given index in the slice.
func (e *Enodes) Set(index int, cnode *Enode) error {
	if index < 0 || index >= len(e.nodes) {
		return errors.New("index out of bounds")
	}
	e.nodes[index] = cnode.node
	return nil
}

// Append adds a new cnode element to the end of the slice.
func (e *Enodes) Append(cnode *Enode) {
	e.nodes = append(e.nodes, cnode.node)
}
