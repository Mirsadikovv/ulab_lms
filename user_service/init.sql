-- Создание базы данных
CREATE DATABASE lms_user_service;

-- Создание пользователя и назначение прав
CREATE USER mirodil WITH PASSWORD 1212;

GRANT ALL PRIVILEGES ON DATABASE lms_user_service TO mirodil;