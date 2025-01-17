package internal

import (
	"context"

	"github.com/petuhovskiy/hydra/client"
	"github.com/petuhovskiy/hydra/driver"
)

func AddFositeExamples(r driver.Registry) {
	for _, c := range []client.Client{
		{
			ClientID:      "my-client",
			Secret:        "foobar",
			RedirectURIs:  []string{"http://localhost:3846/callback"},
			ResponseTypes: []string{"id_token", "code", "token"},
			GrantTypes:    []string{"implicit", "refresh_token", "authorization_code", "password", "client_credentials"},
			Scope:         "fosite,openid,photos,offline",
		},
		{
			ClientID:      "encoded:client",
			Secret:        "encoded&password",
			RedirectURIs:  []string{"http://localhost:3846/callback"},
			ResponseTypes: []string{"id_token", "code", "token"},
			GrantTypes:    []string{"implicit", "refresh_token", "authorization_code", "password", "client_credentials"},
			Scope:         "fosite,openid,photos,offline",
		},
	} {
		if err := r.ClientManager().CreateClient(context.Background(), &c); err != nil {
			panic(err)
		}
	}
}
