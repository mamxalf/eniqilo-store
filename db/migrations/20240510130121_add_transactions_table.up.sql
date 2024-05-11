BEGIN;

CREATE TABLE IF NOT EXISTS transactions (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
    customer_id UUID NOT NULL,
    product_details JSONB NOT NULL,
    paid INT NOT NULL,
    change INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- create trigger trigger for automatically set updated_at on row update
CREATE TRIGGER set_transactions_updated_at BEFORE UPDATE ON transactions FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- alter table
ALTER TABLE transactions ADD FOREIGN KEY (customer_id) REFERENCES customers (id);


COMMIT;