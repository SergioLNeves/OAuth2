package config

import "time"

const (
	Development = "development"
	Staging     = "staging"
	Production  = "production"
)

type Config struct {
	Env       string `env:"ENV,default=development"`
	Server    Server
	OAuth     OAuth
	Cors      Cors
	SQLite    SQLite
	RateLimit RateLimit
	Time      Time
	AuthDB    AuthDB
	Database  Database
	URL       URL
	Session   Session
}

type Server struct {
	Port            int           `env:"SERVER_PORT,default=8080"`
	ShutdownTimeout time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT,default=10s"`
}

type Cors struct {
	AllowedOrigins []string `env:"CORS_ALLOWED_ORIGINS,default=*"`
	AllowedMethods []string `env:"CORS_ALLOWED_METHODS,default=GET|POST|PUT|DELETE|OPTIONS"`
	AllowedHeaders []string `env:"CORS_ALLOWED_HEADERS,default=Content-Type,Authorization"`
}

// OAuth contains OAuth 2.0 and OIDC configuration
type OAuth struct {
	AuthCodeExpiration     time.Duration `env:"OAUTH_AUTH_CODE_EXPIRATION,default=10m"`
	AccessTokenExpiration  time.Duration `env:"OAUTH_ACCESS_TOKEN_EXPIRATION,default=1h"`
	RefreshTokenExpiration time.Duration `env:"OAUTH_REFRESH_TOKEN_EXPIRATION,default=180h"` // 7 days
	IDTokenExpiration      time.Duration `env:"OAUTH_ID_TOKEN_EXPIRATION,default=1h"`
	Issuer                 string        `env:"OAUTH_ISSUER,default=http://localhost:5001"`
}

// SQLite contains SQLite database connection pool configuration.
type SQLite struct {
	MaxConn     int           `env:"DATABASE_MAX_CONN,default=10"`
	MaxIdle     int           `env:"DATABASE_MAX_IDLE,default=5"`
	MaxLifeTime time.Duration `env:"DATABASE_MAX_LIFE_TIME,default=300s"`
}

// RateLimit contains rate limiting configuration.
type RateLimit struct {
	MaxRequests int           `env:"RATE_LIMIT_REQUESTS,default=100"`
	Window      time.Duration `env:"RATE_LIMIT_WINDOW,default=1m"`
}

// Time contains time configuration.
type Time struct {
	Timezone string `env:"TIME_TIMEZONE,default=UTC"`
}

// AuthFB contains authentication database configuration.
type AuthDB struct {
	// DBPath is the path to the authentication SQLite database
	DBPath string `env:"AUTH_DB_PATH,default=pa_data/auth.db"`
}

// Database contains database operation timeout configuration.
type Database struct {
	TimeoutRead   time.Duration `env:"DATABASE_TIMEOUT_READ,default=2s"`
	TimeoutWrite  time.Duration `env:"DATABASE_TIMEOUT_WRITE,default=3s"`
	TimeoutUpdate time.Duration `env:"DATABASE_TIMEOUT_UPDATE,default=3s"`
	TimeoutDelete time.Duration `env:"DATABASE_TIMEOUT_DELETE,default=2s"`
	TimeoutCount  time.Duration `env:"DATABASE_TIMEOUT_COUNT,default=5s"`
}

// URL contains application URLs configuration.
type URL struct {
	APIBaseURL string `env:"API_BASE_URL,default=http://localhost:5001"`
	APPBaseURL string `env:"APP_BASE_URL,default=http://localhost:5173"`
}

// Session contains session management configuration.
type Session struct {
	Secret    string        `env:"SESSION_SECRET,default=V4o6aXjqGN2LUkNDmimHp4JLlx7+Rzf4g/WL74Icdj4="`
	Duration  time.Duration `env:"SESSION_DURATION,default=168h"`
	TokenSize int           `env:"SESSION_TOKEN_SIZE,default=32"`
}

// IsDevelopment returns true if the environment is development.
func (e *Config) IsDevelopment() bool {
	return e.Env == Development
}

// IsStaging returns true if the environment is staging.
func (e *Config) IsStaging() bool {
	return e.Env == Staging
}

// IsProduction returns true if the environment is production.
func (e *Config) IsProduction() bool {
	return e.Env == Production
}
