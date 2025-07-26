package config

import "os"

var (
	DATABASE_URL string
	JWT_SECRET   string
)

func init() {
	// DATABASE_URL = getEnv("DATABASE_URL", "host=localhost user=postgres password=password dbname=doctor_db port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	DATABASE_URL = getEnv("DATABASE_URL", "postgresql://postgres:sanyuanya212223@@db.rtbwvdoclefhiqbteevi.supabase.co:5432/postgres")
	JWT_SECRET = getEnv("JWT_SECRET", "Tyokb2KA5tzKqMzojpAs9k9XvaRbVddsfpvaKowKweWDRhYZBT0G06tpf1y2BQhYtfZWJQ8w3EDbXPxx0TfCHKbEWhKJH3LMPXPC33FG4etifcrrHdwdAFoxWTieHurjrxyeFrZTTc00U26A3EkkhFcqdUbJNTp4Ldwm3W9QxYfAckt4Ziydfyvw2hsdZUqt5HG471AL")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
