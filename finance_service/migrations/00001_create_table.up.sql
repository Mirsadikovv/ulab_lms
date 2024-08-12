CREATE SEQUENCE superadmin_login_seq START WITH 1;


CREATE TYPE gender AS ENUM (
  male,
  female,
  other
);

CREATE TYPE course_category AS ENUM (
  backend,
  frontend,
  mobile
);



CREATE TYPE discount_type AS ENUM (
  two_courses,
  family,
  more_3_monthes
);

CREATE TYPE week_days AS ENUM (
  mon,
  tue,
  wed,
  thu,
  fri,
  sat,
  sun
);


CREATE TABLE student (
  id uuid PRIMARY KEY,
  group_id UUID[],
  schedule_id UUID,
  photo_link VARCHAR,
  fullname VARCHAR(35) NOT NULL,
  email VARCHAR NOT NULL,
  phone VARCHAR(13) NOT NULL,
  student_password VARCHAR NOT NULL,
  discount_type discount_type NOT NULL,
  discount_percent float,
  father_name VARCHAR(35),
  mother_name VARCHAR(35),
  father_phone VARCHAR(13),
  mother_phone VARCHAR(13),
  stuednt_address TEXT,
  comment TEXT,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE course (
  id uuid PRIMARY KEY,
  branch_id uuid,
  slug VARCHAR NOT NULL,
  photo_link VARCHAR,
  course_name VARCHAR NOT NULL,
  category course_category,
  course_description TEXT,
  cost float NOT NULL,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE groups (
  id uuid PRIMARY KEY,
  course_id uuid,
  room_id uuid,
  group_name VARCHAR NOT NULL,
  started_date DATE,
  finished_date DATE,
  started_time TIME,
  finished_time TIME,
  lesson_days week_days[],
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE schedule (
  id uuid PRIMARY KEY,
  teacher_id uuid,
  group_id uuid,
  started_date DATE,
  finished_date DATE,
  started_time TIME,
  finished_time TIME,
  lesson_days week_days[],
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE room (
  id uuid PRIMARY KEY,
  branch_id uuid,
  room_number VARCHAR NOT NULL,
  capacity INT,
  area float,
  rent_cost float,
  cost_for_person float,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE payments (
  id uuid PRIMARY KEY,
  student_id uuid NOT NULL,
  course_id uuid NOT NULL,
  terminal_paid FLOAT,
  cash_paid FLOAT,
  comment TEXT,
  date_of_payment date,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE salaries (
  id uuid PRIMARY KEY,
  cost_name VARCHAR NOT NULL,
  employee_id uuid NOT NULL,
  terminal_paid float,
  cash_paid float,
  comment TEXT,
  date_of_salary DATE,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE expenses (
  id uuid PRIMARY KEY,
  cost_name VARCHAR NOT NULL,
  terminal_paid float,
  cash_paid float,
  comment TEXT,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE student_performance (
  id UUID PRIMARY KEY,
  student_id UUID NOT NULL,
  schedule_id UUID NOT NULL,
  attended BOOL NOT NULL,
  task_score DECIMAL DEFAULT 0,
  created_at TIMESTAMP DEFAULT (NOW()),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

ALTER TABLE groups ADD FOREIGN KEY (room_id) REFERENCES room (id);

ALTER TABLE schedule ADD FOREIGN KEY (group_id) REFERENCES groups (id);

ALTER TABLE student_performance ADD FOREIGN KEY (student_id) REFERENCES student (id);

ALTER TABLE groups ADD FOREIGN KEY (course_id) REFERENCES course (id);

ALTER TABLE student ADD FOREIGN KEY (schedule_id) REFERENCES schedule (id);

ALTER TABLE student ADD FOREIGN KEY (group_id) REFERENCES groups (id);