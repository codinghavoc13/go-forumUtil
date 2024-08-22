package main

var CREATE_USER_TABLE = `
CREATE TABLE IF NOT EXISTS react_forum.users
(
    user_id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    first_name character varying(255) COLLATE pg_catalog."default",
    last_name character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS react_forum.users
    OWNER to postgres;`

var CREATE_ROOMS_TABLE = `
CREATE TABLE IF NOT EXISTS react_forum.rooms
(
    room_id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    room_title character varying(255) COLLATE pg_catalog."default",
    room_description character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT rooms_pkey PRIMARY KEY (room_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS react_forum.rooms
    OWNER to postgres;`

var CREATE_POSTS_TABLE = `
CREATE TABLE IF NOT EXISTS react_forum.posts
(
    post_id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    date_last_updated timestamp(6) without time zone,
    number_responses bigint,
    post_date timestamp(6) without time zone,
    post_text character varying(255) COLLATE pg_catalog."default",
    post_title character varying(255) COLLATE pg_catalog."default",
    user_id bigint,
    room_id bigint,
    CONSTRAINT posts_pkey PRIMARY KEY (post_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS react_forum.posts
    OWNER to postgres;`

var CREATE_RESPONSE_TABLE = `
CREATE TABLE IF NOT EXISTS react_forum.responses
(
    response_id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    post_id bigint,
    response_date timestamp(6) without time zone,
    response_text character varying(255) COLLATE pg_catalog."default",
    user_id bigint,
    CONSTRAINT responses_pkey PRIMARY KEY (response_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS react_forum.responses
    OWNER to postgres;`

var INSERT_ROOMS = `
INSERT INTO react_forum.rooms(room_title, room_description)
	VALUES
	('Star Wars', 'All things Star Wars, legends, or cannon'),
	('Star Trek','All things Star Trek'),
	('Coding - Backend Development','Discussion place for backend specific developerment of any language'),
	('Coding - Frontend Development','Discussion place for frontend specific developerment of any language'),
	('Coding - General','Discussion place for developers of any language'),
	('MCU','All things regarding the Marvel Cinematic Universe movies and shows'),
	('DCEU','All things regarding the DC Extended Universe movies and shows');`

var INSERT_POSTS = `
INSERT INTO react_forum.posts(
	date_last_updated, number_responses, post_date, post_text, post_title, user_id, room_id)
	VALUES `

var INSERT_RESPONSES = `
INSERT INTO notification_demo.responses(
	post_id, response_date, response_text, user_id)
	VALUES (?, ?, ?, ?);`

var INSERT_USERS = `
INSERT INTO notification_demo.users(first_name, last_name)
	VALUES ('Anne', 'Obrien'),
    ('Ira', 'Sprague'),
    ('Beverly', 'Wheeler'),
    ('Marianne', 'Shultz'),
    ('Rosa', 'McBride'),
    ('Willard', 'Neal'),
    ('Pearl', 'Berry'),
    ('Rose', 'Downey'),
    ('Justin', 'Durham'),
    ('Miriam', 'Gutierrez'),
    ('Rodolfo', 'Young'),
    ('Roland', 'Bell'),
    ('Thelma', 'Newton'),
    ('Polly', 'Vazquez'),
    ('Adrian', 'Foster'),
    ('Timothy', 'Leon'),
    ('Mathew', 'Driscoll'),
    ('Calvin', 'Kelley'),
    ('Gwendolyn', 'Stone'),
    ('Kris', 'McKnight');`
