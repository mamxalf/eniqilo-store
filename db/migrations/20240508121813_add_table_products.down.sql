BEGIN;

DROP TABLE IF EXISTS products CASCADE;
DROP TRIGGER IF EXISTS set_products_updated_at ON products CASCADE;

COMMIT;