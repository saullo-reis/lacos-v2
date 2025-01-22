package migrations

import (
	"fmt"
	_ "github.com/lib/pq"
	dbconfig "lacosv2.com/src/database/config"
)

func CreateTables() {

	queryToCreateTables := `
	CREATE TABLE IF NOT EXISTS persons (
	id_person SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    rg VARCHAR(20),
    cpf VARCHAR(14),
    cad_unico VARCHAR(20),
    nis VARCHAR(20),
    school VARCHAR(255),
    address VARCHAR(255),
    address_number VARCHAR(10),
    blood_type VARCHAR(3),
    neighborhood VARCHAR(100),
    city VARCHAR(100),
    cep VARCHAR(9),
    home_phone VARCHAR(15),
    cell_phone VARCHAR(15),
    contact_phone VARCHAR(15),
	created_date timestamp not null default CURRENT_TIMESTAMP,
    email VARCHAR(255),
    current_age INT,
	active VARCHAR(1)
	);

	CREATE TABLE IF NOT EXISTS responsible_person (
		id_responsible SERIAL PRIMARY KEY,
		id_person INT REFERENCES persons(id_person),
		name VARCHAR(255) NOT NULL,
		relationship VARCHAR(50),
		rg VARCHAR(20),
		cpf VARCHAR(14),
		cell_phone VARCHAR(15)
	);

	CREATE TABLE IF NOT EXISTS period (
		id_period INT PRIMARY KEY,
		name VARCHAR(50) NOT NULL
	);

	INSERT INTO period (id_period, name) VALUES
	(1, 'Manhã'),
	(2, 'Tarde'),
	(3, 'Noite') 
	ON CONFLICT (id_period) DO NOTHING;

	CREATE TABLE IF NOT EXISTS activity_list (
		id_activity SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		hour_start VARCHAR(50) NOT NULL,
		hour_end VARCHAR(50) NOT NULL,
		id_period INT REFERENCES period(id_period)
	);

	CREATE TABLE IF NOT EXISTS activities (
		id_activities SERIAL PRIMARY KEY,
		id_activity INT REFERENCES activity_list(id_activity),
		id_person INT REFERENCES persons(id_person)
	);

	CREATE TABLE IF NOT EXISTS users(
		id_user SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL
	);

	INSERT INTO users (username, password)
	SELECT 'admin', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918'
	WHERE NOT EXISTS (
		SELECT 1 FROM users WHERE username = 'admin'
	);

	`
	fmt.Println(dbconfig.DataSourceName)
	db, err := dbconfig.ConnectDB()
	if err != nil {
		panic("Error connecting in database: " + err.Error())
	}
	fmt.Println("Connection with database sucess")
	defer db.Close()

	
	_, err = db.Exec(queryToCreateTables)
	if err != nil {
		panic("Error in migrations: " + err.Error())
	}
}
