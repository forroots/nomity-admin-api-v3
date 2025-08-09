--
-- PostgreSQL database dump
--

-- Dumped from database version 13.18
-- Dumped by pg_dump version 13.18

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

--
-- Name: fiber_storage; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.fiber_storage (
    k character varying(64) DEFAULT ''::character varying NOT NULL,
    v bytea NOT NULL,
    e bigint DEFAULT '0'::bigint NOT NULL
);


ALTER TABLE public.fiber_storage OWNER TO nomity_dev;

--
-- Name: fiber_storage_admin; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.fiber_storage_admin (
    k character varying(64) DEFAULT ''::character varying NOT NULL,
    v bytea NOT NULL,
    e bigint DEFAULT 0 NOT NULL
);


ALTER TABLE public.fiber_storage_admin OWNER TO nomity_dev;

--
-- Name: t_admin_action_logs; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_admin_action_logs (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    t_admin_user_id bigint NOT NULL,
    action_type text NOT NULL,
    message text,
    details jsonb,
    request_id text NOT NULL,
    ip_address text NOT NULL
);


ALTER TABLE public.t_admin_action_logs OWNER TO nomity_dev;

--
-- Name: COLUMN t_admin_action_logs.message; Type: COMMENT; Schema: public; Owner: nomity_dev
--

COMMENT ON COLUMN public.t_admin_action_logs.message IS '自由記述のメッセージ';


--
-- Name: COLUMN t_admin_action_logs.details; Type: COMMENT; Schema: public; Owner: nomity_dev
--

COMMENT ON COLUMN public.t_admin_action_logs.details IS '詳細なパラメータや差分など、任意の構造データ';


--
-- Name: t_admin_action_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_admin_action_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_admin_action_logs_id_seq OWNER TO nomity_dev;

--
-- Name: t_admin_action_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_admin_action_logs_id_seq OWNED BY public.t_admin_action_logs.id;


--
-- Name: t_admin_users; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_admin_users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    email text NOT NULL,
    role text NOT NULL,
    active boolean NOT NULL,
    email_confirmed_at timestamp with time zone
);


ALTER TABLE public.t_admin_users OWNER TO nomity_dev;

--
-- Name: t_admin_users_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_admin_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_admin_users_id_seq OWNER TO nomity_dev;

--
-- Name: t_admin_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_admin_users_id_seq OWNED BY public.t_admin_users.id;


--
-- Name: t_android_tmp_users; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_android_tmp_users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    email text NOT NULL
);


ALTER TABLE public.t_android_tmp_users OWNER TO nomity_dev;

--
-- Name: t_android_tmp_users_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_android_tmp_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_android_tmp_users_id_seq OWNER TO nomity_dev;

--
-- Name: t_android_tmp_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_android_tmp_users_id_seq OWNED BY public.t_android_tmp_users.id;


--
-- Name: t_article_tags; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_article_tags (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    t_article_id bigint NOT NULL,
    tag text NOT NULL,
    display_order bigint NOT NULL
);


ALTER TABLE public.t_article_tags OWNER TO nomity_dev;

--
-- Name: t_article_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_article_tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_article_tags_id_seq OWNER TO nomity_dev;

--
-- Name: t_article_tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_article_tags_id_seq OWNED BY public.t_article_tags.id;


--
-- Name: t_articles; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_articles (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text NOT NULL,
    thumbnail_url text NOT NULL,
    content_url text NOT NULL,
    release_date timestamp with time zone NOT NULL
);


ALTER TABLE public.t_articles OWNER TO nomity_dev;

--
-- Name: t_articles_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_articles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_articles_id_seq OWNER TO nomity_dev;

--
-- Name: t_articles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_articles_id_seq OWNED BY public.t_articles.id;


--
-- Name: t_favorite_alcohols; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_favorite_alcohols (
    t_user_id bigint NOT NULL,
    alcohol_type_cd bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.t_favorite_alcohols OWNER TO nomity_dev;

--
-- Name: t_kensakit_receipt_import_items; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_kensakit_receipt_import_items (
    id bigint NOT NULL,
    created_by_admin_user_id bigint NOT NULL,
    admin_memo text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    t_kensakit_receipt_import_id bigint NOT NULL,
    serial_number text NOT NULL,
    status text NOT NULL,
    update_kensakit_status_batch_at timestamp with time zone,
    update_kensakit_status_batch_success boolean,
    update_kensakit_status_batch_error_message text,
    send_mail_batch_at timestamp with time zone,
    send_mail_batch_success boolean,
    send_mail_batch_error_message text,
    send_mail_to text
);


ALTER TABLE public.t_kensakit_receipt_import_items OWNER TO nomity_dev;

--
-- Name: t_kensakit_receipt_import_items_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_kensakit_receipt_import_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_kensakit_receipt_import_items_id_seq OWNER TO nomity_dev;

--
-- Name: t_kensakit_receipt_import_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_kensakit_receipt_import_items_id_seq OWNED BY public.t_kensakit_receipt_import_items.id;


--
-- Name: t_kensakit_receipt_imports; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_kensakit_receipt_imports (
    id bigint NOT NULL,
    created_by_admin_user_id bigint NOT NULL,
    admin_memo text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    status text NOT NULL,
    imported_at timestamp with time zone NOT NULL,
    imported_file_name text NOT NULL,
    update_kensakit_status_batch_at timestamp with time zone,
    update_kensakit_status_batch_success boolean,
    update_kensakit_status_batch_error_message text,
    send_mail_batch_at timestamp with time zone,
    send_mail_batch_success boolean,
    send_mail_batch_error_message text
);


ALTER TABLE public.t_kensakit_receipt_imports OWNER TO nomity_dev;

--
-- Name: t_kensakit_receipt_imports_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_kensakit_receipt_imports_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_kensakit_receipt_imports_id_seq OWNER TO nomity_dev;

--
-- Name: t_kensakit_receipt_imports_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_kensakit_receipt_imports_id_seq OWNED BY public.t_kensakit_receipt_imports.id;


--
-- Name: t_kensakit_status_histories; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_kensakit_status_histories (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    t_kensakit_id bigint NOT NULL,
    status_cd text NOT NULL,
    executed_time timestamp with time zone NOT NULL
);


ALTER TABLE public.t_kensakit_status_histories OWNER TO nomity_dev;

--
-- Name: t_kensakit_status_histories_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_kensakit_status_histories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_kensakit_status_histories_id_seq OWNER TO nomity_dev;

--
-- Name: t_kensakit_status_histories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_kensakit_status_histories_id_seq OWNED BY public.t_kensakit_status_histories.id;


--
-- Name: t_kensakits; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_kensakits (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    t_user_id bigint,
    serial_number text NOT NULL,
    auth_key1 text NOT NULL,
    auth_key2 text NOT NULL,
    status_for_user bigint NOT NULL,
    test_result bigint
);


ALTER TABLE public.t_kensakits OWNER TO nomity_dev;

--
-- Name: t_kensakits_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_kensakits_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_kensakits_id_seq OWNER TO nomity_dev;

--
-- Name: t_kensakits_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_kensakits_id_seq OWNED BY public.t_kensakits.id;


--
-- Name: t_restaurants; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_restaurants (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    prefecture text NOT NULL,
    city text NOT NULL
);


ALTER TABLE public.t_restaurants OWNER TO nomity_dev;

--
-- Name: t_restaurants_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_restaurants_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_restaurants_id_seq OWNER TO nomity_dev;

--
-- Name: t_restaurants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_restaurants_id_seq OWNED BY public.t_restaurants.id;


--
-- Name: t_sanwa_questionaires; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_sanwa_questionaires (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    t_kensakit_id bigint NOT NULL,
    is_qs1_done boolean,
    is_qs2_done boolean,
    is_qs3_1_done boolean,
    is_qs3_2_done boolean,
    is_qs3_3_done boolean,
    qn_1_1 bigint,
    qn_1_2 text,
    qn_1_3 text,
    qn_1_4 bigint,
    qn_1_5 bigint,
    qn_1_6 bigint,
    qn_1_7 bigint,
    qn_1_8 bigint,
    qn_1_9 bigint,
    qs_1_1 bigint,
    qs_1_2 bigint,
    qs_1_3 bigint,
    qs_1_4 bigint,
    qs_1_5 bigint,
    qs_1_6 bigint,
    qs_1_7 bigint,
    qs_1_8 bigint,
    qs_1_9_1 bigint,
    qs_1_9_2 bigint,
    qs_1_9_3 bigint,
    qs_1_9_4 bigint,
    qs_1_9_5 bigint,
    qs_1_9_6 bigint,
    qs_1_9_7 bigint,
    qs_1_9_8 bigint,
    qs_1_9_9 bigint,
    qs_1_9_10 bigint,
    qs_1_9_11 bigint,
    qs_1_9_12 bigint,
    qs_1_9_13 bigint,
    qs_1_9_14 bigint,
    qs_1_9_15 bigint,
    qs_1_9_16 bigint,
    qs_1_9_17 bigint,
    qs_1_9_18 bigint,
    qs_1_9_19 bigint,
    qs_1_9_20 bigint,
    qs_1_9_21 bigint,
    qs_1_10_1 bigint,
    qs_1_10_2 bigint,
    qs_1_10_3 bigint,
    qs_1_10_4 bigint,
    qs_1_10_5 bigint,
    qs_1_10_6 bigint,
    qs_1_10_7 bigint,
    qs_1_10_8 bigint,
    qs_1_10_9 bigint,
    qs_1_10_10 bigint,
    qs_1_10_11 bigint,
    qs_1_10_12 bigint,
    qs_1_10_13 bigint,
    qs_1_10_14 bigint,
    qs_1_10_15 bigint,
    qs_1_10_16 bigint,
    qs_1_10_17 bigint,
    qs_1_10_18 bigint,
    qs_1_10_19 bigint,
    qs_1_10_20 bigint,
    qs_1_10_21 bigint,
    qs_1_11 bigint,
    qs_1_12 bigint,
    qs_1_13 bigint,
    qs_1_14 bigint,
    qs_1_15 bigint,
    qs_1_16 bigint,
    qs_1_17 bigint,
    qs_1_18_1 bigint,
    qs_1_18_2 bigint,
    qs_1_18_3 bigint,
    qs_1_18_4 bigint,
    qs_1_18_5 bigint,
    qs_1_18_6 bigint,
    qs_1_18_7 bigint,
    qs_1_18_8 bigint,
    qs_1_18_9 bigint,
    qs_1_18_10 bigint,
    qs_1_18_11 bigint,
    qs_1_18_12 bigint,
    qs_1_18_13 bigint,
    qs_1_18_14 bigint,
    qs_1_18_15 bigint,
    qs_1_18_16 bigint,
    qs_1_18_17 bigint,
    qs_1_18_18 bigint,
    qs_1_18_19 bigint,
    qs_1_18_20 bigint,
    qs_1_18_21 bigint,
    qs_1_19 text,
    qs_1_20 text,
    qs_2_1 bigint,
    qs_2_2 bigint,
    qs_2_3 bigint,
    qs_2_4 bigint,
    qs_2_5 bigint,
    qs_2_6 bigint,
    qs_2_7 bigint,
    qs_2_8 bigint,
    qs_2_9 bigint,
    qs_2_10 bigint,
    qs_2_11 bigint,
    qs_2_12 bigint,
    qs_2_13 bigint,
    qs_2_14 bigint,
    qs_2_15 bigint,
    qs_2_16 bigint,
    qs_2_17 bigint,
    qs_2_18 bigint,
    qs_2_19 bigint,
    qs_2_20 bigint,
    qs_2_21 bigint,
    qs_2_22 bigint,
    qs_2_23 bigint,
    qs_2_24 bigint,
    qs_2_25 bigint,
    qs_2_26 bigint
);


ALTER TABLE public.t_sanwa_questionaires OWNER TO nomity_dev;

--
-- Name: t_sanwa_questionaires_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_sanwa_questionaires_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_sanwa_questionaires_id_seq OWNER TO nomity_dev;

--
-- Name: t_sanwa_questionaires_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_sanwa_questionaires_id_seq OWNED BY public.t_sanwa_questionaires.id;


--
-- Name: t_taishitsu_questionaires; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_taishitsu_questionaires (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    t_user_id bigint NOT NULL,
    audit_c_frequency_cd bigint NOT NULL,
    audit_c_quantity_cd bigint NOT NULL,
    audit_c_excessive_alcohol_frequency_cd bigint NOT NULL,
    self_knowledge_about_tolerance_cd bigint NOT NULL,
    self_knowledge_about_type_cd bigint NOT NULL
);


ALTER TABLE public.t_taishitsu_questionaires OWNER TO nomity_dev;

--
-- Name: t_taishitsu_questionaires_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_taishitsu_questionaires_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_taishitsu_questionaires_id_seq OWNER TO nomity_dev;

--
-- Name: t_taishitsu_questionaires_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_taishitsu_questionaires_id_seq OWNED BY public.t_taishitsu_questionaires.id;


--
-- Name: t_user_action_logs; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_user_action_logs (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    t_user_id bigint NOT NULL,
    action_type text NOT NULL,
    details jsonb,
    message text NOT NULL
);


ALTER TABLE public.t_user_action_logs OWNER TO nomity_dev;

--
-- Name: t_user_action_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_user_action_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_user_action_logs_id_seq OWNER TO nomity_dev;

--
-- Name: t_user_action_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_user_action_logs_id_seq OWNED BY public.t_user_action_logs.id;


--
-- Name: t_user_details; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_user_details (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    t_user_id bigint NOT NULL,
    sex_cd bigint NOT NULL,
    birth_date date,
    residence_cd bigint NOT NULL,
    birthplace_cd bigint NOT NULL,
    type_of_work_cd text NOT NULL,
    type_of_occupation_cd text NOT NULL,
    alcohol_taishitsu_type_cd bigint NOT NULL
);


ALTER TABLE public.t_user_details OWNER TO nomity_dev;

--
-- Name: t_user_details_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_user_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_user_details_id_seq OWNER TO nomity_dev;

--
-- Name: t_user_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_user_details_id_seq OWNED BY public.t_user_details.id;


--
-- Name: t_users; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.t_users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    parent_user_id bigint,
    name text NOT NULL,
    email text,
    password_hash text NOT NULL,
    serching_id text,
    allow_expose boolean DEFAULT false NOT NULL,
    role_cd text NOT NULL
);


ALTER TABLE public.t_users OWNER TO nomity_dev;

--
-- Name: t_users_id_seq; Type: SEQUENCE; Schema: public; Owner: nomity_dev
--

CREATE SEQUENCE public.t_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_users_id_seq OWNER TO nomity_dev;

--
-- Name: t_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nomity_dev
--

ALTER SEQUENCE public.t_users_id_seq OWNED BY public.t_users.id;


--
-- Name: w_admin_device_sessions; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.w_admin_device_sessions (
    session_id text NOT NULL,
    created_at timestamp with time zone,
    t_admin_user_id bigint,
    email text,
    otp text,
    otp_expires_at timestamp with time zone
);


ALTER TABLE public.w_admin_device_sessions OWNER TO nomity_dev;

--
-- Name: w_email_authenticatings; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.w_email_authenticatings (
    access_token text NOT NULL,
    email text NOT NULL,
    auth_code text NOT NULL,
    authenticated boolean NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.w_email_authenticatings OWNER TO nomity_dev;

--
-- Name: w_sessions; Type: TABLE; Schema: public; Owner: nomity_dev
--

CREATE TABLE public.w_sessions (
    access_token text NOT NULL,
    refresh_token text NOT NULL,
    t_user_id bigint NOT NULL,
    access_token_expired_in timestamp with time zone NOT NULL,
    refresh_token_expired_in timestamp with time zone NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    device_type text DEFAULT 'mobile'::text NOT NULL
);


ALTER TABLE public.w_sessions OWNER TO nomity_dev;

--
-- Name: t_admin_action_logs id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_admin_action_logs ALTER COLUMN id SET DEFAULT nextval('public.t_admin_action_logs_id_seq'::regclass);


--
-- Name: t_admin_users id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_admin_users ALTER COLUMN id SET DEFAULT nextval('public.t_admin_users_id_seq'::regclass);


--
-- Name: t_android_tmp_users id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_android_tmp_users ALTER COLUMN id SET DEFAULT nextval('public.t_android_tmp_users_id_seq'::regclass);


--
-- Name: t_article_tags id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_article_tags ALTER COLUMN id SET DEFAULT nextval('public.t_article_tags_id_seq'::regclass);


--
-- Name: t_articles id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_articles ALTER COLUMN id SET DEFAULT nextval('public.t_articles_id_seq'::regclass);


--
-- Name: t_kensakit_receipt_import_items id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakit_receipt_import_items ALTER COLUMN id SET DEFAULT nextval('public.t_kensakit_receipt_import_items_id_seq'::regclass);


--
-- Name: t_kensakit_receipt_imports id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakit_receipt_imports ALTER COLUMN id SET DEFAULT nextval('public.t_kensakit_receipt_imports_id_seq'::regclass);


--
-- Name: t_kensakit_status_histories id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakit_status_histories ALTER COLUMN id SET DEFAULT nextval('public.t_kensakit_status_histories_id_seq'::regclass);


--
-- Name: t_kensakits id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakits ALTER COLUMN id SET DEFAULT nextval('public.t_kensakits_id_seq'::regclass);


--
-- Name: t_restaurants id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_restaurants ALTER COLUMN id SET DEFAULT nextval('public.t_restaurants_id_seq'::regclass);


--
-- Name: t_sanwa_questionaires id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_sanwa_questionaires ALTER COLUMN id SET DEFAULT nextval('public.t_sanwa_questionaires_id_seq'::regclass);


--
-- Name: t_taishitsu_questionaires id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_taishitsu_questionaires ALTER COLUMN id SET DEFAULT nextval('public.t_taishitsu_questionaires_id_seq'::regclass);


--
-- Name: t_user_action_logs id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_user_action_logs ALTER COLUMN id SET DEFAULT nextval('public.t_user_action_logs_id_seq'::regclass);


--
-- Name: t_user_details id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_user_details ALTER COLUMN id SET DEFAULT nextval('public.t_user_details_id_seq'::regclass);


--
-- Name: t_users id; Type: DEFAULT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_users ALTER COLUMN id SET DEFAULT nextval('public.t_users_id_seq'::regclass);


--
-- Name: fiber_storage_admin fiber_storage_admin_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.fiber_storage_admin
    ADD CONSTRAINT fiber_storage_admin_pkey PRIMARY KEY (k);


--
-- Name: fiber_storage fiber_storage_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.fiber_storage
    ADD CONSTRAINT fiber_storage_pkey PRIMARY KEY (k);


--
-- Name: t_admin_action_logs t_admin_action_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_admin_action_logs
    ADD CONSTRAINT t_admin_action_logs_pkey PRIMARY KEY (id);


--
-- Name: t_admin_users t_admin_users_email_key; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_admin_users
    ADD CONSTRAINT t_admin_users_email_key UNIQUE (email);


--
-- Name: t_admin_users t_admin_users_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_admin_users
    ADD CONSTRAINT t_admin_users_pkey PRIMARY KEY (id);


--
-- Name: t_android_tmp_users t_android_tmp_users_email_key; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_android_tmp_users
    ADD CONSTRAINT t_android_tmp_users_email_key UNIQUE (email);


--
-- Name: t_android_tmp_users t_android_tmp_users_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_android_tmp_users
    ADD CONSTRAINT t_android_tmp_users_pkey PRIMARY KEY (id);


--
-- Name: t_article_tags t_article_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_article_tags
    ADD CONSTRAINT t_article_tags_pkey PRIMARY KEY (id);


--
-- Name: t_articles t_articles_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_articles
    ADD CONSTRAINT t_articles_pkey PRIMARY KEY (id);


--
-- Name: t_favorite_alcohols t_favorite_alcohols_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_favorite_alcohols
    ADD CONSTRAINT t_favorite_alcohols_pkey PRIMARY KEY (t_user_id, alcohol_type_cd);


--
-- Name: t_kensakit_receipt_import_items t_kensakit_receipt_import_items_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakit_receipt_import_items
    ADD CONSTRAINT t_kensakit_receipt_import_items_pkey PRIMARY KEY (id);


--
-- Name: t_kensakit_receipt_imports t_kensakit_receipt_imports_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakit_receipt_imports
    ADD CONSTRAINT t_kensakit_receipt_imports_pkey PRIMARY KEY (id);


--
-- Name: t_kensakit_status_histories t_kensakit_status_histories_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakit_status_histories
    ADD CONSTRAINT t_kensakit_status_histories_pkey PRIMARY KEY (id);


--
-- Name: t_kensakits t_kensakits_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakits
    ADD CONSTRAINT t_kensakits_pkey PRIMARY KEY (id);


--
-- Name: t_kensakits t_kensakits_serial_number_key; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakits
    ADD CONSTRAINT t_kensakits_serial_number_key UNIQUE (serial_number);


--
-- Name: t_restaurants t_restaurants_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_restaurants
    ADD CONSTRAINT t_restaurants_pkey PRIMARY KEY (id);


--
-- Name: t_sanwa_questionaires t_sanwa_questionaires_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_sanwa_questionaires
    ADD CONSTRAINT t_sanwa_questionaires_pkey PRIMARY KEY (id);


--
-- Name: t_sanwa_questionaires t_sanwa_questionaires_t_kensakit_id_key; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_sanwa_questionaires
    ADD CONSTRAINT t_sanwa_questionaires_t_kensakit_id_key UNIQUE (t_kensakit_id);


--
-- Name: t_taishitsu_questionaires t_taishitsu_questionaires_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_taishitsu_questionaires
    ADD CONSTRAINT t_taishitsu_questionaires_pkey PRIMARY KEY (id);


--
-- Name: t_taishitsu_questionaires t_taishitsu_questionaires_t_user_id_key; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_taishitsu_questionaires
    ADD CONSTRAINT t_taishitsu_questionaires_t_user_id_key UNIQUE (t_user_id);


--
-- Name: t_user_action_logs t_user_action_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_user_action_logs
    ADD CONSTRAINT t_user_action_logs_pkey PRIMARY KEY (id);


--
-- Name: t_user_details t_user_details_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_user_details
    ADD CONSTRAINT t_user_details_pkey PRIMARY KEY (id);


--
-- Name: t_user_details t_user_details_t_user_id_key; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_user_details
    ADD CONSTRAINT t_user_details_t_user_id_key UNIQUE (t_user_id);


--
-- Name: t_users t_users_email_key; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_users
    ADD CONSTRAINT t_users_email_key UNIQUE (email);


--
-- Name: t_users t_users_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_users
    ADD CONSTRAINT t_users_pkey PRIMARY KEY (id);


--
-- Name: t_users t_users_serching_id_key; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_users
    ADD CONSTRAINT t_users_serching_id_key UNIQUE (serching_id);


--
-- Name: w_admin_device_sessions w_admin_device_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.w_admin_device_sessions
    ADD CONSTRAINT w_admin_device_sessions_pkey PRIMARY KEY (session_id);


--
-- Name: w_email_authenticatings w_email_authenticatings_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.w_email_authenticatings
    ADD CONSTRAINT w_email_authenticatings_pkey PRIMARY KEY (access_token);


--
-- Name: w_sessions w_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.w_sessions
    ADD CONSTRAINT w_sessions_pkey PRIMARY KEY (access_token);


--
-- Name: e; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX e ON public.fiber_storage USING btree (e);


--
-- Name: idx_t_admin_action_logs_created_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_admin_action_logs_created_at ON public.t_admin_action_logs USING btree (created_at);


--
-- Name: idx_t_admin_users_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_admin_users_deleted_at ON public.t_admin_users USING btree (deleted_at);


--
-- Name: idx_t_android_tmp_users_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_android_tmp_users_deleted_at ON public.t_android_tmp_users USING btree (deleted_at);


--
-- Name: idx_t_article_tags_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_article_tags_deleted_at ON public.t_article_tags USING btree (deleted_at);


--
-- Name: idx_t_articles_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_articles_deleted_at ON public.t_articles USING btree (deleted_at);


--
-- Name: idx_t_kensakit_status_histories_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_kensakit_status_histories_deleted_at ON public.t_kensakit_status_histories USING btree (deleted_at);


--
-- Name: idx_t_kensakits_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_kensakits_deleted_at ON public.t_kensakits USING btree (deleted_at);


--
-- Name: idx_t_restaurants_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_restaurants_deleted_at ON public.t_restaurants USING btree (deleted_at);


--
-- Name: idx_t_sanwa_questionaires_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_sanwa_questionaires_deleted_at ON public.t_sanwa_questionaires USING btree (deleted_at);


--
-- Name: idx_t_taishitsu_questionaires_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_taishitsu_questionaires_deleted_at ON public.t_taishitsu_questionaires USING btree (deleted_at);


--
-- Name: idx_t_user_action_logs_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_user_action_logs_deleted_at ON public.t_user_action_logs USING btree (deleted_at);


--
-- Name: idx_t_user_details_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_user_details_deleted_at ON public.t_user_details USING btree (deleted_at);


--
-- Name: idx_t_users_deleted_at; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE INDEX idx_t_users_deleted_at ON public.t_users USING btree (deleted_at);


--
-- Name: idx_unique_t_article_tags_01; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE UNIQUE INDEX idx_unique_t_article_tags_01 ON public.t_article_tags USING btree (t_article_id, tag);


--
-- Name: idx_unique_t_kensakits_01; Type: INDEX; Schema: public; Owner: nomity_dev
--

CREATE UNIQUE INDEX idx_unique_t_kensakits_01 ON public.t_kensakits USING btree (auth_key1, auth_key2);


--
-- Name: t_kensakit_receipt_import_items fk__t_kensakit_receipt_import_items__t_kensakit_receipt_import_; Type: FK CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakit_receipt_import_items
    ADD CONSTRAINT fk__t_kensakit_receipt_import_items__t_kensakit_receipt_import_ FOREIGN KEY (t_kensakit_receipt_import_id) REFERENCES public.t_kensakit_receipt_imports(id);


--
-- Name: t_admin_action_logs fk_t_admin_action_logs_owner; Type: FK CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_admin_action_logs
    ADD CONSTRAINT fk_t_admin_action_logs_owner FOREIGN KEY (t_admin_user_id) REFERENCES public.t_admin_users(id);


--
-- Name: w_admin_device_sessions fk_t_admin_users_w_admin_device_sessions; Type: FK CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.w_admin_device_sessions
    ADD CONSTRAINT fk_t_admin_users_w_admin_device_sessions FOREIGN KEY (t_admin_user_id) REFERENCES public.t_admin_users(id);


--
-- Name: t_article_tags fk_t_articles_tags; Type: FK CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_article_tags
    ADD CONSTRAINT fk_t_articles_tags FOREIGN KEY (t_article_id) REFERENCES public.t_articles(id);


--
-- Name: t_kensakit_status_histories fk_t_kensakits_t_status_histories; Type: FK CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakit_status_histories
    ADD CONSTRAINT fk_t_kensakits_t_status_histories FOREIGN KEY (t_kensakit_id) REFERENCES public.t_kensakits(id);


--
-- Name: t_favorite_alcohols fk_t_users_t_favorite_alcohols; Type: FK CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_favorite_alcohols
    ADD CONSTRAINT fk_t_users_t_favorite_alcohols FOREIGN KEY (t_user_id) REFERENCES public.t_users(id);


--
-- Name: t_kensakits fk_t_users_t_kensakits; Type: FK CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_kensakits
    ADD CONSTRAINT fk_t_users_t_kensakits FOREIGN KEY (t_user_id) REFERENCES public.t_users(id);


--
-- Name: t_taishitsu_questionaires fk_t_users_t_taishitsu_questionaire; Type: FK CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_taishitsu_questionaires
    ADD CONSTRAINT fk_t_users_t_taishitsu_questionaire FOREIGN KEY (t_user_id) REFERENCES public.t_users(id);


--
-- Name: t_user_details fk_t_users_t_user_detail; Type: FK CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_user_details
    ADD CONSTRAINT fk_t_users_t_user_detail FOREIGN KEY (t_user_id) REFERENCES public.t_users(id);


--
-- Name: t_users fk_t_users_t_users; Type: FK CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.t_users
    ADD CONSTRAINT fk_t_users_t_users FOREIGN KEY (parent_user_id) REFERENCES public.t_users(id);


--
-- Name: w_sessions fk_t_users_w_session; Type: FK CONSTRAINT; Schema: public; Owner: nomity_dev
--

ALTER TABLE ONLY public.w_sessions
    ADD CONSTRAINT fk_t_users_w_session FOREIGN KEY (t_user_id) REFERENCES public.t_users(id);


--
-- PostgreSQL database dump complete
--

