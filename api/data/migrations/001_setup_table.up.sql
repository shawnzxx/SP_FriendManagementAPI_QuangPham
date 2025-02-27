BEGIN;

CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    email      TEXT NOT NULL UNIQUE,
    phone      TEXT,
    is_active  BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS relations
(
    id         SERIAL PRIMARY KEY,
    requester_id     INTEGER NOT NULL REFERENCES users(id),
    addressee_id     INTEGER NOT NULL REFERENCES users(id),
    requester_email     TEXT NOT NULL,
    addressee_email     TEXT NOT NULL,
    relation_type INTEGER DEFAULT 0,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

INSERT INTO "public"."users" ("id",
                              "email",
                              "phone",
                              "is_active")
VALUES (10000, 'andy@example.com', '123456', 'true'),
       (10001, 'john@example.com', '123456', 'true'),
       (10003, 'common@example.com', '123456', 'true'),
       (10004, 'lisa@example.com', '123456', 'true');

COMMIT;



