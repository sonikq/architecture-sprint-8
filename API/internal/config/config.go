package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	RunAddress string
	CtxTimeout time.Duration

	Keycloak Keycloak
}

type Keycloak struct {
	URI         string
	Realm       string
	Client      string
	AllowedRole string
}

const (
	defaultMode           = "debug"
	defaultRunAddress     = "localhost:8000"
	defaultCtxTimeOut     = 1000
	defaultKeycloakURI    = "http://keycloak:8080"
	defaultKeycloakRealm  = "reports-realm"
	defaultKeycloakClient = "reports-api"
	defaultAllowedRole    = "prothetic_user"
)

func Load() (Config, error) {
	startupMode := flag.String("mode", defaultMode, "start mode(debug or release)")
	runAddress := flag.String("run_address", defaultRunAddress, "run address defines on what port and host the server will be started")
	ctxTimeout := flag.Int("context_timeout", defaultCtxTimeOut, "value for define context timeout")
	keycloakURI := flag.String("keycloak_uri", defaultKeycloakURI, "defines address where placed keycloak service")
	keycloakRealm := flag.String("keycloak_realm", defaultKeycloakRealm, "defines keycloak realm")
	keycloakClient := flag.String("keycloak_client", defaultKeycloakClient, "defines name of the keycloak client")
	keycloakAllowedRole := flag.String("keycloak_allowed_role", defaultAllowedRole, "defines name who has rights to access")
	flag.Parse()

	switch *startupMode {
	case "debug":
		if err := godotenv.Load("internal/config/.env"); err != nil {
			return Config{}, err
		}
	case "release":

	default:
		log.Fatal("invalid mode: " + *startupMode)
	}

	cfg := new(Config)

	cfg.RunAddress = getEnvString(*runAddress, "RUN_ADDRESS")
	cfg.CtxTimeout = time.Millisecond * time.Duration(getEnvInt(*ctxTimeout, "CTX_TIMEOUT"))

	cfg.Keycloak.URI = getEnvString(*keycloakURI, "KEYCLOAK_URI")
	cfg.Keycloak.Realm = getEnvString(*keycloakRealm, "KEYCLOAK_REALM")
	cfg.Keycloak.Client = getEnvString(*keycloakClient, "KEYCLOAK_CLIENT")
	cfg.Keycloak.AllowedRole = getEnvString(*keycloakAllowedRole, "KEYCLOAK_ALLOWED_ROLE")

	return *cfg, nil
}

// getEnvString - function for determining the priority between flags and environment variables in string format
func getEnvString(flagValue string, envKey string) string {
	envValue, exists := os.LookupEnv(envKey)
	if exists {
		return envValue
	}
	return flagValue
}

// getEnvInt - function for determining the priority between flags and environment variables in int format
func getEnvInt(flagValue int, envKey string) int {
	envValue, exists := os.LookupEnv(envKey)
	if exists {
		intVal, err := strconv.Atoi(envValue)
		if err != nil {
			log.Printf("cant convert env-key: %s to int", envValue)
			return 1
		}

		return intVal
	}

	return flagValue
}
