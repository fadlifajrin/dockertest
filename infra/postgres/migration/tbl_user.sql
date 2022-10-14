create table tbl_user
(
	user_id bigserial not null primary key,
	email varchar not null,
    full_name varchar not null
);

insert into tbl_user(email, full_name) values('aaa@yahoo.com', 'aaa');
insert into tbl_user(email, full_name) values('bbb@yahoo.com', 'bbb');
insert into tbl_user(email, full_name) values('ccc@yahoo.com', 'ccc');
insert into tbl_user(email, full_name) values('ddd@yahoo.com', 'ddd');
