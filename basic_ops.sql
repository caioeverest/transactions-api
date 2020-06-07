create table operations(operation_type_id serial not null constraint operations_pkey primary key, description text not null);
alter table operations owner to admin;
insert into operations (operation_type_id, description) values (1, 'COMPRA A VISTA');
insert into operations (operation_type_id, description) values (2, 'COMPRA PARCELADA');
insert into operations (operation_type_id, description) values (3, 'SAQUE');
insert into operations (operation_type_id, description) values (4, 'PAGAMENTO');