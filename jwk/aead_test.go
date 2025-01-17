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
	"crypto/rand"
	"fmt"
	"io"
	"testing"

	"github.com/ory/viper"

	"github.com/petuhovskiy/hydra/driver/configuration"
	"github.com/petuhovskiy/hydra/internal"
	. "github.com/petuhovskiy/hydra/jwk"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func secret(t *testing.T) string {
	bytes := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, bytes)
	require.NoError(t, err)
	return fmt.Sprintf("%X", bytes)
}

func TestAEAD(t *testing.T) {
	c := internal.NewConfigurationWithDefaults()
	t.Run("case=without-rotation", func(t *testing.T) {
		viper.Set(configuration.ViperKeyGetSystemSecret, []string{secret(t)})
		a := NewAEAD(c)

		plain := []byte(uuid.New())
		ct, err := a.Encrypt(plain)
		assert.NoError(t, err)

		res, err := a.Decrypt(ct)
		assert.NoError(t, err)
		assert.Equal(t, plain, res)
	})

	t.Run("case=wrong-secret", func(t *testing.T) {
		viper.Set(configuration.ViperKeyGetSystemSecret, []string{secret(t)})
		a := NewAEAD(c)

		ct, err := a.Encrypt([]byte(uuid.New()))
		require.NoError(t, err)

		viper.Set(configuration.ViperKeyGetSystemSecret, []string{secret(t)})
		_, err = a.Decrypt(ct)
		require.Error(t, err)
	})

	t.Run("case=with-rotation", func(t *testing.T) {
		old := secret(t)
		viper.Set(configuration.ViperKeyGetSystemSecret, []string{old})
		a := NewAEAD(c)

		plain := []byte(uuid.New())
		ct, err := a.Encrypt(plain)
		require.NoError(t, err)

		// Sets the old secret as a rotated secret and creates a new one.
		viper.Set(configuration.ViperKeyGetSystemSecret, []string{secret(t), old})
		res, err := a.Decrypt(ct)
		require.NoError(t, err)
		assert.Equal(t, plain, res)

		// THis should also work when we re-encrypt the same plain text.
		ct2, err := a.Encrypt(plain)
		require.NoError(t, err)
		assert.NotEqual(t, ct2, ct)

		res, err = a.Decrypt(ct)
		require.NoError(t, err)
		assert.Equal(t, plain, res)
	})

	t.Run("case=with-rotation-wrong-secret", func(t *testing.T) {
		viper.Set(configuration.ViperKeyGetSystemSecret, []string{secret(t)})
		a := NewAEAD(c)

		plain := []byte(uuid.New())
		ct, err := a.Encrypt(plain)
		require.NoError(t, err)

		// When the secrets do not match, an error should be thrown during decryption.
		viper.Set(configuration.ViperKeyGetSystemSecret, []string{secret(t), secret(t)})
		_, err = a.Decrypt(ct)
		require.Error(t, err)
	})
}
