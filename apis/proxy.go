// Package interfaces for common interfaces
// MIT License
//
// Copyright (c) 2017 yroffin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package apis

import (
	"log"
	"net/http"

	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	core_models "github.com/yroffin/go-boot-sqllite/core/models"
	proxy_models "github.com/yroffin/go-google-gateway/models"
)

// Proxy internal members
type Proxy struct {
	// Base component
	*core_apis.API
	// internal members
	Name string
	// mounts
	crud    string `path:"/api/proxies"`
	execute string `path:"/api/execute" handler:"ExecuteProxy" method:"GET" mime-type:""`
	ssl     string `path:"/api/version" handler:"GetVersion" method:"GET" mime-type:""`
}

// IProxy implements IBean
type IProxy interface {
	core_bean.IBean
}

// PostConstruct this API
func (p *Proxy) Init() error {
	// Crud
	p.HandlerGetAll = func() (string, error) {
		return p.GenericGetAll(&proxy_models.ProxyBean{}, core_models.IPersistents(&proxy_models.ProxyBeans{Collection: make([]core_models.IPersistent, 0)}))
	}
	p.HandlerGetByID = func(id string) (string, error) {
		return p.GenericGetByID(id, &proxy_models.ProxyBean{})
	}
	p.HandlerPost = func(body string) (string, error) {
		return p.GenericPost(body, &proxy_models.ProxyBean{})
	}
	p.HandlerPutByID = func(id string, body string) (string, error) {
		return p.GenericPutByID(id, body, &proxy_models.ProxyBean{})
	}
	p.HandlerDeleteByID = func(id string) (string, error) {
		return p.GenericDeleteByID(id, &proxy_models.ProxyBean{})
	}
	p.HandlerPatchByID = func(id string, body string) (string, error) {
		return p.GenericPatchByID(id, body, &proxy_models.ProxyBean{})
	}
	return p.API.Init()
}

// PostConstruct this API
func (p *Proxy) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p)
	return nil
}

// Validate this API
func (p *Proxy) Validate(name string) error {
	return nil
}

// ExecuteProxy render ExecuteProxy
func (p *Proxy) ExecuteProxy() func(w http.ResponseWriter, r *http.Request) {
	anonymous := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request Url %v", r.URL)
		log.Printf("Request Headers %v", r.Header)
		log.Printf("Request Encoding %v", r.TransferEncoding)
		log.Printf("Request IP %v", r.RemoteAddr)
		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(200)
		w.Write([]byte("Message de test"))
	}
	return anonymous
}

// GetVersion get gateway version
func (p *Proxy) GetVersion() func(w http.ResponseWriter, r *http.Request) {
	anonymous := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte("{\"version\":\"v1.0\"}"))
	}
	return anonymous
}
