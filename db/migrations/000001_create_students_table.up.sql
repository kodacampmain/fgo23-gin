-- public.students definition
CREATE TABLE public.students (
	id serial NOT NULL,
	"name" varchar NULL,
	CONSTRAINT students_pkey PRIMARY KEY (id)
);