CREATE TABLE "user" (
    id serial primary key ,
    first_name varchar(50) DEFAULT NULL,
    last_name varchar(50) DEFAULT NULL,
    password varchar(255) DEFAULT NULL,
    age int DEFAULT NULL,
    email varchar(50) DEFAULT NULL,
    status_user varchar(50) DEFAULT NULL,
    education_level varchar(50) DEFAULT NULL,
    study_program_id int DEFAULT NULL,
    school_id int DEFAULT NULL,
    admission_year timestamp DEFAULT NULL,
    graduation_year timestamp DEFAULT NULL,
    is_admin bool DEFAULT NULL,
    is_verificated bool DEFAULT NULL,
    is_visible bool DEFAULT NULL,
    avatar_path varchar(100) DEFAULT NULL,
    is_moderated bool DEFAULT NULL,
    activation_code varchar(10)
)

CREATE TABLE "coincidence" (
   id serial primary key ,
   user1_id int DEFAULT NULL,
   user2_id int DEFAULT NULL,
   coincidence_time timestamp NULL DEFAULT NULL
)

CREATE TABLE "dialog" (
  id serial primary key,
  user1_id int DEFAULT NULL,
  user2_id int DEFAULT NULL,
  time_mess timestamp NULL DEFAULT NULL
)

CREATE TABLE "faculty" (
   id serial primary key,
   name_program varchar(240) DEFAULT NULL
)

CREATE TABLE "post" (
    id serial primary key,
    user_id int DEFAULT NULL,
    text_post text,
    is_moderated bool DEFAULT NULL,
    publication_time timestamp NULL DEFAULT NULL
)

CREATE TABLE "school" (
  id serial primary key,
  school_name varchar(255) DEFAULT NULL
)