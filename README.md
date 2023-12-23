## Learning mySQL side of databases

- [X] MySQL and stuff
    - login msql -u sername -p < you will be prompted for password >
    - SHOW DATABASES; (will list databases)
    - if database exist:
        - USE database_name; (will change to database)
        - SHOW TABLES;  (will list usernames)
- [X] Create a new database and setup new user
    - CREATE DATABASE dB_name;
    - CREATE USER 'new_user'@'localhost' IDENTIFIED BY 'new_password';
    - GRANT ALL PRIVILEGES ON dB_name.* TO 'new_user'@'localhost' identified by 'new_password';
    - FLUSH PRIVILEGES;



