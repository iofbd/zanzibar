// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package module

import (
	barclientgenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/bar"
	barclientmodule "github.com/uber/zanzibar/examples/example-gateway/build/clients/bar/module"
	bazclientgenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz"
	bazclientmodule "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz/module"
	contactsclientgenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/contacts"
	contactsclientmodule "github.com/uber/zanzibar/examples/example-gateway/build/clients/contacts/module"
	googlenowclientgenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/google-now"
	googlenowclientmodule "github.com/uber/zanzibar/examples/example-gateway/build/clients/google-now/module"
	multiclientgenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/multi"
	multiclientmodule "github.com/uber/zanzibar/examples/example-gateway/build/clients/multi/module"
	quuxclientmodule "github.com/uber/zanzibar/examples/example-gateway/build/clients/quux/module"
	withexceptionsclientgenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/withexceptions"
	withexceptionsclientmodule "github.com/uber/zanzibar/examples/example-gateway/build/clients/withexceptions/module"
	barendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar"
	barendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
	bazendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz"
	bazendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	contactsendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/contacts"
	contactsendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/contacts/module"
	googlenowendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow"
	googlenowendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow/module"
	multiendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/multi"
	multiendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/multi/module"
	panicendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/panic"
	panicendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/panic/module"
	serverlessendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/serverless"
	serverlessendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/serverless/module"
	baztchannelendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/baz"
	baztchannelendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/baz/module"
	panictchannelendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/panic"
	panictchannelendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/panic/module"
	quuxendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/quux"
	quuxendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/quux/module"
	withexceptionsendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/withexceptions"
	withexceptionsendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/withexceptions/module"
	defaultexamplemiddlewaregenerated "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example"
	defaultexamplemiddlewaremodule "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example/module"
	defaultexample2middlewaregenerated "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example2"
	defaultexample2middlewaremodule "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example2/module"
	defaultexampletchannelmiddlewaregenerated "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example_tchannel"
	defaultexampletchannelmiddlewaremodule "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example_tchannel/module"
	examplemiddlewaregenerated "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/example"
	examplemiddlewaremodule "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/example/module"
	exampletchannelmiddlewaregenerated "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/example_tchannel"
	exampletchannelmiddlewaremodule "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/example_tchannel/module"
	quuxclientstatic "github.com/uber/zanzibar/examples/example-gateway/clients/quux"

	zanzibar "github.com/uber/zanzibar/runtime"
)

// DependenciesTree contains all deps for this service.
type DependenciesTree struct {
	Client     *ClientDependenciesNodes
	Middleware *MiddlewareDependenciesNodes
	Endpoint   *EndpointDependenciesNodes
}

// ClientDependenciesNodes contains client dependencies
type ClientDependenciesNodes struct {
	Bar            barclientgenerated.Client
	Baz            bazclientgenerated.Client
	Contacts       contactsclientgenerated.Client
	GoogleNow      googlenowclientgenerated.Client
	Multi          multiclientgenerated.Client
	Quux           quuxclientstatic.IClient
	Withexceptions withexceptionsclientgenerated.Client
}

// MiddlewareDependenciesNodes contains middleware dependencies
type MiddlewareDependenciesNodes struct {
	DefaultExample         defaultexamplemiddlewaregenerated.Middleware
	DefaultExample2        defaultexample2middlewaregenerated.Middleware
	DefaultExampleTchannel defaultexampletchannelmiddlewaregenerated.Middleware
	Example                examplemiddlewaregenerated.Middleware
	ExampleTchannel        exampletchannelmiddlewaregenerated.Middleware
}

// EndpointDependenciesNodes contains endpoint dependencies
type EndpointDependenciesNodes struct {
	Bar            barendpointgenerated.Endpoint
	Baz            bazendpointgenerated.Endpoint
	Contacts       contactsendpointgenerated.Endpoint
	Googlenow      googlenowendpointgenerated.Endpoint
	Multi          multiendpointgenerated.Endpoint
	Panic          panicendpointgenerated.Endpoint
	Serverless     serverlessendpointgenerated.Endpoint
	BazTChannel    baztchannelendpointgenerated.Endpoint
	PanicTChannel  panictchannelendpointgenerated.Endpoint
	Quux           quuxendpointgenerated.Endpoint
	Withexceptions withexceptionsendpointgenerated.Endpoint
}

// InitializeDependencies fully initializes all dependencies in the dep tree
// for the example-gateway service
func InitializeDependencies(
	g *zanzibar.Gateway,
) (*DependenciesTree, *Dependencies) {
	tree := &DependenciesTree{}

	initializedDefaultDependencies := &zanzibar.DefaultDependencies{
		Logger:           g.Logger,
		ContextExtractor: g.ContextExtractor,
		ContextLogger:    g.ContextLogger,
		ContextMetrics:   zanzibar.NewContextMetrics(g.RootScope),
		Scope:            g.RootScope,
		Tracer:           g.Tracer,
		Config:           g.Config,
		Channel:          g.Channel,

		GRPCClientDispatcher: g.GRPCClientDispatcher,
	}

	initializedClientDependencies := &ClientDependenciesNodes{}
	tree.Client = initializedClientDependencies
	initializedClientDependencies.Bar = barclientgenerated.NewClient(&barclientmodule.Dependencies{
		Default: initializedDefaultDependencies,
	})
	initializedClientDependencies.Baz = bazclientgenerated.NewClient(&bazclientmodule.Dependencies{
		Default: initializedDefaultDependencies,
	})
	initializedClientDependencies.Contacts = contactsclientgenerated.NewClient(&contactsclientmodule.Dependencies{
		Default: initializedDefaultDependencies,
	})
	initializedClientDependencies.GoogleNow = googlenowclientgenerated.NewClient(&googlenowclientmodule.Dependencies{
		Default: initializedDefaultDependencies,
	})
	initializedClientDependencies.Multi = multiclientgenerated.NewClient(&multiclientmodule.Dependencies{
		Default: initializedDefaultDependencies,
	})
	initializedClientDependencies.Quux = quuxclientstatic.NewClient(&quuxclientmodule.Dependencies{
		Default: initializedDefaultDependencies,
	})
	initializedClientDependencies.Withexceptions = withexceptionsclientgenerated.NewClient(&withexceptionsclientmodule.Dependencies{
		Default: initializedDefaultDependencies,
	})

	initializedMiddlewareDependencies := &MiddlewareDependenciesNodes{}
	tree.Middleware = initializedMiddlewareDependencies
	initializedMiddlewareDependencies.DefaultExample = defaultexamplemiddlewaregenerated.NewMiddleware(&defaultexamplemiddlewaremodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &defaultexamplemiddlewaremodule.ClientDependencies{
			Baz: initializedClientDependencies.Baz,
		},
	})
	initializedMiddlewareDependencies.DefaultExample2 = defaultexample2middlewaregenerated.NewMiddleware(&defaultexample2middlewaremodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &defaultexample2middlewaremodule.ClientDependencies{
			Baz: initializedClientDependencies.Baz,
		},
	})
	initializedMiddlewareDependencies.DefaultExampleTchannel = defaultexampletchannelmiddlewaregenerated.NewMiddleware(&defaultexampletchannelmiddlewaremodule.Dependencies{
		Default: initializedDefaultDependencies,
	})
	initializedMiddlewareDependencies.Example = examplemiddlewaregenerated.NewMiddleware(&examplemiddlewaremodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &examplemiddlewaremodule.ClientDependencies{
			Baz: initializedClientDependencies.Baz,
		},
	})
	initializedMiddlewareDependencies.ExampleTchannel = exampletchannelmiddlewaregenerated.NewMiddleware(&exampletchannelmiddlewaremodule.Dependencies{
		Default: initializedDefaultDependencies,
	})

	initializedEndpointDependencies := &EndpointDependenciesNodes{}
	tree.Endpoint = initializedEndpointDependencies
	initializedEndpointDependencies.Bar = barendpointgenerated.NewEndpoint(&barendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &barendpointmodule.ClientDependencies{
			Bar: initializedClientDependencies.Bar,
		},
		Middleware: &barendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
			Example:                initializedMiddlewareDependencies.Example,
		},
	})
	initializedEndpointDependencies.Baz = bazendpointgenerated.NewEndpoint(&bazendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &bazendpointmodule.ClientDependencies{
			Baz: initializedClientDependencies.Baz,
		},
		Middleware: &bazendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
		},
	})
	initializedEndpointDependencies.Contacts = contactsendpointgenerated.NewEndpoint(&contactsendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &contactsendpointmodule.ClientDependencies{
			Contacts: initializedClientDependencies.Contacts,
		},
		Middleware: &contactsendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
		},
	})
	initializedEndpointDependencies.Googlenow = googlenowendpointgenerated.NewEndpoint(&googlenowendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &googlenowendpointmodule.ClientDependencies{
			GoogleNow: initializedClientDependencies.GoogleNow,
		},
		Middleware: &googlenowendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
		},
	})
	initializedEndpointDependencies.Multi = multiendpointgenerated.NewEndpoint(&multiendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &multiendpointmodule.ClientDependencies{
			Multi: initializedClientDependencies.Multi,
		},
		Middleware: &multiendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
		},
	})
	initializedEndpointDependencies.Panic = panicendpointgenerated.NewEndpoint(&panicendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &panicendpointmodule.ClientDependencies{
			Multi: initializedClientDependencies.Multi,
		},
		Middleware: &panicendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
		},
	})
	initializedEndpointDependencies.Serverless = serverlessendpointgenerated.NewEndpoint(&serverlessendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Middleware: &serverlessendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
		},
	})
	initializedEndpointDependencies.BazTChannel = baztchannelendpointgenerated.NewEndpoint(&baztchannelendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &baztchannelendpointmodule.ClientDependencies{
			Baz:  initializedClientDependencies.Baz,
			Quux: initializedClientDependencies.Quux,
		},
		Middleware: &baztchannelendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
			ExampleTchannel:        initializedMiddlewareDependencies.ExampleTchannel,
		},
	})
	initializedEndpointDependencies.PanicTChannel = panictchannelendpointgenerated.NewEndpoint(&panictchannelendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &panictchannelendpointmodule.ClientDependencies{
			Baz: initializedClientDependencies.Baz,
		},
		Middleware: &panictchannelendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
		},
	})
	initializedEndpointDependencies.Quux = quuxendpointgenerated.NewEndpoint(&quuxendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &quuxendpointmodule.ClientDependencies{
			Quux: initializedClientDependencies.Quux,
		},
		Middleware: &quuxendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
		},
	})
	initializedEndpointDependencies.Withexceptions = withexceptionsendpointgenerated.NewEndpoint(&withexceptionsendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &withexceptionsendpointmodule.ClientDependencies{
			Withexceptions: initializedClientDependencies.Withexceptions,
		},
		Middleware: &withexceptionsendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
		},
	})

	dependencies := &Dependencies{
		Default: initializedDefaultDependencies,
		Endpoint: &EndpointDependencies{
			Bar:            initializedEndpointDependencies.Bar,
			Baz:            initializedEndpointDependencies.Baz,
			Contacts:       initializedEndpointDependencies.Contacts,
			Googlenow:      initializedEndpointDependencies.Googlenow,
			Multi:          initializedEndpointDependencies.Multi,
			Panic:          initializedEndpointDependencies.Panic,
			Serverless:     initializedEndpointDependencies.Serverless,
			BazTChannel:    initializedEndpointDependencies.BazTChannel,
			PanicTChannel:  initializedEndpointDependencies.PanicTChannel,
			Quux:           initializedEndpointDependencies.Quux,
			Withexceptions: initializedEndpointDependencies.Withexceptions,
		},
	}

	return tree, dependencies
}
