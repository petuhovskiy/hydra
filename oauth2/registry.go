package oauth2

import (
	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/openid"
	"github.com/petuhovskiy/hydra/client"
	"github.com/petuhovskiy/hydra/consent"
	"github.com/petuhovskiy/hydra/driver/configuration"
	"github.com/petuhovskiy/hydra/jwk"
	"github.com/petuhovskiy/hydra/x"
)

type InternalRegistry interface {
	client.Registry
	x.RegistryWriter
	x.RegistryLogger
	consent.Registry
	Registry
}

type Registry interface {
	OAuth2Storage() x.FositeStorer
	OAuth2Provider() fosite.OAuth2Provider
	AudienceStrategy() fosite.AudienceMatchingStrategy
	ScopeStrategy() fosite.ScopeStrategy

	AccessTokenJWTStrategy() jwk.JWTStrategy
	OpenIDJWTStrategy() jwk.JWTStrategy

	OpenIDConnectRequestValidator() *openid.OpenIDConnectRequestValidator
}

type Configuration interface {
	configuration.Provider
}
