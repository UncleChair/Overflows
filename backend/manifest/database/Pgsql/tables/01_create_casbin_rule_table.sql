CREATE TABLE casbin_rule
(
    id SERIAL PRIMARY KEY,
    p_type VARCHAR(10),
    v0 VARCHAR(25),
    v1 VARCHAR(25),
    v2 VARCHAR(25),
    v3 VARCHAR(25),
    v4 VARCHAR(25),
    v5 VARCHAR(25),
    v6 VARCHAR(25),
    v7 VARCHAR(25),
    CONSTRAINT idx_casbin_rule UNIQUE (p_type, v0, v1, v2, v3, v4, v5, v6, v7)
) TABLESPACE pg_default;