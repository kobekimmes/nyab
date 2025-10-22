



CREATE TABLE products {
    id SERIAL PRIMARY KEY AUTO,
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL,
    discount NUMERIC(10,2) DEFAULT 0,
    images TEXT[] DEFAULT ARRAY[]::TEXT[]
}

CREATE TABLE order_history {
    id SERIAL PRIMARY KEY AUTO,
    timestamp TIMESTAMP

}