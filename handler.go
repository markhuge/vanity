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
	"strings"
)

const logfmt = "%s	%s	%s	%d	%s"

func handler(opts options) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// This should only ever be an http GET
		// Anything else is shenanigans
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			log.Printf(logfmt, r.RemoteAddr, r.Method, r.URL, http.StatusMethodNotAllowed, r.UserAgent())
			return
		}

		qp := r.URL.Query()

		if qp.Get("go-get") == "1" {
			log.Printf(logfmt, r.RemoteAddr, r.Method, r.URL, http.StatusOK, r.UserAgent())

			host := strings.Split(r.Host, ":")[0]
			repo := r.URL.Path

			fmt.Fprintf(w, `<meta name="go-import" content="%s%s git %s%s.git">`, host, repo, opts.Dest, repo)

			if opts.Debug {
				log.Printf(`Sent: <meta name="go-import" content="%s%s git %s%s.git">`, host, repo, opts.Dest, repo)
			}

		} else {

			log.Printf(logfmt, r.RemoteAddr, r.Method, r.URL, http.StatusNotFound, r.UserAgent())
			w.WriteHeader(http.StatusNotFound)

		}
	}
}
