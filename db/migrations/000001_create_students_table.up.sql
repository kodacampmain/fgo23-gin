-- public.students definition
CREATE TABLE IF NOT EXISTS public.students (
	id serial NOT NULL,
	"name" varchar NULL,
	CONSTRAINT students_pkey PRIMARY KEY (id)
);