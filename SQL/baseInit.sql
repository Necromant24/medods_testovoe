PGDMP  "                    |            medodsDb    17.2    17.2     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false            �           1262    16388    medodsDb    DATABASE     ~   CREATE DATABASE "medodsDb" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';
    DROP DATABASE "medodsDb";
                     postgres    false            �            1259    16396    tokens    TABLE     �   CREATE TABLE public.tokens (
    id uuid NOT NULL,
    token character varying NOT NULL,
    "userId" uuid NOT NULL,
    "userIp" character varying
);
    DROP TABLE public.tokens;
       public         heap r       postgres    false            �            1259    16389    users    TABLE     Y   CREATE TABLE public.users (
    id uuid NOT NULL,
    name character varying NOT NULL
);
    DROP TABLE public.users;
       public         heap r       postgres    false            %           2606    16402    tokens tokens_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.tokens DROP CONSTRAINT tokens_pkey;
       public                 postgres    false    218            #           2606    16395    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public                 postgres    false    217           