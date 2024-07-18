create table chats
(
    id int auto_increment primary key,
    message text not null,
    sender_id varchar(50) not null,
    receiver_id varchar(50) not null,
    created_at timestamp default current_timestamp,
    foreign key (sender_id) references users(id),
    foreign key (receiver_id) references users(id)
) engine = innodb;