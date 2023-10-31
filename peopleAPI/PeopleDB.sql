--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0 (Homebrew)
-- Dumped by pg_dump version 16.0 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.people (
    id integer,
    firstname character varying(255),
    lastname character varying(255),
    age integer
);

--
-- Sample Row Data
-- 

COPY public.people (id, firstname, lastname, age) FROM stdin;
1	john	Cooper	40
2	Dave	Supra	23
3	Yennifer	Sicklemore	45
4	Julieta	Verretti	25
5	Loto	Picard	34
\.


--
-- PostgreSQL database dump complete
--

