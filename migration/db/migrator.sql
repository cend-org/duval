drop database if exists cend;
create database cend;
use cend;

create table user
(
    id                     int auto_increment
        primary key,
    created_at             datetime     default CURRENT_TIMESTAMP     not null,
    updated_at             datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at             datetime     default '0000-00-00 00:00:00' null,
    name                   varchar(500) default ''                    null,
    family_name            varchar(500) default ''                    null,
    nick_name              varchar(100) default ''                    null,
    email                  varchar(100) default ''                    null,
    matricule              varchar(32)  default ''                    null,
    age                    int          default 0                     null,
    birth_date             datetime     default '0000-00-00 00:00:00' null,
    sex                    int          default 0                     null,
    lang                   int          default 0                     null,
    status                 int          default 0                     null,
    profile_image_xid      varchar(250) default ''                    null,
    description            varchar(500) default ''                    null,
    cover_text             varchar(500) default ''                    null,
    profile                varchar(500) default ''                    null,
    experience_detail      varchar(500) default ''                    null,
    additional_description varchar(500) default ''                    null,
    add_on_title           varchar(250) default ''                    null,
    constraint user_pk
        unique (email)
);

create table authorization
(
    id         int auto_increment
        primary key,
    created_at datetime default CURRENT_TIMESTAMP     not null,
    updated_at datetime default CURRENT_TIMESTAMP     not null,
    deleted_at datetime default '0000-00-00 00:00:00' null,
    user_id    int      default 0                     null,
    level      int      default 0                     null,
    constraint authorization_pk
        unique (user_id, level),
    constraint authorization_user_id_fk
        foreign key (user_id) references user (id)
            on update cascade on delete cascade
);

create table password
(
    id         int auto_increment
        primary key,
    created_at datetime     default CURRENT_TIMESTAMP     not null,
    updated_at datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at datetime     default '0000-00-00 00:00:00' null,
    user_id    int          default 0                     null,
    hash       varchar(500) default ''                    null,
    constraint password_user_id_fk
        foreign key (user_id) references user (id)
);

create index password_user_id_index
    on password (user_id);

create table media
(
    id           int auto_increment
        primary key,
    created_at   datetime     default CURRENT_TIMESTAMP     not null,
    updated_at   datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at   datetime     default '0000-00-00 00:00:00' null,
    file_name    varchar(500) default ''                    null,
    extension    varchar(10)  default ''                    null,
    xid          varchar(100) default ''                    null
);

create table media_thumb
(
    id         int auto_increment
        primary key,
    created_at datetime     default CURRENT_TIMESTAMP,
    updated_at datetime     default CURRENT_TIMESTAMP,
    deleted_at datetime     default '0000-00-00 00:00:00',
    extension  varchar(10)  default '',
    media_xid  varchar(100) default '',
    xid        varchar(500) default ''
);

alter table media_thumb
    add constraint media_thumb_media_xid_fk
        foreign key (media_xid) references media (xid);

create table user_media_detail
(
    id            int auto_increment primary key,
    created_at    datetime            default CURRENT_TIMESTAMP,
    updated_at    datetime            default CURRENT_TIMESTAMP,
    deleted_at    datetime            default '0000-00-00 00:00:00',
    owner_id      int                 default 0,
    document_type int                 default 0,
    document_xid  varchar(100) unique default ''
);

alter table user_media_detail
    add constraint user_media_detail_user_id_fk
        foreign key (owner_id) references user (id);

