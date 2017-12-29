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
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	core_apis "github.com/yroffin/go-boot-sqllite/core/apis"
	core_bean "github.com/yroffin/go-boot-sqllite/core/bean"
)

// Proxy internal members
type Proxy struct {
	// Base component
	*core_apis.API
	// internal members
	Name string
	// Mqtt
	client mqtt.Client
	// mounts
	post string `path:"/api/assistant" handler:"ExecuteProxy" method:"POST" mime-type:""`
}

// IProxy implements IBean
type IProxy interface {
	core_bean.IBean
}

// PostConstruct this API
func (p *Proxy) Init() error {
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
	// options
	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.1.111:1883")
	opts.SetClientID("GoogleHome")

	// client
	p.client = mqtt.NewClient(opts)
	if token := p.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	var wg sync.WaitGroup
	wg.Add(1)

	const TOPIC = "testtopic/test"
	if token := p.client.Subscribe(TOPIC, 0, func(client mqtt.Client, msg mqtt.Message) {
		if string(msg.Payload()) != "mymessage" {
			log.Fatalf("want mymessage, got %s", msg.Payload())
		}
		wg.Done()
	}); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	if token := p.client.Publish(TOPIC, 0, false, "mymessage"); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	wg.Wait()
	log.Printf("MQTT is ok")
	return nil
}

// SendMessage send MQTT message
func (p *Proxy) SendMessage(topic string, message string) {
	if token := p.client.Publish(topic, 0, false, message); token.Wait() && token.Error() != nil {
		log.Print(token.Error())
	}
}

// ExecuteProxy render ExecuteProxy
func (p *Proxy) ExecuteProxy() func(w http.ResponseWriter, r *http.Request) {
	anonymous := func(w http.ResponseWriter, r *http.Request) {
		var headers = r.Header["X-Api-Google"]
		if len(headers) > 0 && headers[0] == "YES" {
			log.Printf("Request IP %v with good header", r.RemoteAddr)
			body, _ := ioutil.ReadAll(r.Body)
			var strBody = string(body)
			if strings.Contains(strBody, "yroffin-dialogflow") {
				p.SendMessage("/assistant", strBody)
			}
			w.Header().Set("Content-type", "text/plain")
			w.WriteHeader(200)
			w.Write(body)
		} else {
			log.Printf("Request IP %v", r.RemoteAddr)
			w.WriteHeader(400)
		}
	}
	return anonymous
}
