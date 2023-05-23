CREATE TABLE IF NOT EXISTS refresh_tokens
(
    id         serial primary key NOT NULL,
    user_id    integer            NOT NULL,
    token      text               NOT NULL,
    event_date timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS recovery_password
(
    id         serial primary key NOT NULL,
    user_id    integer            NOT NULL,
    token      text               NOT NULL,
    email      text               NOT NULL,
    event_date timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS education
(
    id   serial primary key NOT NULL,
    name text               NOT NULL
);

CREATE TABLE IF NOT EXISTS users
(
    id                   serial primary key NOT NULL,
    email                text               NOT NULL,
    password_hash        text               NOT NULL,
    event_date           timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    surname              text               NOT NULL,
    first_name           text               NOT NULL,
    patronymic           text               NOT NULL,
    city                 text               NOT NULL,
    university           text               NOT NULL,
    age                  integer            NOT NULL,
    education_id         int                NOT NULL,
    direction_internship text               NOT NULL,

    CONSTRAINT fk_detail_stage_versions_details
        FOREIGN KEY (education_id)
            REFERENCES education (id)
            ON DELETE SET NULL
);



