// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && ignore
// +build darwin,amd64,ignore

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"go-darwin.dev/exp/libproc"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("usage: %s [pid]", os.Args[0])
	}
	s := os.Args[1]

	pid, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	path, err := libproc.ProcPidpath(libproc.PID(pid))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("libproc.ProcPidpath: pid: %d, process name: %s\n", pid, path)
}
