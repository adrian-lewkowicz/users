--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

-- Started on 2024-10-15 00:00:02

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 5 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: gorm
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO gorm;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 16401)
-- Name: users; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name text,
    email text,
    age bigint
);


ALTER TABLE public.users OWNER TO gorm;

--
-- TOC entry 215 (class 1259 OID 16400)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO gorm;

--
-- TOC entry 4839 (class 0 OID 0)
-- Dependencies: 215
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 4688 (class 2604 OID 16404)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 4690 (class 2606 OID 16408)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2024-10-15 00:00:02

--
-- PostgreSQL database dump complete
--

