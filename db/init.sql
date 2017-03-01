CREATE DATABASE spary;
CREATE DATABASE spary_test;
CREATE USER spary_admin IDENTIFIED BY 'spary@YRAPS';
GRANT ALL PRIVILEGES ON spary.* TO 'spary_admin'@'%';
GRANT ALL PRIVILEGES ON spary_test.* TO 'spary_admin'@'%';
FLUSH PRIVILEGES;
-- grant all on spary.* to spary_admin@localhost identified by 'spary@YRAPS';
