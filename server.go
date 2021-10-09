/*
   Copyright 2021 Mark Wilkerson

   This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// populated at build time from ldflags
var VERSION string

func main() {

	opts := Init()

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", opts.BindAddr, opts.Port),
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
		Handler:      handler(opts),
	}

	log.Printf("vanity v%s listening on %s:%d", VERSION, opts.BindAddr, opts.Port)

	if opts.SSLCert != "" {

		log.Fatal(server.ListenAndServeTLS(opts.SSLCert, opts.SSLKeyFile))

	} else {

		log.Fatal(server.ListenAndServe())

	}

}
