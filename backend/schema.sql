


GRANT ALL PRIVILEGES ON DATABASE mydb TO myuser;

-- Products table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL,
    discount NUMERIC(10,2) DEFAULT 0,
    images TEXT[] DEFAULT ARRAY[]::TEXT[],
    sold BOOLEAN DEFAULT false
);

-- Orders table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    total_cost NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    product_ids INT[] NOT NULL,
    first_name TEXT,
    last_name TEXT,
    email TEXT
);

-- Update log
CREATE TABLE updates (
    id SERIAL PRIMARY KEY,
    table TEXT NOT NULL,
    record INT NOT NULL,
    action TEXT NOT NULL,
    old JSONB,
    new JSONB,
    updated_at TIMESTAMP DEFAULT NOW(),
)