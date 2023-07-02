-- v0 -> v1: Latest revision

CREATE TABLE "user" (
    -- only: postgres
    rowid BIGINT  PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    -- only: sqlite
    rowid INTEGER PRIMARY KEY,

    mxid    TEXT NOT NULL UNIQUE,
    phone   TEXT UNIQUE,
    session jsonb,

    management_room TEXT,
    space_room      TEXT,

    access_token TEXT
);

CREATE TABLE puppet (
    id               TEXT    NOT NULL,
    receiver         BIGINT  NOT NULL,
    phone            TEXT    NOT NULL,
    name             TEXT    NOT NULL,
    name_set         BOOLEAN NOT NULL DEFAULT false,
    avatar_id        TEXT    NOT NULL,
    avatar_mxc       TEXT    NOT NULL,
    avatar_set       BOOLEAN NOT NULL DEFAULT false,
    contact_info_set BOOLEAN NOT NULL DEFAULT false,

    FOREIGN KEY (receiver) REFERENCES "user"(rowid) ON DELETE CASCADE,
    UNIQUE (phone, receiver),
    PRIMARY KEY (id, receiver)
);

CREATE TABLE portal (
    id         TEXT    NOT NULL,
    receiver   BIGINT  NOT NULL,
    self_user  TEXT,
    other_user TEXT,
    mxid       TEXT    UNIQUE,
    name       TEXT    NOT NULL,
    name_set   BOOLEAN NOT NULL DEFAULT false,
    avatar_id  TEXT    NOT NULL,
    avatar_mxc TEXT    NOT NULL,
    avatar_set BOOLEAN NOT NULL DEFAULT false,
    encrypted  BOOLEAN NOT NULL DEFAULT false,
    in_space   BOOLEAN NOT NULL DEFAULT false,

    FOREIGN KEY (receiver) REFERENCES "user"(rowid) ON DELETE CASCADE,
    FOREIGN KEY (other_user, receiver) REFERENCES puppet(id, receiver) ON DELETE CASCADE,
    PRIMARY KEY (id, receiver)
);

CREATE TABLE message (
    conv_id       TEXT   NOT NULL,
    conv_receiver BIGINT NOT NULL,
    id            TEXT   NOT NULL,
    mxid          TEXT   NOT NULL UNIQUE,
    sender        TEXT   NOT NULL,
    timestamp     BIGINT NOT NULL,

    PRIMARY KEY (conv_id, conv_receiver, id),
    FOREIGN KEY (conv_id, conv_receiver) REFERENCES portal(id, receiver) ON DELETE CASCADE
);
