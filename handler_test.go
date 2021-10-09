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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tt := []struct {
		label      string
		input      string
		method     string
		want       string
		statuscode int
		debug      bool
	}{
		{
			label:      "Valid request",
			input:      "https://whatver.biz:8080/foo?go-get=1",
			method:     http.MethodGet,
			want:       `<meta name="go-import" content="whatver.biz/foo git test.com/foo.git">`,
			statuscode: http.StatusOK,
			debug:      true,
		},
		{
			label:      "Valid request, nested path",
			input:      "https://whatver.biz:8080/foo/bar/baz?go-get=1",
			method:     http.MethodGet,
			want:       `<meta name="go-import" content="whatver.biz/foo/bar/baz git test.com/foo/bar/baz.git">`,
			statuscode: http.StatusOK,
			debug:      true,
		},
		{
			label:      "Valid URI, bad method: POST",
			input:      "https://whatver.biz:8080/foo?go-get=1",
			method:     http.MethodPost,
			want:       "",
			statuscode: http.StatusMethodNotAllowed,
		},
		{
			label:      "Valid URI, bad method: PUT",
			input:      "https://whatver.biz:8080/foo?go-get=1",
			method:     http.MethodPut,
			want:       "",
			statuscode: http.StatusMethodNotAllowed,
		},
		{
			label:      "Valid URI, bad method: HEAD",
			input:      "https://whatver.biz:8080/foo?go-get=1",
			method:     http.MethodHead,
			want:       "",
			statuscode: http.StatusMethodNotAllowed,
		},
		{
			label:      "Valid URI, bad method: DELETE",
			input:      "https://whatver.biz:8080/foo?go-get=1",
			method:     http.MethodDelete,
			want:       "",
			statuscode: http.StatusMethodNotAllowed,
		},
		{
			label:      "Wrong go-get value",
			input:      "https://whatver.biz:8080/foo?go-get=",
			method:     http.MethodGet,
			want:       "",
			statuscode: http.StatusNotFound,
		},
		{
			label:      "Missing go-get",
			input:      "https://whatver.biz:8080/foo?go-away=1",
			method:     http.MethodGet,
			want:       "",
			statuscode: http.StatusNotFound,
		},
		{
			label:      "No query params",
			input:      "https://whatver.biz:8080/foo",
			method:     http.MethodGet,
			want:       "",
			statuscode: http.StatusNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.label, func(t *testing.T) {
			opts := options{
				BindAddr: "0.0.0.0",
				Port:     8080,
				Dest:     "test.com",
				Debug:    tc.debug,
			}
			h := handler(opts)
			req, err := http.NewRequest(tc.method, tc.input, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			mock := httptest.NewRecorder()

			h(mock, req)

			result := mock.Result()
			if result.StatusCode != tc.statuscode {
				t.Fatalf("expected status %d, got %d", tc.statuscode, result.StatusCode)
			}

			body, err := ioutil.ReadAll(result.Body)
			result.Body.Close()
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}

			if string(body) != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, body)
			}

		})
	}
}
