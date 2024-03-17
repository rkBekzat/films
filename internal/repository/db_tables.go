package repository

// creating table
const (
	usersQuey = `CREATE TABLE IF NOT EXISTS users(
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		username VARCHAR(50),
		email VARCHAR(50) UNIQUE,
		gender VARCHAR(50),
		password VARCHAR(256),
		role VARCHAR(50)
	);`

	filmsQuey = `CREATE TABLE IF NOT EXISTS film(
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title VARCHAR(50),
		description VARCHAR(500),
		release_date DATE,
		rating real 
	);`

	actorsQuey = `CREATE TABLE IF NOT EXISTS actor(
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(50),
		gender VARCHAR(50),
		birth_date DATE 
	);`

	participantsQuey = `CREATE TABLE IF NOT EXISTS participants(
		film_id UUID,
		FOREIGN KEY (film_id) REFERENCES film (id) ON DELETE SET NULL,
		actor_id UUID,
		FOREIGN KEY (actor_id) REFERENCES actor (id) ON DELETE SET NULL
	);`
)

var (
	quries = []string{usersQuey, filmsQuey, actorsQuey, participantsQuey}
)
