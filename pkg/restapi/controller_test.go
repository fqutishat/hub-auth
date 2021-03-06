/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package restapi

import (
	"crypto/aes"
	"crypto/rand"
	"testing"

	"github.com/trustbloc/hub-auth/pkg/internal/common/mockoidc"

	"github.com/stretchr/testify/require"
	"github.com/trustbloc/edge-core/pkg/storage/memstore"

	"github.com/trustbloc/hub-auth/pkg/restapi/operation"
)

func TestController_New(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		config := config(t)

		controller, err := New(config)
		require.NoError(t, err)
		require.NotNil(t, controller)
	})

	t.Run("error if operations cannot start", func(t *testing.T) {
		config := config(t)
		config.OIDCProviderURL = "BadURL"

		_, err := New(config)
		require.Error(t, err)
	})
}

func TestController_GetOperations(t *testing.T) {
	config := config(t)

	controller, err := New(config)
	require.NoError(t, err)
	require.NotNil(t, controller)

	ops := controller.GetOperations()
	require.NotEmpty(t, ops)
}

func config(t *testing.T) *operation.Config {
	path := mockoidc.StartProvider(t)

	return &operation.Config{
		OIDCProviderURL:        path,
		TransientStoreProvider: memstore.NewProvider(),
		StoreProvider:          memstore.NewProvider(),
		Cookies: &operation.CookieConfig{
			AuthKey: cookieKey(t),
			EncKey:  cookieKey(t),
		},
		StartupTimeout: 1,
	}
}

func cookieKey(t *testing.T) []byte {
	key := make([]byte, aes.BlockSize)

	_, err := rand.Read(key)
	require.NoError(t, err)

	return key
}
