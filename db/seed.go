package db

func SeedRoles() {
	// Define roles to be seeded
	roles := []string{"user", "driver", "admin"}

	db := Connect()
	defer db.Close()

	for _, role := range roles {
		var exists bool
		// Check if the role already exists
		query := `SELECT EXISTS(SELECT 1 FROM roles WHERE name = $1)`
		err := db.QueryRow(query, role).Scan(&exists)
		if err != nil {
			// panic(err)
			panic("Failed to check seedRoles: " + err.Error())
		}

		// Insert the role if it does not exist
		if !exists {
			// Insert the role into the database
			query = `INSERT INTO roles (name) VALUES ($1) ON CONFLICT (name) DO NOTHING`
			if _, err := db.Exec(query, role); err != nil {
				// panic(err)
				panic("Failed insert role: " + err.Error())
			}
		}
	}
}
