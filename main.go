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
package main

import (
	// fmt has methods for formatted IO

	// the "net/http" library has methods for HTTP

	// Gorilla router

	// Apis
	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
	core_business "github.com/yroffin/go-boot-sqllite/core/business"
	core_manager "github.com/yroffin/go-boot-sqllite/core/manager"
	proxy_apis "github.com/yroffin/go-google-gateway/apis"
)

// Rest()
func main() {
	// declare manager and boot it
	var m = core_manager.Manager{}
	m.Init()
	// Core beans
	m.Register("router", &core_apis.Router{Bean: &core_bean.Bean{}})
	m.Register("crud-business", &core_business.CrudBusiness{Bean: &core_bean.Bean{}})
	// API beans
	m.Register("proxy", &proxy_apis.Proxy{API: &core_apis.API{Bean: &core_bean.Bean{}}})
	m.Boot()
	// Declarre listener
	m.HTTP(8080)
	m.HTTPS(8443)
	m.Wait()
}
