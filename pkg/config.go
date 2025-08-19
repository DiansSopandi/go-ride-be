package pkg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host        string `mapstructure:"db_host"`
	Port        int    `mapstructure:"db_port"`
	User        string `mapstructure:"db_user"`
	Password    string `mapstructure:"db_password"`
	DBName      string `mapstructure:"db_name"`
	SSLMode     string `mapstructure:"sslmode"`
	Timezone    string `mapstructure:"timezone"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"redis_host"`
	Port     int    `mapstructure:"redis_port"`
	Password string `mapstructure:"redis_password"`
	DB       int    `mapstructure:"redis_db"`
}

type ApplicationConfig struct {
	Name                       string `mapstructure:"name"`
	Version                    string `mapstructure:"version"`
	Env                        string `mapstructure:"env"`
	Host                       string `mapstructure:"host"`
	AppPort                    int    `mapstructure:"app_port"`
	AppUrl                     string `mapstructure:"app_url"`
	AppPath                    string `mapstructure:"app_path"`
	RedirectPath               string `mapstructure:"redirect_path"`
	WsUrl                      string `mapstructure:"ws_url"`
	Timezone                   string `mapstructure:"timezone"`
	EnableLog                  bool   `mapstructure:"enable_log"`
	EnableLogToFile            bool   `mapstructure:"enable_log_to_file"`
	LogPath                    string `mapstructure:"log_path"`
	Prefork                    bool   `mapstructure:"prefork"`
	AllowOrigins               string `mapstructure:"allow_origins"`
	AllowHeaders               string `mapstructure:"allow_headers"`
	AllowMethods               string `mapstructure:"allow_methods"`
	EnableTrustedProxyCheck    bool   `mapstructure:"enable_trusted_proxy_check"`
	EnableCache                bool   `mapstructure:"enable_cache"`
	AppKey                     string `mapstructure:"app_key"`
	JwtSecretKey               string `mapstructure:"jwt_secret_key"`
	SsoJwtSecret               string `mapstructure:"sso_jwt_secret"`
	FilePath                   string `mapstructure:"file_path"`
	DefaultMaxRequestPerMinute int    `mapstructure:"default_max_requests_per_minute"`
	// DefaultRequestDuration     time.Duration `mapstructure:"default_request_duration"`
}

type Config struct {
	Database    DatabaseConfig    `mapstructure:"database"`
	Redis       RedisConfig       `mapstructure:"redis"`
	Application ApplicationConfig `mapstructure:"application"`
	// SsoClientCredentials SsoClientCredentials  `mapstructure:"sso_client_credentials"`
	// Email                EmailConfig           `mapstructure:"email"`
	// Instrumentation      InstrumentationConfig `mapstructure:"instrumentation"`
	// RingBufferQueue      RingBufferQueue       `mapstructure:"ring_buffer_queue"`
	// Installers           InstallerConfig       `mapstructure:"installers"`
	// Manticore            ManticoreConfig       `mapstructure:"manticore"`
	// Minio                MinioConfig           `mapstructure:"minio"`
}

// type Config struct {
//     Application struct {
//         AppPort           int    `mapstructure:"app_port"`
//         AppPath           string `mapstructure:"app_path"`
//         AppURL            string `mapstructure:"app_url"`
//         EnableDBMigration bool   `mapstructure:"enable_db_migration"`
//         JWTSecretKey      string `mapstructure:"jwt_secret_key"`
//         EnableLog         bool   `mapstructure:"enable_log"`
//         EnableLogToFile   bool   `mapstructure:"enable_log_to_file"`
//         LogPath           string `mapstructure:"log_path"`
//     } `mapstructure:"application"`

//     Database struct {
//         DBHost     string `mapstructure:"db_host"`
//         DBPort     int    `mapstructure:"db_port"`
//         DBUser     string `mapstructure:"db_user"`
//         DBPassword string `mapstructure:"db_password"`
//         DBName     string `mapstructure:"db_name"`
//     } `mapstructure:"database"`
// }

var Cfg Config

// LoadConfig loads configuration from file.
func LoadConfig(configFileName string, path string) {
	viper.SetConfigName(configFileName)
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

}

// GetEnv gets environment variable value by key
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Warning: Environment variable %s is not set", key)
	}
	return value
}

// GetEnvOrDefault gets environment variable value by key or returns default value if not set
func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Environment variable %s is not set, using default value: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
