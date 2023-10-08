create table roles
(
    id         bigserial   not null primary key,
    role       text        not null,
    active     boolean     not null default true,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz not null default current_timestamp
);

create table notifications
(
    id           bigserial                 not null primary key,
    notification text                      not null,
    recipient_id int references users (id) not null,
    active       boolean                   not null default true,
    created_at   timestamptz               not null default current_timestamp,
    updated_at   timestamptz               not null default current_timestamp,
    deleted_at   timestamptz               not null default current_timestamp
);

alter table notifications
    add is_read boolean not null default false;

create table users
(
    id         bigserial                 not null primary key,
    name       text                      not null,
    age        int                       not null,
    login      text                      not null,
    password   text                      not null,
    role_id    int references roles (id) not null default 2,
    active     boolean                   not null default true,
    created_at timestamptz               not null default current_timestamp,
    updated_at timestamptz               not null default current_timestamp,
    deleted_at timestamptz               not null default current_timestamp

);

create table tokens
(
    id         bigserial                 not null primary key,
    token      text                      not null,
    user_id    int references users (id) not null,
    expire     timestamptz               not null default current_timestamp + interval '3 hour',
    created_at timestamptz               not null default current_timestamp,
    updated_at timestamptz               not null default current_timestamp,
    deleted_at timestamptz               not null default current_timestamp
);

create table genres
(
    id         bigserial   not null primary key,
    name       text        not null,
    active     boolean     not null default true,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz not null default current_timestamp
);

create table content_type
(
    id         bigserial   not null primary key,
    name       text        not null,
    active     boolean     not null default true,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz not null default current_timestamp
);


create table if not exists content
(
    id               bigserial                        not null primary key,
    name             text                             not null,
    content_type_id  int references content_type (id) not null,
    description      text                             not null default 'no description',
    production_year  int                              not null,
    age_limit        int                              not null,
    producers        text                             not null,
    directors        text                             not null,
    actors           text                             not null,
    main_characters  text                             not null,
    duration         text                             not null,
    genre_id         int references genres (id)       not null,
    cover_image_name text                             not null default 'no image',
    active           boolean                          not null default true,
    created_at       timestamptz                      not null default current_timestamp,
    updated_at       timestamptz                      not null default current_timestamp,
    deleted_at       timestamptz                      not null default current_timestamp
);

alter table content
    add cover_image_name text not null default 'no image';


create table links
(
    id         bigserial                   not null primary key,
    active     boolean                     not null default true,
    alfa       text                        not null default 'no link',
    kinopoisk  text                        not null default 'no link',
    okko       text                        not null default 'no link',
    wink       text                        not null default 'no link',
    content_id int references content (id) not null default 0,
    created_at timestamptz                 not null default current_timestamp,
    updated_at timestamptz                 not null default current_timestamp,
    deleted_at timestamptz                 not null default current_timestamp
);

create table user_recommendations
(
    id         bigserial                   not null primary key,
    active     boolean                     not null default true,
    user_id    int references users (id)   not null,
    content_id int references content (id) not null,
    created_at timestamptz                 not null default current_timestamp,
    updated_at timestamptz                 not null default current_timestamp,
    deleted_at timestamptz                 not null default current_timestamp
);



create table recommendation
(
    id         bigserial   not null primary key,
    content_id int         not null references content (id),
    active     boolean     not null default true,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz not null default current_timestamp
);

create table playlist_types
(
    id         bigserial   not null primary key,
    name       text        not null,
    active     boolean     not null default true,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz not null default current_timestamp
);

create table playlists
(
    id          bigserial                          not null primary key,
    active      boolean                            not null default true,
    playlist_id int references playlist_types (id) not null,
    user_id     int references users (id)          not null,
    content_id  int references content (id)        not null,
    created_at  timestamptz                        not null default current_timestamp,
    updated_at  timestamptz                        not null default current_timestamp,
    deleted_at  timestamptz                        not null default current_timestamp
);