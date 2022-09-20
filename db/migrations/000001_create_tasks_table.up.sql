CREATE TABLE IF NOT EXISTS tasks(
    id UUID NOT NULL DEFAULT gen_random_uuid (),
    title varchar(1000) NULL,
    description varchar(1000) NULL,
    created_at TIMESTAMP NULL ,
    image text NULL ,
    status VARCHAR(20) NULL ,
    PRIMARY KEY(id)
);