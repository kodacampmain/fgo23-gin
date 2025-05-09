-- public.promos definition
CREATE TABLE IF NOT EXISTS public.promos (
    id serial4 NOT NULL,
    "name" varchar NOT NULL,
    discount int4 NOT NULL,
    CONSTRAINT promos_pkey PRIMARY KEY (id)
);