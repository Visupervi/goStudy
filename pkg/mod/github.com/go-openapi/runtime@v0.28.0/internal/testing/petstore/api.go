// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package petstore

import (
	goerrors "errors"
	"io"
	"net/http"
	"strings"
	gotest "testing"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	testingutil "github.com/go-openapi/runtime/internal/testing"
	"github.com/go-openapi/runtime/middleware/untyped"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/runtime/yamlpc"
	"github.com/stretchr/testify/require"
)

const (
	apiPrincipal = "admin"
	apiUser      = "topuser"
	otherUser    = "anyother"
)

// NewAPI registers a stub api for the pet store
func NewAPI(t gotest.TB) (*loads.Document, *untyped.API) {
	spec, err := loads.Analyzed(testingutil.PetStoreJSONMessage, "")
	require.NoError(t, err)
	api := untyped.NewAPI(spec)

	api.RegisterConsumer("application/json", runtime.JSONConsumer())
	api.RegisterProducer("application/json", runtime.JSONProducer())
	api.RegisterConsumer("application/xml", new(stubConsumer))
	api.RegisterProducer("application/xml", new(stubProducer))
	api.RegisterProducer("text/plain", new(stubProducer))
	api.RegisterProducer("text/html", new(stubProducer))
	api.RegisterConsumer("application/x-yaml", yamlpc.YAMLConsumer())
	api.RegisterProducer("application/x-yaml", yamlpc.YAMLProducer())

	api.RegisterAuth("basic", security.BasicAuth(func(username, password string) (interface{}, error) {
		switch {
		case username == apiPrincipal && password == apiPrincipal:
			return apiPrincipal, nil
		case username == apiUser && password == apiUser:
			return apiUser, nil
		case username == otherUser && password == otherUser:
			return otherUser, nil
		default:
			return nil, errors.Unauthenticated("basic")
		}
	}))
	api.RegisterAuth("apiKey", security.APIKeyAuth("X-API-KEY", "header", func(token string) (interface{}, error) {
		if token == "token123" {
			return apiPrincipal, nil
		}
		return nil, errors.Unauthenticated("token")
	}))
	api.RegisterAuthorizer(runtime.AuthorizerFunc(func(r *http.Request, user interface{}) error {
		if r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/api/pets") && user.(string) != apiPrincipal {
			if user.(string) == apiUser {
				return errors.CompositeValidationError(errors.New(errors.InvalidTypeCode, "unauthorized"))
			}

			return goerrors.New("unauthorized")
		}
		return nil
	}))
	api.RegisterOperation("get", "/pets", new(stubOperationHandler))
	api.RegisterOperation("post", "/pets", new(stubOperationHandler))
	api.RegisterOperation("delete", "/pets/{id}", new(stubOperationHandler))
	api.RegisterOperation("get", "/pets/{id}", new(stubOperationHandler))

	api.Models["pet"] = func() interface{} { return new(Pet) }
	api.Models["newPet"] = func() interface{} { return new(Pet) }
	api.Models["tag"] = func() interface{} { return new(Tag) }

	return spec, api
}

// NewRootAPI registers a stub api for the pet store
func NewRootAPI(t gotest.TB) (*loads.Document, *untyped.API) {
	spec, err := loads.Analyzed(testingutil.RootPetStoreJSONMessage, "")
	require.NoError(t, err)
	api := untyped.NewAPI(spec)

	api.RegisterConsumer("application/json", runtime.JSONConsumer())
	api.RegisterProducer("application/json", runtime.JSONProducer())
	api.RegisterConsumer("application/xml", new(stubConsumer))
	api.RegisterProducer("application/xml", new(stubProducer))
	api.RegisterProducer("text/plain", new(stubProducer))
	api.RegisterProducer("text/html", new(stubProducer))
	api.RegisterConsumer("application/x-yaml", yamlpc.YAMLConsumer())
	api.RegisterProducer("application/x-yaml", yamlpc.YAMLProducer())

	api.RegisterAuth("basic", security.BasicAuth(func(username, password string) (interface{}, error) {
		if username == apiPrincipal && password == apiPrincipal {
			return apiPrincipal, nil
		}
		return nil, errors.Unauthenticated("basic")
	}))
	api.RegisterAuth("apiKey", security.APIKeyAuth("X-API-KEY", "header", func(token string) (interface{}, error) {
		if token == "token123" {
			return apiPrincipal, nil
		}
		return nil, errors.Unauthenticated("token")
	}))
	api.RegisterAuthorizer(security.Authorized())
	api.RegisterOperation("get", "/pets", new(stubOperationHandler))
	api.RegisterOperation("post", "/pets", new(stubOperationHandler))
	api.RegisterOperation("delete", "/pets/{id}", new(stubOperationHandler))
	api.RegisterOperation("get", "/pets/{id}", new(stubOperationHandler))

	api.Models["pet"] = func() interface{} { return new(Pet) }
	api.Models["newPet"] = func() interface{} { return new(Pet) }
	api.Models["tag"] = func() interface{} { return new(Tag) }

	return spec, api
}

// Tag the tag model
type Tag struct {
	ID   int64
	Name string
}

// Pet the pet model
type Pet struct {
	ID        int64
	Name      string
	PhotoURLs []string
	Status    string
	Tags      []Tag
}

type stubConsumer struct {
}

func (s *stubConsumer) Consume(_ io.Reader, _ interface{}) error {
	return nil
}

type stubProducer struct {
}

func (s *stubProducer) Produce(_ io.Writer, _ interface{}) error {
	return nil
}

type stubOperationHandler struct {
}

func (s *stubOperationHandler) ParameterModel() interface{} {
	return nil
}

func (s *stubOperationHandler) Handle(_ interface{}) (interface{}, error) {
	return map[string]interface{}{}, nil
}
