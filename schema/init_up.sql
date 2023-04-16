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
    is_moderated bool DEFAULT NULL
)

