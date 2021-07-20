-- drop table news;

create table news
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title       varchar(255),
    date_time   timestamp,
    downloadurl text,
    imageurl    text,
    pageurl     varchar(255),
    playlist    text,
    posted      boolean                           NOT NULL default true,
    notified    bool                              NOT NULL default false,
    created_at  timestamp
);

-- alter table news
--     owner to "postgres";
--
-- alter table news
--     add column notified bool NOT NULL default false;
--
-- alter table news
--     add column created_at timestamp;
--
-- update news n
-- set created_at = n.date_time
-- where n.created_at is null;
