create table news
(
    id                     serial not null
        constraint news_pkey
            primary key,
    country                varchar(255),
    date_time              timestamp,
    downloadurl            text,
    format                 varchar(255),
    genre                  varchar(255),
    imageurl               text,
    news_type              varchar(255),
    pageurl                varchar(255),
    playlist               text,
    title                  varchar(255),
    posted                 boolean default true,
    path_to_file           varchar(255),
    real_file_download_url varchar(255)
);

alter table news
    owner to "postgres";

