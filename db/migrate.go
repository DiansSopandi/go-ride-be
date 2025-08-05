package db

func Migrate() {
	MigrateUsers()
	MigrateRoles()
}

func MigrateUsers() {
	// Create the users table if it doesn't exist
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		roles VARCHAR(50) NOT NULL DEFAULT 'user',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL
	);
	`
	db := InitDatabase()
	defer db.Close()

	if _, err := db.Exec(query); err != nil {
		panic(err)
	}
}

func MigrateRoles() {
	// Create the roles table if it doesn't exist
	query := `
	CREATE TABLE IF NOT EXISTS roles (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) UNIQUE NOT NULL,
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL
	);
	`
	db := InitDatabase()
	defer db.Close()

	if _, err := db.Exec(query); err != nil {
		panic(err)
	}
}
