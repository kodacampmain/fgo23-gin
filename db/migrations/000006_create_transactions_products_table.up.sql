-- public.transactions_products definition
CREATE TABLE public.transactions_products (
    transaction_id int4 NOT NULL,
    product_id int4 NOT NULL,
    qty int4 DEFAULT 1 NOT NULL
);
-- public.transactions_products foreign keys
ALTER TABLE public.transactions_products
ADD CONSTRAINT transactions_products_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id);
ALTER TABLE public.transactions_products
ADD CONSTRAINT transactions_products_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);