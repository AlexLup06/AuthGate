# AuthGate Database Migrations

This directory contains the **database schema migrations required by AuthGate**.

AuthGate **does not manage database schema automatically**.  
Migrations must be applied **explicitly** before running the AuthGate server.

---

## Overview

- AuthGate requires a PostgreSQL database
- The database schema is versioned using SQL migrations
- Migrations are applied using `sql-migrate`
- AuthGate **will refuse to start** if the database schema version is incompatible

This design is intentional and ensures:
- explicit infrastructure changes
- safe upgrades and rollbacks
- no hidden side effects at runtime

---

## Migration Tooling

AuthGate provides a **dedicated migration runner Docker image** built from this directory.

The migration image:
- contains all SQL migrations
- contains the `sql-migrate` CLI
- does **not** start AuthGate
- exits after applying migrations

---

## Running Migrations (Docker – Recommended)

### Prerequisites

You must provide database connection details via environment variables, for example:

POSTGRESQL_HOST=localhost  
POSTGRESQL_PORT=5432  
POSTGRESQL_DATABASE=authgate  
POSTGRESQL_USERNAME=authgate  
POSTGRESQL_PASSWORD=secret  

### Apply migrations

docker run --rm \
  --env-file .env \
  ghcr.io/<your-org>/authgate-migrate:<version> up

This command:
- connects to the database
- applies all pending migrations
- exits immediately

---

## Checking Migration Status

docker run --rm \
  --env-file .env \
  ghcr.io/<your-org>/authgate-migrate:<version> status

---

## Rolling Back (Advanced)

docker run --rm \
  --env-file .env \
  ghcr.io/<your-org>/authgate-migrate:<version> down

⚠️ **Warning:** Rolling back migrations may result in data loss.  
Only perform rollbacks if you fully understand the impact.

---

## Schema Compatibility Enforcement

AuthGate **validates schema compatibility on startup**.

If the database schema version does not match the version required by the AuthGate server, startup will fail with an error similar to:

database schema version mismatch  
current: 002_users  
required: 003_sessions  

In this case, apply the correct migrations before starting AuthGate.

---

## Design Principles

- AuthGate **does not own your database**
- Schema changes are **explicit**
- Migrations are **operator-controlled**
- Runtime behavior is **deterministic and safe**

This is infrastructure software — not an application that mutates shared state automatically.

---

## Summary

- Migrations are required to run AuthGate
- Migrations must be applied explicitly
- AuthGate will not auto-migrate
- Schema compatibility is enforced at startup

If AuthGate starts successfully, the database schema is guaranteed to be correct.
