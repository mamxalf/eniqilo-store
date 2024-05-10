BEGIN;

DROP TABLE IF EXISTS customers CASCADE;
DROP TRIGGER IF EXISTS set_customers_updated_at ON customers CASCADE;

COMMIT;