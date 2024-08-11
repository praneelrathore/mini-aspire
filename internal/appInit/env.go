package appInit

import (
	"context"
)

// Env is the application environment exposing application scoped instances
// in request handlers
type Env struct {
	appDB *Connection
}

type envContextKey string

const (
	// EnvCtxKey is the key to set and retrieve Env in context
	EnvCtxKey envContextKey = "env"
)

// NewEnv returns a new Env instance
func NewEnv(options ...func(env *Env)) *Env {
	env := &Env{}

	for _, option := range options {
		option(env)
	}

	return env
}

func (env *Env) AddEnv(options ...func(env *Env)) {
	for _, option := range options {
		option(env)
	}
}

// WithDatabaseConnection sets a database connection in the Env
func WithDatabaseConnection(db *Connection) func(*Env) {
	return func(env *Env) {
		env.appDB = db
	}
}

// DB retrieves the database connection from the Env
func (env *Env) DB() *Connection {
	return env.appDB
}

// WithContext returns a context containing the env Value
func (env *Env) WithContext(ctx context.Context) context.Context {
	nctx := context.WithValue(ctx, EnvCtxKey, env)
	return nctx
}
