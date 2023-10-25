package migration

var schema = `
	DO $$
	BEGIN

		IF NOT EXISTS (SELECT 1 FROM pg_type  WHERE typname = 'gender') THEN
			CREATE TYPE gender AS ENUM('male', 'female');
		END IF;

		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
			CREATE TYPE role AS ENUM('admin', 'manager', 'user');
		END IF;	

		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_credit') THEN
			CREATE TYPE status_credit AS ENUM('waiting', 'processed', 'ongoing', 'done');
		END IF;	

		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'credit_type') THEN
			CREATE TYPE credit_type AS ENUM('QORD');
		END IF;	

		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_payment') THEN
			CREATE TYPE status_payment AS ENUM('success', 'failed');
		END IF;	

	END $$;

	CREATE TABLE IF NOT EXISTS users (
		id SERIAL NOT NULL,
		email VARCHAR (255) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		phone_number VARCHAR (255) UNIQUE NOT NULL,
		created_at TIMESTAMPTZ NOT NULL,
		updated_at TIMESTAMPTZ NOT NULL,
		deleted_at TIMESTAMPTZ,
		PRIMARY KEY (id)
	);

	CREATE TABLE IF NOT EXISTS user_details (
		id SERIAL NOT NULL,
		user_id INTEGER NOT NULL,
		first_name VARCHAR (255) NOT NULL,
		last_name VARCHAR (255),
		address VARCHAR (255),
		gender gender,
		date_of_birth DATE,
		img_ktp BYTEA,
		role role NOT NULL,
		PRIMARY KEY (id),
		FOREIGN KEY (user_id) REFERENCES users (id)
	);

	CREATE TABLE IF NOT EXISTS credits (
		id SERIAL NOT NULL,
		user_id INTEGER NOT NULL,
		credit_type credit_type,
		name VARCHAR (255) NOT NULL,
		total_transaction NUMERIC(20,2) DEFAULT 0 NOT NULL,
		tenor VARCHAR (255) NOT NULL,
		total_credit NUMERIC(20,2) NOT NULL,
		status status_credit,
		created_at TIMESTAMPTZ NOT NULL,
		updated_at TIMESTAMPTZ NOT NULL,
		deleted_at TIMESTAMPTZ,
		PRIMARY KEY (id),
		FOREIGN KEY (user_id) REFERENCES users (id)
	);

	CREATE TABLE IF NOT EXISTS payments (
		id SERIAL NOT NULL,
		credit_id INTEGER NOT NULL,
		amount NUMERIC(20,2) NOT NULL,
		status status_payment,
		created_at TIMESTAMPTZ NOT NULL,
		updated_at TIMESTAMPTZ NOT NULL,
		deleted_at TIMESTAMPTZ,
		PRIMARY KEY (id),
		FOREIGN KEY (credit_id) REFERENCES credits (id)
	);

`
