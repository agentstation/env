// Package env provides a structure around the environment setting that the web application is running in.
package env

import (
	"log"
	"strings"
	"sync"
)

// Environment represents the application environment
type Environment string

var (
	current Environment  // used to optionally track the current env in this package
	mu      sync.RWMutex // mutex to protect access to current
)

const (
	// PROD is the production environment
	PROD Environment = "PROD"

	// STAGE is the staging environment
	STAGE Environment = "STAGE"

	// DEV is the development environment
	DEV Environment = "DEV"
)

// String returns a human-readable string representation of the Environment
func (e Environment) String() string {
	switch e {
	case PROD:
		return "Production"
	case STAGE:
		return "Staging"
	case DEV:
		return "Development"
	default:
		return "Unknown"
	}
}

// IsDev returns true if the Environment is DEV
func (e Environment) IsDev() bool { return e == DEV }

// IsStage returns true if the Environment is STAGE
func (e Environment) IsStage() bool { return e == STAGE }

// IsProd returns true if the Environment is PROD
func (e Environment) IsProd() bool { return e == PROD }

// Parse converts a string to an Environment type
func Parse(env string) Environment {

	// Trim spaces and remove quotes
	env = strings.TrimSpace(env)
	env = strings.Trim(env, `"'`)

	// Convert to uppercase and match known environments
	switch strings.ToUpper(env) {
	case "PROD", "PRODUCTION":
		return PROD
	case "STAGE", "STAGING":
		return STAGE
	case "DEV", "DEVELOP", "DEVELOPMENT":
		return DEV
	default:
		log.Printf("Warning: invalid environment '%s', defaulting to PROD", env)
		return PROD
	}
}

// Set sets the current environment
func Set(env Environment) {
	mu.Lock()
	defer mu.Unlock()
	current = env
}

// Current returns the current environment
func Current() Environment {
	mu.RLock()
	defer mu.RUnlock()
	if current == "" {
		return PROD // Default to PROD if not set
	}
	return current
}
