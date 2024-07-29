DO $$
    BEGIN
        BEGIN
            EXECUTE 'CREATE ROLE postgres WITH LOGIN SUPERUSER PASSWORD ''12345''';
        EXCEPTION
            WHEN duplicate_object THEN
                RAISE NOTICE 'Role postgres already exists.';
        END;
    END $$;

CREATE TABLE IF NOT EXISTS message
(
    id   SERIAL PRIMARY KEY,
    data TEXT DEFAULT ''::text NOT NULL
);

ALTER TABLE message
    OWNER TO postgres;