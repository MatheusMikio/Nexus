package config


var(
	db *gorm.DB
	cfg Config
)

type Config struct{
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	DBSSLMode          string
	JWTSecret          string
	PasswordPepper     string
	ServerPort         string
	CORSAllowedOrigins []string
}

func Init() error {
	_ = godotenv.Load(filepath.Join("..", ".env"))

	var err error
	cfg, err = loadConfig()
	if err != nil{
		return err
	}

	db, err := initPostgreSQL(cfg)
	
	if err != nil{
		return err
	}

	return nil
}

func loadConfig() (Config, error) {
	config := Config{
		DBHost:         strings.TrimSpace(os.Getenv("DB_HOST")),
		DBPort:         strings.TrimSpace(os.Getenv("DB_PORT")),
		DBUser:         strings.TrimSpace(os.Getenv("DB_USER")),
		DBPassword:     strings.TrimSpace(os.Getenv("DB_PASSWORD")),
		DBName:         strings.TrimSpace(os.Getenv("DB_NAME")),
		DBSSLMode:      strings.TrimSpace(os.Getenv("DB_SSLMODE")),
		JWTSecret:      strings.TrimSpace(os.Getenv("JWT_SECRET")),
		ServerPort:     strings.TrimSpace(os.Getenv("SERVER_PORT")),
		PasswordPepper: strings.TrimSpace(os.Getenv("PASSWORD_PEPPER")),
	}

	config.CORSAllowedOrigins = splitCSV(os.Getenv("CORS_ALLOWED_ORIGINS"))

	var missing []string
	required := []struct {
		key   string
		value string
	}{
		{key: "DB_HOST", value: config.DBHost},
		{key: "DB_PORT", value: config.DBPort},
		{key: "DB_USER", value: config.DBUser},
		{key: "DB_NAME", value: config.DBName},
		{key: "JWT_SECRET", value: config.JWTSecret},
	}
	for _, item := range required {
		if strings.TrimSpace(item.value) == "" {
			missing = append(missing, item.key)
		}
	}

	if len(missing) > 0 {
		return Config{}, fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	if strings.TrimSpace(config.PasswordPepper) == "" {
		return Config{}, errors.New("missing required environment variable: PASSWORD_PEPPER")
	}

	return config, nil
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func GetDb() *gorm.DB {
	return db
}