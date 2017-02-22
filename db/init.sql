CREATE DATABASE spary;
CREATE USER spary_admin IDENTIFIED BY 'spary@YRAPS';
GRANT ALL PRIVILEGES ON spary.* TO 'spary_admin'@'%';
FLUSH PRIVILEGES;
-- grant all on spary.* to spary_admin@localhost identified by 'spary@YRAPS';
