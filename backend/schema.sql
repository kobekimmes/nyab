-- Products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL,
    discount NUMERIC(10,2) DEFAULT 0,
    images TEXT[] DEFAULT ARRAY[]::TEXT[],
    sold BOOLEAN DEFAULT false
);

-- Orders table
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    total_cost NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    product_ids INT[] NOT NULL,
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    paid BOOLEAN DEFAULT false,
    payment_id TEXT
);

-- Update log
CREATE TABLE IF NOT EXISTS updates (
    id SERIAL PRIMARY KEY,
    table_name TEXT NOT NULL,
    record INT NOT NULL,
    action_method TEXT NOT NULL,
    old JSONB,
    new JSONB,
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Migration log
CREATE TABLE IF NOT EXISTS migrations (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
)