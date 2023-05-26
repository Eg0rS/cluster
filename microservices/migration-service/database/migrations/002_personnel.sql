-- +goose Up
create table if not exists PersonnelTest
(
    id               serial primary key,
    test_title       varchar not null,
    test_description varchar not null,
    test_type        varchar not null
);

create table if not exists TestQuestion
(
    id            serial primary key,
    question_text varchar not null,
    test_id       int     not null,

    foreign key (test_id) references PersonnelTest (id)
);

create table if not exists QuestionAnswer
(
    id          serial primary key,
    answer      varchar not null,
    is_right    boolean not null,
    question_id int     not null,

    foreign key (question_id) references TestQuestion (id)
);

create table if not exists Organizations
(
    id      serial primary key,
    org_name    text not null,
    address text not null,
    first_coordinates decimal not null,
    second_coordinates decimal not null
);

create table if not exists PersonnelRequest
(
    id                  serial primary key,
    request_title       varchar not null,
    request_description varchar not null,
    test_id             int     not null,
    user_id             int     not null,
    organization_id     int     not null,

    foreign key (user_id) references Users (id),
    foreign key (test_id) references PersonnelTest (id),
    foreign key (organization_id) references Organizations (id)
);

create table if not exists TestFile
(
    id       serial primary key,
    filename varchar not null,
    test_id  int     not null,

    foreign key (test_id) references PersonnelTest (id)
);
-- +goose Down
