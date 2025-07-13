CREATE TABLE `global` (
    `id`                    INTEGER PRIMARY KEY AUTOINCREMENT,
    `writing_labels`        TEXT,
    `writing_categories`    TEXT,
    `document_labels`       TEXT,
    `matter_labels`         TEXT,
    `inspiration_labels`    TEXT
);

CREATE TABLE `writings` (
    `id`                    INTEGER PRIMARY KEY AUTOINCREMENT,
    `created`               DATE DEFAULT CURRENT_TIMESTAMP,
    `modified`              DATE DEFAULT CURRENT_TIMESTAMP,
    `title`                 VARCHAR(64),
    `subtitle`              VARCHAR(64) NULL,
    `content`               TEXT NULL,
    `published`             INTEGER NOT NULL,
    `pinned`                INTEGER NOT NULL,
    `labels`                TEXT,
    `category`              TEXT
);

CREATE TABLE `todo_docs` (
    `id`                    INTEGER PRIMARY KEY AUTOINCREMENT,
    `created`               DATE DEFAULT CURRENT_TIMESTAMP,
    `name`                  VARCHAR(64) NOT NULL,
    `comment`               VARCHAR(64) NULL DEFAULT NULL,
    `priority`              INTEGER NOT NULL,
    `labels`                TEXT NULL DEFAULT NULL
);

CREATE TABLE `cplt_docs` (
    `id`                    INTEGER PRIMARY KEY AUTOINCREMENT,
    `created`               DATE,
    `completed`             DATE DEFAULT CURRENT_TIMESTAMP,
    `name`                  VARCHAR(64) NOT NULL,
    `comment`               VARCHAR(64) NULL,
    `labels`                TEXT
);

CREATE TABLE `todo_matters` (
    `id`                    INTEGER PRIMARY KEY AUTOINCREMENT,
    `created`               DATE DEFAULT CURRENT_TIMESTAMP,
    `name`                  VARCHAR(64) NOT NULL,
    `comment`               VARCHAR(64) NULL DEFAULT NULL,
    `priority`              INTEGER NOT NULL,
    `labels`                TEXT NULL DEFAULT NULL
);

CREATE TABLE `cplt_matters` (
    `id`                    INTEGER PRIMARY KEY AUTOINCREMENT,
    `created`               DATE,
    `completed`             DATE DEFAULT CURRENT_TIMESTAMP,
    `name`                  VARCHAR(64) NOT NULL,
    `comment`               VARCHAR(64) NULL,
    `labels`                TEXT
);

CREATE TABLE `inspirations` (
    `id`                    INTEGER PRIMARY KEY AUTOINCREMENT,
    `created`               DATE DEFAULT CURRENT_TIMESTAMP,
    `content`               TEXT NOT NULL,
    `labels`                TEXT
);

CREATE TABLE `pops` (
    `id`                    INTEGER PRIMARY KEY AUTOINCREMENT,
    `created`               DATE DEFAULT CURRENT_TIMESTAMP,
    `content`               TEXT NOT NULL,
    `labels`                TEXT
);
