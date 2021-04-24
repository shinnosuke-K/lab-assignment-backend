use questionnaire;

create table if not exists students (
    user_id varchar(255) not null,
    student_num int,
    password varchar(255),
    entered bool,
    PRIMARY KEY (user_id)
);

create table if not exists  labs (
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
    FOREIGN KEY (user_id) REFERENCES students(user_id)
)
