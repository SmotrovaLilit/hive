package configuration

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/ory/x/tracing"
)

type HasherArgon2Config struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

type SelfServiceHook struct {
	Run    string          `json:"run"`
	Config json.RawMessage `json:"config"`
}

type SelfServiceStrategy struct {
	Enabled bool            `json:"enabled"`
	Config  json.RawMessage `json:"config"`
}

type Provider interface {
	AdminListenOn() string
	PublicListenOn() string
	DSN() string

	SessionSecrets() [][]byte

	SelfPublicURL() *url.URL
	SelfAdminURL() *url.URL

	DefaultReturnToURL() *url.URL

	LoginURL() *url.URL
	ErrorURL() *url.URL
	MultiFactorURL() *url.URL

	SessionLifespan() time.Duration
	SelfServiceLoginRequestLifespan() time.Duration
	SelfServiceRegistrationRequestLifespan() time.Duration

	SelfServiceStrategy(strategy string) *SelfServiceStrategy
	SelfServiceLoginBeforeHooks() []SelfServiceHook
	SelfServiceRegistrationBeforeHooks() []SelfServiceHook
	SelfServiceLoginAfterHooks(strategy string) []SelfServiceHook
	SelfServiceRegistrationAfterHooks(strategy string) []SelfServiceHook
	SelfServiceLogoutRedirectURL() *url.URL

	DefaultIdentityTraitsSchemaURL() *url.URL

	WhitelistedReturnToDomains() []url.URL

	RegisterURL() *url.URL

	HashersArgon2() *HasherArgon2Config

	TracingServiceName() string
	TracingProvider() string
	TracingJaegerConfig() *tracing.JaegerConfig
}
