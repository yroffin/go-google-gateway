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
	"flag"
	"log"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/yroffin/go-boot-sqllite/core/engine"
	"github.com/yroffin/go-boot-sqllite/core/winter"
)

func init() {
	winter.Helper.Register("ProxyBean", (&Proxy{}).New())
}

// Proxy internal members
type Proxy struct {
	// Base component
	*engine.API
	// internal members
	Name string
	// Mqtt
	client mqtt.Client
	// Swagger with injection mecanism
	Swagger engine.ISwaggerService `@autowired:"swagger"`
	// Api
	post interface{} `@handler:"ExecuteProxy" path:"/api/assistant" method:"POST" mime-type:"/application/json"`
	// Intent header control
	intentHeader *string
	// Value of header control
	intentHeaderValue *string
}

// IProxy implements IBean
type IProxy interface {
	engine.IAPI
}

// New constructor
func (p *Proxy) New() IProxy {
	bean := Proxy{API: &engine.API{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *Proxy) Init() error {
	return p.API.Init()
}

// PostConstruct this API
func (p *Proxy) PostConstruct(name string) error {
	// Scan struct and init all handler
	p.ScanHandler(p.Swagger, p)
	// scan flags
	p.intentHeader = flag.String("intent-header", "DEFINE-IT", "Intent header control")
	p.intentHeaderValue = flag.String("intent-header-value", "DEFINE-IT", "Intent header control value")
	flag.Parse()
	log.Printf("Header control %v must be to %v", *p.intentHeader, *p.intentHeaderValue)
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
func (p *Proxy) ExecuteProxy() func(engine.IHttpContext) {
	anonymous := func(c engine.IHttpContext) {
		c.Header("Content-type", "application/json")
		var data, _ = c.GetRawData()
		p.SendMessage("/assistant", string(data))
		c.String(200, string(data))
	}
	return anonymous
}
