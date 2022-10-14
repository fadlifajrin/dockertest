create table tbl_product
(
	product_id bigserial not null primary key,
	product_name varchar not null,
    product_price float not null
);

insert into tbl_product(product_name, product_price) values('prod a', 1000);
insert into tbl_product(product_name, product_price) values('prod b', 2000);
insert into tbl_product(product_name, product_price) values('prod c', 3000);
insert into tbl_product(product_name, product_price) values('prod d', 4000);
