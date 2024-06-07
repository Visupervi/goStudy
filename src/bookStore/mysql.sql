create table carts
(
    id           varchar(100) primary key,
    total_count  int not null,
    total_amount double(11, 2
) not null, user_id int not null, foreign key(user_id) references users(id));

create table sessions
(
    session_id varchar(100) primary key,
    user_id    int          not null,
    foreign key (user_id) references users (id),
    username   varchar(100) not null
);

create table books
(
    id     int primary key auto_increment,
    title  varchar(100) not null,
    author varchar(100) not null,
    price  double(11, 2
) not null, sales int not null, stock int not null, img_path varchar(100));

create table cart_items
(
    id     int primary key auto_increment,
    count  int not null,
    amount double(11, 2) not null,
    book_id int not null,
    cart_id varchar(100) not null,
     foreign key(book_id) references books(id),
      foreign key(cart_id) references carts(id)
);


-- 创建订单表
create table orders
(
    id           varchar(100) primary key,
    create_time  datetime not null,
    total_count  int      not null,
    total_amount double (11, 2) not null,
    state int not null,
    user_id int,
    foreign key(user_id) references users(id)
    );

-- 创建订单项表
create table order_items
(
    id     int primary key auto_increment,
    count  int not null,
    amount double(11, 2) not null,
    title varchar (100) not null,
    author varchar (100) not null,
    price double(11,2) not null,
    img_path varchar(100) not null,
    order_id varchar (100)  not null,
    foreign key (order_id) references orders(id)
);