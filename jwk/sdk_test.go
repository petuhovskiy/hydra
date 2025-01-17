/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package jwk_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/petuhovskiy/hydra/sdk/go/hydra/client"
	"github.com/petuhovskiy/hydra/sdk/go/hydra/client/admin"
	"github.com/petuhovskiy/hydra/sdk/go/hydra/models"
	"github.com/ory/x/pointerx"
	"github.com/ory/x/urlx"

	"github.com/petuhovskiy/hydra/x"

	"github.com/petuhovskiy/hydra/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/petuhovskiy/hydra/jwk"
)

func TestJWKSDK(t *testing.T) {
	conf := internal.NewConfigurationWithDefaults()
	reg := internal.NewRegistry(conf)

	router := x.NewRouterAdmin()
	h := NewHandler(reg, conf)
	h.SetRoutes(router, x.NewRouterPublic(), func(h http.Handler) http.Handler {
		return h
	})
	server := httptest.NewServer(router)
	sdk := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{Schemes: []string{"http"}, Host: urlx.ParseOrPanic(server.URL).Host})

	t.Run("JSON Web Key", func(t *testing.T) {
		t.Run("CreateJwkSetKey", func(t *testing.T) {
			// Create a key called set-foo
			resultKeys, err := sdk.Admin.CreateJSONWebKeySet(admin.NewCreateJSONWebKeySetParams().WithSet("set-foo").WithBody(&models.CreateRequest{
				Algorithm: pointerx.String("HS256"),
				KeyID:     pointerx.String("key-bar"),
				Use:       pointerx.String("sig"),
			}))
			require.NoError(t, err)
			require.Len(t, resultKeys.Payload.Keys, 1)
			assert.Equal(t, "key-bar", *resultKeys.Payload.Keys[0].Kid)
			assert.Equal(t, "HS256", *resultKeys.Payload.Keys[0].Alg)
			assert.Equal(t, "sig", *resultKeys.Payload.Keys[0].Use)
		})

		var resultKeys *models.SwaggerJSONWebKeySet
		t.Run("GetJwkSetKey after create", func(t *testing.T) {
			result, err := sdk.Admin.GetJSONWebKey(admin.NewGetJSONWebKeyParams().WithKid("key-bar").WithSet("set-foo"))
			require.NoError(t, err)
			require.Len(t, result.Payload.Keys, 1)
			require.Equal(t, "key-bar", *result.Payload.Keys[0].Kid)
			require.Equal(t, "HS256", *result.Payload.Keys[0].Alg)

			resultKeys = result.Payload
		})

		t.Run("UpdateJwkSetKey", func(t *testing.T) {
			require.Len(t, resultKeys.Keys, 1)
			resultKeys.Keys[0].Alg = pointerx.String("RS256")

			resultKey, err := sdk.Admin.UpdateJSONWebKey(admin.NewUpdateJSONWebKeyParams().WithKid("key-bar").WithSet("set-foo").WithBody(resultKeys.Keys[0]))
			require.NoError(t, err)
			assert.Equal(t, "key-bar", *resultKey.Payload.Kid)
			assert.Equal(t, "RS256", *resultKey.Payload.Alg)
		})

		t.Run("DeleteJwkSetKey after delete", func(t *testing.T) {
			_, err := sdk.Admin.DeleteJSONWebKey(admin.NewDeleteJSONWebKeyParams().WithKid("key-bar").WithSet("set-foo"))
			require.NoError(t, err)
		})

		t.Run("GetJwkSetKey after delete", func(t *testing.T) {
			_, err := sdk.Admin.GetJSONWebKey(admin.NewGetJSONWebKeyParams().WithKid("key-bar").WithSet("set-foo"))
			require.Error(t, err)
		})

	})

	t.Run("JWK Set", func(t *testing.T) {
		t.Run("CreateJwkSetKey", func(t *testing.T) {
			resultKeys, err := sdk.Admin.CreateJSONWebKeySet(admin.NewCreateJSONWebKeySetParams().WithSet("set-foo2").WithBody(&models.CreateRequest{
				Algorithm: pointerx.String("HS256"),
				KeyID:     pointerx.String("key-bar"),
			}))
			require.NoError(t, err)

			require.Len(t, resultKeys.Payload.Keys, 1)
			assert.Equal(t, "key-bar", *resultKeys.Payload.Keys[0].Kid)
			assert.Equal(t, "HS256", *resultKeys.Payload.Keys[0].Alg)
		})

		resultKeys, err := sdk.Admin.GetJSONWebKeySet(admin.NewGetJSONWebKeySetParams().WithSet("set-foo2"))
		t.Run("GetJwkSet after create", func(t *testing.T) {
			require.NoError(t, err)
			require.Len(t, resultKeys.Payload.Keys, 1)
			assert.Equal(t, "key-bar", *resultKeys.Payload.Keys[0].Kid)
			assert.Equal(t, "HS256", *resultKeys.Payload.Keys[0].Alg)
		})

		t.Run("UpdateJwkSet", func(t *testing.T) {
			require.Len(t, resultKeys.Payload.Keys, 1)
			resultKeys.Payload.Keys[0].Alg = pointerx.String("RS256")

			result, err := sdk.Admin.UpdateJSONWebKeySet(admin.NewUpdateJSONWebKeySetParams().WithSet("set-foo2").WithBody(resultKeys.Payload))
			require.NoError(t, err)
			require.Len(t, result.Payload.Keys, 1)
			assert.Equal(t, "key-bar", *result.Payload.Keys[0].Kid)
			assert.Equal(t, "RS256", *result.Payload.Keys[0].Alg)
		})

		t.Run("DeleteJwkSet", func(t *testing.T) {
			_, err := sdk.Admin.DeleteJSONWebKeySet(admin.NewDeleteJSONWebKeySetParams().WithSet("set-foo2"))
			require.NoError(t, err)
		})

		t.Run("GetJwkSet after delete", func(t *testing.T) {
			_, err := sdk.Admin.GetJSONWebKeySet(admin.NewGetJSONWebKeySetParams().WithSet("set-foo2"))
			require.Error(t, err)
		})

		t.Run("GetJwkSetKey after delete", func(t *testing.T) {
			_, err := sdk.Admin.GetJSONWebKey(admin.NewGetJSONWebKeyParams().WithSet("set-foo2").WithKid("key-bar"))
			require.Error(t, err)
		})
	})
}
