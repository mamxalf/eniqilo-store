BEGIN;

CREATE TABLE IF NOT EXISTS staffs (
                                     id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
    name varchar NOT NULL,
    phone varchar NOT NULL UNIQUE,
    password varchar NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );
-- create trigger trigger for automatically set updated_at on row update
CREATE TRIGGER set_staffs_updated_at BEFORE UPDATE ON staffs FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

COMMIT;