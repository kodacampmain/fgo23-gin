-- public.employee definition
CREATE TABLE IF NOT EXISTS public.employee (
    id serial NOT NULL,
    "name" varchar NOT NULL,
    salary int4 NOT NULL,
    city varchar NULL,
    CONSTRAINT employee_pkey PRIMARY KEY (id)
);