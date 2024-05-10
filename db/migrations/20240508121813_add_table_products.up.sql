BEGIN;

CREATE TABLE IF NOT EXISTS products (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
    staff_id UUID NOT NULL,
    name varchar NOT NULL,
    sku varchar NOT NULL,
    category varchar NOT NULL,
    imageurl varchar NOT NULL,
    notes TEXT NOT NULL,
    price int NOT NULL DEFAULT 1,
    stock int NOT NULL DEFAULT 1,
    location TEXT,
    isavailable boolean NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- create trigger trigger for automatically set updated_at on row update
CREATE TRIGGER set_products_updated_at BEFORE UPDATE ON products FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- alter table
ALTER TABLE products ADD FOREIGN KEY (staff_id) REFERENCES staffs (id);

COMMIT;