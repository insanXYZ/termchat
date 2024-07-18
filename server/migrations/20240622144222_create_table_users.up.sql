create table users
(
    id varchar(50) primary key ,
    name varchar(50) not null ,
    email varchar(50) not null ,
    password text not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
) engine = innodb;