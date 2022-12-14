PGDMP         $                z            postgres    14.4    14.4     ?           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            ?           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            ?           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            ?           1262    13754    postgres    DATABASE     l   CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
    DROP DATABASE postgres;
                postgres    false            ?           0    0    DATABASE postgres    COMMENT     N   COMMENT ON DATABASE postgres IS 'default administrative connection database';
                   postgres    false    3322                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                postgres    false            ?           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   postgres    false    4            ?            1259    17048    book    TABLE     ?   CREATE TABLE public.book (
    id integer NOT NULL,
    title character varying NOT NULL,
    genre character varying,
    author character varying NOT NULL
);
    DROP TABLE public.book;
       public         heap    postgres    false    4            ?            1259    17047    book_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.book_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.book_id_seq;
       public          postgres    false    4    211            ?           0    0    book_id_seq    SEQUENCE OWNED BY     ;   ALTER SEQUENCE public.book_id_seq OWNED BY public.book.id;
          public          postgres    false    210            ?            1259    17057    customer    TABLE     ?   CREATE TABLE public.customer (
    id integer NOT NULL,
    firstname character varying NOT NULL,
    lastname character varying NOT NULL,
    age integer
);
    DROP TABLE public.customer;
       public         heap    postgres    false    4            ?            1259    17056    customer_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.customer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.customer_id_seq;
       public          postgres    false    213    4            ?           0    0    customer_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.customer_id_seq OWNED BY public.customer.id;
          public          postgres    false    212            b           2604    17051    book id    DEFAULT     b   ALTER TABLE ONLY public.book ALTER COLUMN id SET DEFAULT nextval('public.book_id_seq'::regclass);
 6   ALTER TABLE public.book ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    210    211    211            c           2604    17060    customer id    DEFAULT     j   ALTER TABLE ONLY public.customer ALTER COLUMN id SET DEFAULT nextval('public.customer_id_seq'::regclass);
 :   ALTER TABLE public.customer ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    213    212    213            ?          0    17048    book 
   TABLE DATA           8   COPY public.book (id, title, genre, author) FROM stdin;
    public          postgres    false    211   ?       ?          0    17057    customer 
   TABLE DATA           @   COPY public.customer (id, firstname, lastname, age) FROM stdin;
    public          postgres    false    213   S       ?           0    0    book_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.book_id_seq', 21, true);
          public          postgres    false    210                        0    0    customer_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.customer_id_seq', 3, true);
          public          postgres    false    212            e           2606    17055    book book_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.book
    ADD CONSTRAINT book_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.book DROP CONSTRAINT book_pkey;
       public            postgres    false    211            ?   ?   x?e??
?@???)?B?.?UQ?"?.????&??ý?7?K????	pT??? ?#]?<)???y?p6u?Z?W?t(6:?.C3y???y}	;5S????y???y??V?
L?o??2??=?? \?&?-? R8c      ?   4   x?3????K?bΐ????\N#S.#??̒????N????JN#c?=... M?     