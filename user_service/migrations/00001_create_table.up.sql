CREATE SEQUENCE admin_login_seq START WITH 1;



CREATE TYPE gender AS ENUM (
  male,
  female,
  other
);


CREATE TYPE salary_type AS ENUM (
  fixed,
  percent
);

CREATE TYPE discount_type AS ENUM (
  two_courses,
  family,
  more_3_monthes
);


CREATE TABLE superadmin (
  id uuid PRIMARY KEY,
  fullname VARCHAR(35),
  superadmin_login VARCHAR NOT NULL,
  superadmin_password VARCHAR NOT NULL,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE branch (
  id uuid PRIMARY KEY,
  branch_name VARCHAR(35) NOT NULL,
  branch_description TEXT,
  branch_address VARCHAR,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE admins (
  id uuid PRIMARY KEY,
  branch_id uuid,
  fullname VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  phone VARCHAR(13) NOT NULL,
  birthday DATE,
  admin_password VARCHAR NOT NULL,
  gender gender NOT NULL,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE free_user (
  id uuid PRIMARY KEY,
  fullname VARCHAR NOT NULL,
  email VARCHAR,
  phone VARCHAR(13),
  user_password VARCHAR NOT NULL,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE employee_role (
  id uuid PRIMARY KEY,
  branch_id uuid,
  role_name VARCHAR(35) NOT NULL,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE employee (
  id uuid PRIMARY KEY,
  branch_id uuid,
  role_id uuid,
  first_name VARCHAR(35) NOT NULL,
  second_name VARCHAR(35) NOT NULL,
  email VARCHAR(55) NOT NULL,
  phone VARCHAR(13) NOT NULL,
  birthday DATE,
  teacher_password VARCHAR NOT NULL,
  gender gender NOT NULL,
  salary_type salary_type NOT NULL,
  salary_percent float,
  comment TEXT,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);




ALTER TABLE admins ADD FOREIGN KEY (branch_id) REFERENCES branch (id);

ALTER TABLE employee_role ADD FOREIGN KEY (branch_id) REFERENCES branch (id);

ALTER TABLE employee ADD FOREIGN KEY (branch_id) REFERENCES branch (id);

ALTER TABLE employee ADD FOREIGN KEY (role_id) REFERENCES employee_role (id);












