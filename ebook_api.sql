-- Database: ebook_appDB
CREATE DATABASE "ebook_appDB"
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Create tables: users, categories, favorites, ebooks

    -- categories    ->
-- Table: public.categories
CREATE TABLE IF NOT EXISTS public.categories
(
    id integer NOT NULL,
    name "char",
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


    --  ebooks   ->
-- Table: public.ebooks
CREATE TABLE IF NOT EXISTS public.ebooks
(
    id integer NOT NULL,
    title "char",
    author "char",
    categorie_id integer NOT NULL,
    CONSTRAINT ebooks_pkey PRIMARY KEY (id),
    CONSTRAINT fk_category FOREIGN KEY (categorie_id)
        REFERENCES public.categories (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.ebooks
    OWNER to postgres;


    --  favorites   ->
-- Table: public.favorites
CREATE TABLE IF NOT EXISTS public.favorites
(
    id integer NOT NULL,
    user_id integer NOT NULL,
    ebook_id integer NOT NULL,
    CONSTRAINT favorites_pkey PRIMARY KEY (id),
    CONSTRAINT fk_ebook FOREIGN KEY (ebook_id)
        REFERENCES public.ebooks (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_user FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.favorites
    OWNER to postgres;