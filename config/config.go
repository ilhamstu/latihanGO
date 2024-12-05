package config

import "os"

var JwtSecret = os.Getenv("JWT_SECRET")
