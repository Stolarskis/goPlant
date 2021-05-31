--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
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

UPDATE pg_catalog.pg_database SET datistemplate = false WHERE datname = 'template1';
DROP DATABASE template1;
--
-- Name: template1; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE template1 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE template1 OWNER TO postgres;

\connect template1

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
-- Name: DATABASE template1; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE template1 IS 'default template for new databases';


--
-- Name: template1; Type: DATABASE PROPERTIES; Schema: -; Owner: postgres
--

ALTER DATABASE template1 IS_TEMPLATE = true;


\connect template1

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
-- Name: DATABASE template1; Type: ACL; Schema: -; Owner: postgres
--

REVOKE CONNECT,TEMPORARY ON DATABASE template1 FROM PUBLIC;
GRANT CONNECT ON DATABASE template1 TO PUBLIC;


--
-- PostgreSQL database dump complete
--

--
-- Database "gobase" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2 (Debian 13.2-1.pgdg100+1)
-- Dumped by pg_dump version 13.2 (Debian 13.2-1.pgdg100+1)

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
-- Name: gobase; Type: DATABASE; Schema: -; Owner: postgres
--

-- CREATE DATABASE gobase WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE gobase OWNER TO postgres;

\connect gobase

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
-- Name: goplant; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA goplant;


ALTER SCHEMA goplant OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: airtemperature; Type: TABLE; Schema: goplant; Owner: postgres
--

CREATE TABLE goplant.airtemperature (
    id integer NOT NULL,
    value integer NOT NULL,
    createdate date DEFAULT CURRENT_DATE,
    recordtime date
);


ALTER TABLE goplant.airtemperature OWNER TO postgres;

--
-- Name: airtemperature_id_seq; Type: SEQUENCE; Schema: goplant; Owner: postgres
--

CREATE SEQUENCE goplant.airtemperature_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE goplant.airtemperature_id_seq OWNER TO postgres;

--
-- Name: airtemperature_id_seq; Type: SEQUENCE OWNED BY; Schema: goplant; Owner: postgres
--

ALTER SEQUENCE goplant.airtemperature_id_seq OWNED BY goplant.airtemperature.id;


--
-- Name: humidity; Type: TABLE; Schema: goplant; Owner: postgres
--

CREATE TABLE goplant.humidity (
    id integer NOT NULL,
    value integer NOT NULL,
    createdate date DEFAULT CURRENT_DATE,
    recordtime date
);


ALTER TABLE goplant.humidity OWNER TO postgres;

--
-- Name: humidity_id_seq; Type: SEQUENCE; Schema: goplant; Owner: postgres
--

CREATE SEQUENCE goplant.humidity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE goplant.humidity_id_seq OWNER TO postgres;

--
-- Name: humidity_id_seq; Type: SEQUENCE OWNED BY; Schema: goplant; Owner: postgres
--

ALTER SEQUENCE goplant.humidity_id_seq OWNED BY goplant.humidity.id;


--
-- Name: light; Type: TABLE; Schema: goplant; Owner: postgres
--

CREATE TABLE goplant.light (
    id integer NOT NULL,
    value integer NOT NULL,
    createdate date DEFAULT CURRENT_DATE,
    recordtime date
);


ALTER TABLE goplant.light OWNER TO postgres;

--
-- Name: light_id_seq; Type: SEQUENCE; Schema: goplant; Owner: postgres
--

CREATE SEQUENCE goplant.light_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE goplant.light_id_seq OWNER TO postgres;

--
-- Name: light_id_seq; Type: SEQUENCE OWNED BY; Schema: goplant; Owner: postgres
--

ALTER SEQUENCE goplant.light_id_seq OWNED BY goplant.light.id;


--
-- Name: soilmoisture; Type: TABLE; Schema: goplant; Owner: postgres
--

CREATE TABLE goplant.soilmoisture (
    id integer NOT NULL,
    value integer NOT NULL,
    createdate date DEFAULT CURRENT_DATE,
    recordtime date
);


ALTER TABLE goplant.soilmoisture OWNER TO postgres;

--
-- Name: soilmoisture_id_seq; Type: SEQUENCE; Schema: goplant; Owner: postgres
--

CREATE SEQUENCE goplant.soilmoisture_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE goplant.soilmoisture_id_seq OWNER TO postgres;

--
-- Name: soilmoisture_id_seq; Type: SEQUENCE OWNED BY; Schema: goplant; Owner: postgres
--

ALTER SEQUENCE goplant.soilmoisture_id_seq OWNED BY goplant.soilmoisture.id;


--
-- Name: soiltemperature; Type: TABLE; Schema: goplant; Owner: postgres
--

CREATE TABLE goplant.soiltemperature (
    id integer NOT NULL,
    value integer NOT NULL,
    createdate date DEFAULT CURRENT_DATE,
    recordtime date
);


ALTER TABLE goplant.soiltemperature OWNER TO postgres;

--
-- Name: soiltemperature_id_seq; Type: SEQUENCE; Schema: goplant; Owner: postgres
--

CREATE SEQUENCE goplant.soiltemperature_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE goplant.soiltemperature_id_seq OWNER TO postgres;

--
-- Name: soiltemperature_id_seq; Type: SEQUENCE OWNED BY; Schema: goplant; Owner: postgres
--

ALTER SEQUENCE goplant.soiltemperature_id_seq OWNED BY goplant.soiltemperature.id;


--
-- Name: airtemperature id; Type: DEFAULT; Schema: goplant; Owner: postgres
--

ALTER TABLE ONLY goplant.airtemperature ALTER COLUMN id SET DEFAULT nextval('goplant.airtemperature_id_seq'::regclass);


--
-- Name: humidity id; Type: DEFAULT; Schema: goplant; Owner: postgres
--

ALTER TABLE ONLY goplant.humidity ALTER COLUMN id SET DEFAULT nextval('goplant.humidity_id_seq'::regclass);


--
-- Name: light id; Type: DEFAULT; Schema: goplant; Owner: postgres
--

ALTER TABLE ONLY goplant.light ALTER COLUMN id SET DEFAULT nextval('goplant.light_id_seq'::regclass);


--
-- Name: soilmoisture id; Type: DEFAULT; Schema: goplant; Owner: postgres
--

ALTER TABLE ONLY goplant.soilmoisture ALTER COLUMN id SET DEFAULT nextval('goplant.soilmoisture_id_seq'::regclass);


--
-- Name: soiltemperature id; Type: DEFAULT; Schema: goplant; Owner: postgres
--

ALTER TABLE ONLY goplant.soiltemperature ALTER COLUMN id SET DEFAULT nextval('goplant.soiltemperature_id_seq'::regclass);


--
-- Data for Name: airtemperature; Type: TABLE DATA; Schema: goplant; Owner: postgres
--

COPY goplant.airtemperature (id, value, createdate, recordtime) FROM stdin;
\.


--
-- Data for Name: humidity; Type: TABLE DATA; Schema: goplant; Owner: postgres
--

COPY goplant.humidity (id, value, createdate, recordtime) FROM stdin;
\.


--
-- Data for Name: light; Type: TABLE DATA; Schema: goplant; Owner: postgres
--

COPY goplant.light (id, value, createdate, recordtime) FROM stdin;
\.


--
-- Data for Name: soilmoisture; Type: TABLE DATA; Schema: goplant; Owner: postgres
--

COPY goplant.soilmoisture (id, value, createdate, recordtime) FROM stdin;
\.


--
-- Data for Name: soiltemperature; Type: TABLE DATA; Schema: goplant; Owner: postgres
--

COPY goplant.soiltemperature (id, value, createdate, recordtime) FROM stdin;
\.


--
-- Name: airtemperature_id_seq; Type: SEQUENCE SET; Schema: goplant; Owner: postgres
--

SELECT pg_catalog.setval('goplant.airtemperature_id_seq', 1, false);


--
-- Name: humidity_id_seq; Type: SEQUENCE SET; Schema: goplant; Owner: postgres
--

SELECT pg_catalog.setval('goplant.humidity_id_seq', 1, false);


--
-- Name: light_id_seq; Type: SEQUENCE SET; Schema: goplant; Owner: postgres
--

SELECT pg_catalog.setval('goplant.light_id_seq', 1, false);


--
-- Name: soilmoisture_id_seq; Type: SEQUENCE SET; Schema: goplant; Owner: postgres
--

SELECT pg_catalog.setval('goplant.soilmoisture_id_seq', 1, false);


--
-- Name: soiltemperature_id_seq; Type: SEQUENCE SET; Schema: goplant; Owner: postgres
--

SELECT pg_catalog.setval('goplant.soiltemperature_id_seq', 1, false);


--
-- Name: airtemperature airtemperature_pk; Type: CONSTRAINT; Schema: goplant; Owner: postgres
--

ALTER TABLE ONLY goplant.airtemperature
    ADD CONSTRAINT airtemperature_pk PRIMARY KEY (id);


--
-- Name: humidity humidity_pk; Type: CONSTRAINT; Schema: goplant; Owner: postgres
--

ALTER TABLE ONLY goplant.humidity
    ADD CONSTRAINT humidity_pk PRIMARY KEY (id);


--
-- Name: light light_pk; Type: CONSTRAINT; Schema: goplant; Owner: postgres
--

ALTER TABLE ONLY goplant.light
    ADD CONSTRAINT light_pk PRIMARY KEY (id);


--
-- Name: soilmoisture soilmoisture_pk; Type: CONSTRAINT; Schema: goplant; Owner: postgres
--

ALTER TABLE ONLY goplant.soilmoisture
    ADD CONSTRAINT soilmoisture_pk PRIMARY KEY (id);


--
-- Name: soiltemperature soiltemperature_pk; Type: CONSTRAINT; Schema: goplant; Owner: postgres
--

ALTER TABLE ONLY goplant.soiltemperature
    ADD CONSTRAINT soiltemperature_pk PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2 (Debian 13.2-1.pgdg100+1)
-- Dumped by pg_dump version 13.2 (Debian 13.2-1.pgdg100+1)

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

DROP DATABASE postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE postgres OWNER TO postgres;

\connect postgres

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
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

