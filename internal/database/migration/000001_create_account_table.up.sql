CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS accounts(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email text NOT NULL,
    password text NOT NULL
);
