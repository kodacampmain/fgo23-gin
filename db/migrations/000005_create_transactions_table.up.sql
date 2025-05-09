-- public.transactions definition
CREATE TABLE IF NOT EXISTS public.transactions (
    id serial4 NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at timestamp NULL,
    student_id int4 NOT NULL,
    CONSTRAINT transactions_pkey PRIMARY KEY (id)
);
-- public.transactions foreign keys
ALTER TABLE iF EXISTS public.transactions
ADD CONSTRAINT transactions_student_id_fkey FOREIGN KEY (student_id) REFERENCES public.students(id);