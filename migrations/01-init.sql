-- +migrate Up
CREATE SCHEMA IF NOT EXISTS app;

CREATE TABLE IF NOT EXISTS app.user (
	id uuid NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,

	username varchar(255) NOT NULL,
	email varchar(255) NOT NULL,

	CONSTRAINT unique_user_email UNIQUE (email),
	CONSTRAINT user_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS app.auth_provider (
	id uuid NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,

	user_id UUID NOT NULL REFERENCES app.user(id) ON DELETE CASCADE,
	method VARCHAR(50) NOT NULL,
	provider_user_id VARCHAR(255),
	password_hash VARCHAR(255),
	two_factor_authentication BOOLEAN DEFAULT FALSE,

	CONSTRAINT unique_provider_user UNIQUE (method, provider_user_id),
	CONSTRAINT auth_provider_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS app.session (
	id uuid NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,

	user_id UUID NOT NULL REFERENCES app.user(id) ON DELETE CASCADE,  

	refresh_token VARCHAR(512) NOT NULL UNIQUE,           
	issued_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),         
	expires_at TIMESTAMPTZ NOT NULL,                      
	revoked BOOLEAN NOT NULL DEFAULT FALSE,               

	user_agent VARCHAR(255),                              

	CONSTRAINT unique_token UNIQUE (refresh_token),                
	CONSTRAINT refresh_token_pkey PRIMARY KEY (id)
);

-- +migrate Down
DROP SCHEMA app cascade;
