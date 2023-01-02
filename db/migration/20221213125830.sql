CREATE TABLE public.cart_details (
    cd_id integer NOT NULL,
    cart_id integer,
    product_id integer,
    product_price numeric,
    quantity integer,
    delivery_price numeric
);

CREATE SEQUENCE public.cart_details_cd_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cart_details_cd_id_seq OWNED BY public.cart_details.cd_id;


CREATE TABLE public.carts (
    cart_id integer NOT NULL,
    user_id integer,
    price numeric,
    delivery_price numeric,
    created_on timestamp with time zone DEFAULT now(),
    updated_on timestamp with time zone DEFAULT now(),
    total numeric
);

CREATE SEQUENCE public.carts_cart_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.carts_cart_id_seq OWNED BY public.carts.cart_id;



CREATE TABLE public.categories (
    category_id integer NOT NULL,
    category_name text NOT NULL,
    created_on timestamp with time zone DEFAULT now(),
    image_url text,
    status text
);

CREATE SEQUENCE public.categories_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.categories_category_id_seq OWNED BY public.categories.category_id;

CREATE TABLE public.countries (
    country_id integer NOT NULL,
    country_name character varying(20),
    currency character varying(10),
    currency_symbol character varying(5)
);

CREATE SEQUENCE public.countries_country_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.countries_country_id_seq OWNED BY public.countries.country_id;


CREATE TABLE public.order (
    order_id character varying(20) NOT NULL,
    user_id integer NOT NULL,
    price numeric NOT NULL,
    delivery_price numeric NOT NULL,
    total numeric NOT NULL,
    created_on timestamp with time zone DEFAULT now(),
    email character varying(50) NOT NULL,
    address text NOT NULL
);
CREATE TABLE public.order_details (
    od_id integer NOT NULL,
    order_id character varying(20),
    product_id integer,
    product_price numeric,
    quantity integer,
    delivery_price numeric
);

CREATE SEQUENCE public.order_details_od_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.order_details_od_id_seq OWNED BY public.order_details.od_id;

CREATE TABLE public.products (
    product_id integer NOT NULL,
    category_id integer,
    product_name character varying NOT NULL,
    product_image character varying,
    quantity integer,
    created_on timestamp with time zone DEFAULT now(),
    updated_on timestamp with time zone DEFAULT now(),
    status text DEFAULT 'active'::text,
    price numeric NOT NULL,
    delivery_price numeric DEFAULT 0,
    product_desc character varying(500) DEFAULT 'No description'::character varying,
    gender character varying(10) DEFAULT 'Not specified'::character varying,
    country_id integer DEFAULT 1
);

CREATE SEQUENCE public.products_product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

--
-- TOC entry 3156 (class 0 OID 0)
-- Dependencies: 206
-- Name: products_product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.products_product_id_seq OWNED BY public.products.product_id;


--
-- TOC entry 203 (class 1259 OID 16536)
-- Name: tokens; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tokens (
    token_id integer NOT NULL,
    user_id integer,
    token character varying NOT NULL,
    created_on timestamp with time zone DEFAULT now(),
    last_access timestamp with time zone DEFAULT now()
);

--
-- TOC entry 202 (class 1259 OID 16534)
-- Name: tokens_token_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tokens_token_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3157 (class 0 OID 0)
-- Dependencies: 202
-- Name: tokens_token_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tokens_token_id_seq OWNED BY public.tokens.token_id;


--
-- TOC entry 201 (class 1259 OID 16520)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    user_id integer NOT NULL,
    first_name text NOT NULL,
    last_name text,
    email character varying NOT NULL,
    phone bigint,
    password character varying NOT NULL,
    created_on timestamp with time zone DEFAULT now(),
    updated_on timestamp with time zone DEFAULT now(),
    status character varying(10) DEFAULT 'active'::character varying,
    access varchar(20) default 'user'
);


CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;

ALTER TABLE ONLY public.cart_details ALTER COLUMN cd_id SET DEFAULT nextval('public.cart_details_cd_id_seq'::regclass);


--
-- TOC entry 2980 (class 2604 OID 16591)
-- Name: carts cart_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts ALTER COLUMN cart_id SET DEFAULT nextval('public.carts_cart_id_seq'::regclass);


--
-- TOC entry 2970 (class 2604 OID 16557)
-- Name: categories category_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN category_id SET DEFAULT nextval('public.categories_category_id_seq'::regclass);


--
-- TOC entry 2986 (class 2604 OID 25517)
-- Name: countries country_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.countries ALTER COLUMN country_id SET DEFAULT nextval('public.countries_country_id_seq'::regclass);


--
-- TOC entry 2985 (class 2604 OID 16762)
-- Name: order_details od_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_details ALTER COLUMN od_id SET DEFAULT nextval('public.order_details_od_id_seq'::regclass);


--
-- TOC entry 2972 (class 2604 OID 16569)
-- Name: products product_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN product_id SET DEFAULT nextval('public.products_product_id_seq'::regclass);


--
-- TOC entry 2967 (class 2604 OID 16539)
-- Name: tokens token_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tokens ALTER COLUMN token_id SET DEFAULT nextval('public.tokens_token_id_seq'::regclass);


--
-- TOC entry 2963 (class 2604 OID 16523)
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- TOC entry 3000 (class 2606 OID 16608)
-- Name: cart_details cart_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_details
    ADD CONSTRAINT cart_details_pkey PRIMARY KEY (cd_id);


--
-- TOC entry 2998 (class 2606 OID 16595)
-- Name: carts carts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_pkey PRIMARY KEY (cart_id);


--
-- TOC entry 2994 (class 2606 OID 16563)
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (category_id);


--
-- TOC entry 3006 (class 2606 OID 25521)
-- Name: countries countries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.countries
    ADD CONSTRAINT countries_pkey PRIMARY KEY (country_id);


--
-- TOC entry 3004 (class 2606 OID 16764)
-- Name: order_details order_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_details
    ADD CONSTRAINT order_details_pkey PRIMARY KEY (od_id);


--
-- TOC entry 3002 (class 2606 OID 16751)
-- Name: order order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order
    ADD CONSTRAINT order_pkey PRIMARY KEY (order_id);


--
-- TOC entry 2996 (class 2606 OID 16580)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (product_id);


--
-- TOC entry 2992 (class 2606 OID 16546)
-- Name: tokens tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (token_id);


--
-- TOC entry 2988 (class 2606 OID 16533)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 2990 (class 2606 OID 16531)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- TOC entry 3011 (class 2606 OID 16609)
-- Name: cart_details fk_cart; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_details
    ADD CONSTRAINT fk_cart FOREIGN KEY (cart_id) REFERENCES public.carts(cart_id) ON DELETE CASCADE;


--
-- TOC entry 3008 (class 2606 OID 16581)
-- Name: products fk_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES public.categories(category_id) ON DELETE CASCADE;


--
-- TOC entry 3009 (class 2606 OID 25522)
-- Name: products fk_country; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_country FOREIGN KEY (country_id) REFERENCES public.countries(country_id);


--
-- TOC entry 3014 (class 2606 OID 16765)
-- Name: order_details fk_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_details
    ADD CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES public.order(order_id) ON DELETE CASCADE;


--
-- TOC entry 3012 (class 2606 OID 16614)
-- Name: cart_details fk_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_details
    ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES public.products(product_id) ON DELETE CASCADE;


--
-- TOC entry 3015 (class 2606 OID 16770)
-- Name: order_details fk_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_details
    ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES public.products(product_id) ON DELETE CASCADE;


--
-- TOC entry 3013 (class 2606 OID 16752)
-- Name: order fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;


--
-- TOC entry 3010 (class 2606 OID 16596)
-- Name: carts fk_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;


--
-- TOC entry 3007 (class 2606 OID 16547)
-- Name: tokens fk_usertoken; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT fk_usertoken FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;


-- Completed on 2021-10-28 21:51:55 IST

--
-- PostgreSQL database dump complete
--
