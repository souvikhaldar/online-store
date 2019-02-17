create table pdt(PID serial primary key,name text,type text,price decimal);
create table agt(aid serial primary key,name text,adhar_num text,availability boolean);
create table purchase(pur_id serial primary key, pid integer,aid integer,dop timestamp default now());