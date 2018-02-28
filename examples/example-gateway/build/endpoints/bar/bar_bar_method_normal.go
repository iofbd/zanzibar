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

package barendpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	"github.com/uber/zanzibar/examples/example-gateway/middlewares/example"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// BarNormalHandler is the handler for "/bar/bar-path"
type BarNormalHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewBarNormalHandler creates a handler
func NewBarNormalHandler(deps *module.Dependencies) *BarNormalHandler {
	handler := &BarNormalHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"bar", "normal",
		zanzibar.NewStack([]zanzibar.MiddlewareHandle{
			deps.Middleware.Example.NewMiddlewareHandle(
				example.Options{
					Baz: []string{"foo", "bar"},
					Foo: "test",
				},
			),
		}, handler.HandleRequest).Handle,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *BarNormalHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"POST", "/bar/bar-path",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/bar/bar-path".
func (h *BarNormalHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	var requestBody endpointsBarBar.Bar_Normal_Args
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	workflow := BarNormalEndpoint{
		Clients: h.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBarBar.BarException:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		default:
			res.SendError(500, "Unexpected server error", err)
			return
		}

	}
	// TODO(jakev): implement writing fields into response headers

	res.WriteJSON(200, cliRespHeaders, response)
}

// BarNormalEndpoint calls thrift client Bar.Normal
type BarNormalEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w BarNormalEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_Normal_Args,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {
	clientRequest := convertToNormalClientRequest(r)

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Bar.Normal(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBarBar.BarException:
			serverErr := convertNormalBarException(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.Error(errValue),
				zap.String("client", "Bar"),
			)

			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertBarNormalClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToNormalClientRequest(in *endpointsBarBar.Bar_Normal_Args) *clientsBarBar.Bar_Normal_Args {
	out := &clientsBarBar.Bar_Normal_Args{}

	if in.Request != nil {
		out.Request = &clientsBarBar.BarRequest{}
		out.Request.StringField = string(in.Request.StringField)
		out.Request.BoolField = bool(in.Request.BoolField)
		out.Request.BinaryField = []byte(in.Request.BinaryField)
		out.Request.Timestamp = clientsBarBar.Timestamp(in.Request.Timestamp)
		out.Request.EnumField = clientsBarBar.Fruit(in.Request.EnumField)
		out.Request.LongField = clientsBarBar.Long(in.Request.LongField)
	} else {
		out.Request = nil
	}

	return out
}

func convertNormalBarException(
	clientError *clientsBarBar.BarException,
) *endpointsBarBar.BarException {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBarBar.BarException{}
	return serverError
}

func convertBarNormalClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	out := &endpointsBarBar.BarResponse{}

	out.StringField = string(in.StringField)
	out.IntWithRange = int32(in.IntWithRange)
	out.IntWithoutRange = int32(in.IntWithoutRange)
	out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
	for key1, value2 := range in.MapIntWithRange {
		out.MapIntWithRange[endpointsBarBar.UUID(key1)] = int32(value2)
	}
	out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
	for key3, value4 := range in.MapIntWithoutRange {
		out.MapIntWithoutRange[key3] = int32(value4)
	}
	out.BinaryField = []byte(in.BinaryField)
	var convertBarResponseHelper5 func(in *clientsBarBar.BarResponse) (out *endpointsBarBar.BarResponse)
	convertBarResponseHelper5 = func(in *clientsBarBar.BarResponse) (out *endpointsBarBar.BarResponse) {
		if in != nil {
			out = &endpointsBarBar.BarResponse{}
			out.StringField = string(in.StringField)
			out.IntWithRange = int32(in.IntWithRange)
			out.IntWithoutRange = int32(in.IntWithoutRange)
			out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
			for key6, value7 := range in.MapIntWithRange {
				out.MapIntWithRange[endpointsBarBar.UUID(key6)] = int32(value7)
			}
			out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
			for key8, value9 := range in.MapIntWithoutRange {
				out.MapIntWithoutRange[key8] = int32(value9)
			}
			out.BinaryField = []byte(in.BinaryField)
			out.NextResponse = convertBarResponseHelper5(in.NextResponse)
		} else {
			out = nil
		}
		return
	}
	out.NextResponse = convertBarResponseHelper5(in.NextResponse)

	return out
}
