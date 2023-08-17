-- Database: ebook_appDB

-- DROP DATABASE "ebook_appDB";

CREATE DATABASE "ebook_appDB"
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Create tables: users, categories, favorites, ebooks

-- Table: public.categories

-- DROP TABLE public.categories;

CREATE TABLE IF NOT EXISTS public.categories
(
    id integer NOT NULL DEFAULT nextval('categories_id_seq'::regclass),
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT categories_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.categories
    OWNER to postgres;

    --  users   ->
-- Table: public.users
CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL,
    username "char",
    email "char",
    CONSTRAINT users_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to postgres;


-- Table: public.ebooks

-- DROP TABLE public.ebooks;

CREATE TABLE IF NOT EXISTS public.ebooks
(
    id integer NOT NULL DEFAULT nextval('ebooks_id_seq'::regclass),
    title character varying(255) COLLATE pg_catalog."default" NOT NULL,
    author character varying(255) COLLATE pg_catalog."default",
    categorie_id integer NOT NULL,
    CONSTRAINT ebooks_pkey PRIMARY KEY (id),
    CONSTRAINT ebooks_categorie_id_fkey FOREIGN KEY (categorie_id)
        REFERENCES public.categories (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.ebooks
    OWNER to postgres;


-- Table: public.favorites

-- DROP TABLE public.favorites;

CREATE TABLE IF NOT EXISTS public.favorites
(
    id integer NOT NULL DEFAULT nextval('favorites_id_seq'::regclass),
    user_id integer NOT NULL,
    ebook_id integer NOT NULL,
    CONSTRAINT favorites_pkey PRIMARY KEY (id),
    CONSTRAINT favorites_ebook_id_fkey FOREIGN KEY (ebook_id)
        REFERENCES public.ebooks (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT favorites_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.favorites
    OWNER to postgres;

-- Table: public.users

-- DROP TABLE public.users;

CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    username character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to postgres;