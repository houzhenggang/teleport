/*
Copyright 2015 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package web

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/url"

	"github.com/gravitational/teleport/lib/httplib"

	"github.com/gravitational/roundtrip"
	"github.com/gravitational/trace"
)

func newInsecureClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func newClientWithPool(pool *x509.CertPool) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: pool},
		},
	}
}

func newWebClient(url string, opts ...roundtrip.ClientParam) (*webClient, error) {
	clt, err := roundtrip.NewClient(url, APIVersion, opts...)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &webClient{clt}, nil
}

// webClient is a package local lightweight client used
// in tests and some functions to handle errors properly
type webClient struct {
	*roundtrip.Client
}

func (w *webClient) PostJSON(
	endpoint string, val interface{}) (*roundtrip.Response, error) {
	return httplib.ConvertResponse(w.Client.PostJSON(endpoint, val))
}

func (w *webClient) PutJSON(
	endpoint string, val interface{}) (*roundtrip.Response, error) {
	return httplib.ConvertResponse(w.Client.PutJSON(endpoint, val))
}

func (w *webClient) Get(endpoint string, val url.Values) (*roundtrip.Response, error) {
	return httplib.ConvertResponse(w.Client.Get(endpoint, val))
}

func (w *webClient) Delete(endpoint string) (*roundtrip.Response, error) {
	return httplib.ConvertResponse(w.Client.Delete(endpoint))
}
