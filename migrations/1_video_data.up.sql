CREATE TABLE IF NOT EXISTS videoData
(
    id             text        not null PRIMARY KEY,
    title          text        not null,
    description    text        not null,
    published_time timestamptz not null,
    url            text        not null,
    UNIQUE (url)
);