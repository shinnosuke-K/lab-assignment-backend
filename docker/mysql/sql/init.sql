use questionnaire;

create table if not exists users (
    id varchar(255) not null,
    student_num bigint,
    password varchar(255),
    graduate bool,
    entered bool,
    PRIMARY KEY (id)
);

create table if not exists  professors (
    id varchar(255) not null ,
    lab_name varchar(255),
    prof_name varchar(255),
    prof_roman varchar(255),
    PRIMARY KEY (id)
);

create table if not exists answers (
    id int not null,
    user_id varchar(255),
    prof_id varchar(255),
    point int,
    reg_at datetime,
    update_at datetime,
    PRIMARY KEY (id, user_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)
