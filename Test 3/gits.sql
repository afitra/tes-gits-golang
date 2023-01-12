-- Adminer 4.8.1 PostgreSQL 9.6.24 dump

DROP TABLE IF EXISTS "authors";
DROP SEQUENCE IF EXISTS authors_id_seq;
CREATE SEQUENCE authors_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."authors" (
    "id" bigint DEFAULT nextval('authors_id_seq') NOT NULL,
    "name" text NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    CONSTRAINT "authors_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "authors" ("id", "name", "created_at", "updated_at") VALUES
(2,	'author2',	'2023-01-11 07:47:27.552125+00',	'2023-01-11 07:47:27.552125+00'),
(3,	'author100',	'2023-01-11 07:47:50.874001+00',	'2023-01-11 08:37:44.770284+00'),
(4,	'author100',	'2023-01-11 08:37:59.19345+00',	'2023-01-11 08:37:59.19345+00'),
(8,	'publisher1',	'2023-01-11 13:59:03.522328+00',	'2023-01-11 13:59:03.522328+00'),
(9,	'publisher1',	'2023-01-11 13:59:37.96146+00',	'2023-01-11 13:59:37.96146+00'),
(10,	'author1',	'2023-01-12 04:07:01.620545+00',	'2023-01-12 04:07:01.620545+00'),
(11,	'author1',	'2023-01-12 04:07:15.450973+00',	'2023-01-12 04:07:15.450973+00'),
(12,	'author1',	'2023-01-12 04:09:14.384243+00',	'2023-01-12 04:09:14.384243+00'),
(1,	'authoryess',	'2023-01-11 07:47:24.415223+00',	'2023-01-12 06:20:23.891046+00');

DROP TABLE IF EXISTS "books";
DROP SEQUENCE IF EXISTS books_id_seq;
CREATE SEQUENCE books_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."books" (
    "id" bigint DEFAULT nextval('books_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text NOT NULL,
    "author_id" bigint NOT NULL,
    "publisher_id" bigint NOT NULL,
    CONSTRAINT "books_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_books_deleted_at" ON "public"."books" USING btree ("deleted_at");

INSERT INTO "books" ("id", "created_at", "updated_at", "deleted_at", "name", "author_id", "publisher_id") VALUES
(6,	'2023-01-12 01:54:43.65795+00',	'2023-01-12 01:54:43.65795+00',	NULL,	'publisher2',	1,	1),
(7,	'2023-01-12 01:55:01.471339+00',	'2023-01-12 01:55:01.471339+00',	NULL,	'publisher2',	1,	1),
(8,	'2023-01-12 01:56:45.875246+00',	'2023-01-12 01:56:45.875246+00',	NULL,	'publisher2',	1,	1),
(9,	'2023-01-12 01:59:32.812653+00',	'2023-01-12 01:59:32.812653+00',	'2023-01-12 02:42:41.773726+00',	'publisher2',	1,	1),
(10,	'2023-01-12 04:12:34.904369+00',	'2023-01-12 04:12:34.904369+00',	NULL,	'publisher2',	1,	1),
(5,	'2023-01-12 00:40:04.998675+00',	'2023-01-12 06:40:46.718433+00',	NULL,	'toooo',	3,	2);

DROP TABLE IF EXISTS "publishers";
DROP SEQUENCE IF EXISTS publishers_id_seq;
CREATE SEQUENCE publishers_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."publishers" (
    "id" bigint DEFAULT nextval('publishers_id_seq') NOT NULL,
    "name" text NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    CONSTRAINT "publishers_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "publishers" ("id", "name", "created_at", "updated_at") VALUES
(2,	'publisher2',	'2023-01-11 14:21:45.208017+00',	'2023-01-11 14:21:45.208017+00'),
(3,	'publisher2',	'2023-01-11 14:21:49.300869+00',	'2023-01-11 14:21:49.300869+00'),
(6,	'publisher2',	'2023-01-12 04:10:59.290922+00',	'2023-01-12 04:10:59.290922+00'),
(1,	'publishertokkkk',	'2023-01-11 14:07:01.247448+00',	'2023-01-12 07:23:19.678868+00');

DROP TABLE IF EXISTS "users";
DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE users_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."users" (
    "id" bigint DEFAULT nextval('users_id_seq') NOT NULL,
    "name" text NOT NULL,
    "email" text NOT NULL,
    "password_hash" text NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "users" ("id", "name", "email", "password_hash", "created_at", "updated_at") VALUES
(1,	'budi',	'a@mail.co',	'$2a$04$0KTZidBf2S.CjG.PuW3bkuXUsSBKS3xxjttzggB1Yu6F2Tz22Tl0y',	'2023-01-11 05:49:00.446502+00',	'2023-01-11 05:49:00.446502+00');

ALTER TABLE ONLY "public"."books" ADD CONSTRAINT "fk_books_author" FOREIGN KEY (author_id) REFERENCES authors(id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."books" ADD CONSTRAINT "fk_books_publisher" FOREIGN KEY (publisher_id) REFERENCES publishers(id) NOT DEFERRABLE;

-- 2023-01-12 09:06:38.863619+00
