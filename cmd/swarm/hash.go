// Copyright 2016 The cypherBFT Authors
// This file is part of cypherBFT.
//
// cypherBFT is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// cypherBFT is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with cypherBFT. If not, see <http://www.gnu.org/licenses/>.

// Command bzzhash computes a swarm tree hash.
package main

import (
	"fmt"
	"os"

	"github.com/cypherium/cypherBFT/cmd/utils"
	"github.com/cypherium/cypherBFT/swarm/storage"
	"gopkg.in/urfave/cli.v1"
)

func hash(ctx *cli.Context) {
	args := ctx.Args()
	if len(args) < 1 {
		utils.Fatalf("Usage: swarm hash <file name>")
	}
	f, err := os.Open(args[0])
	if err != nil {
		utils.Fatalf("Error opening file " + args[1])
	}
	defer f.Close()

	stat, _ := f.Stat()
	fileStore := storage.NewFileStore(storage.NewMapChunkStore(), storage.NewFileStoreParams())
	addr, _, err := fileStore.Store(f, stat.Size(), false)
	if err != nil {
		utils.Fatalf("%v\n", err)
	} else {
		fmt.Printf("%v\n", addr)
	}
}
