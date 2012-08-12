//
// httphere
// Copyright 2012 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: Robert Sesek (rsesek@google.com)
//

package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

var (
	port = flag.String("port", "0", "Port to run the server. 0 for a random port.")
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		print(err)
		os.Exit(1)
	}

	server := http.FileServer(http.Dir(pwd))

	addr := net.JoinHostPort("", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		print(err)
		os.Exit(2)
	}

	fmt.Printf("Listening on %s\n", listener.Addr().String())
	fmt.Printf("Serving conents of %s\n", pwd)
	http.Serve(listener, server)
}
