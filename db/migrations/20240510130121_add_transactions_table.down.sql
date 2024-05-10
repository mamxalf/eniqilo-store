BEGIN;

DROP TABLE IF EXISTS transactions CASCADE;
DROP TRIGGER IF EXISTS set_transactions_updated_at ON transactions CASCADE;

COMMIT;