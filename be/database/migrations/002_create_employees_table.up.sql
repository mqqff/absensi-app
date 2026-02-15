CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE employees (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    position SMALLINT NOT NULL DEFAULT 4 CHECK (position BETWEEN 1 AND 4),
    department SMALLINT NOT NULL DEFAULT 1 CHECK (department BETWEEN 1 AND 5),
    salary NUMERIC(10, 2) NOT NULL,
    address VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL DEFAULT 2 CHECK (status BETWEEN 1 AND 2),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TRIGGER trg_update_employees_updated_at
	BEFORE UPDATE ON employees
	FOR EACH ROW
	EXECUTE FUNCTION update_timestamp();

CREATE INDEX IF NOT EXISTS idx_employees_filter
    ON employees (department, position, status);

CREATE INDEX IF NOT EXISTS idx_employees_active_not_deleted
    ON employees (department, position)
    WHERE deleted_at IS NULL AND status = 1;

CREATE INDEX IF NOT EXISTS idx_employees_name_trgm
    ON employees
    USING gin (name gin_trgm_ops);