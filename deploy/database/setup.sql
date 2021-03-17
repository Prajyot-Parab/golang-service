CREATE DATABASE directory;
\c directory;
CREATE TABLE employees (id int, name varchar(80), role varchar(80), org varchar(80), updated_at timestamp with time zone, deleted_at timestamp with time zone, created_at timestamp with time zone);
INSERT INTO employees (id,name,role,org) VALUES (101, 'Rob Tucker', 'Software Engineer', 'Persistent Systems');
INSERT INTO employees (id,name,role,org) VALUES (102, 'Bob Marley', 'Senior Software Engineer', 'RedHat');
INSERT INTO employees (id,name,role,org) VALUES (103, 'Kenny Boss', 'Software Engineer', 'IBM');
INSERT INTO employees (id,name,role,org) VALUES (104, 'Roger Turk', 'Lead Software Engineer', 'Persistent Systems');
INSERT INTO employees (id,name,role,org) VALUES (105, 'Steve Hope', 'Team Lead', 'Persistent Systems');

