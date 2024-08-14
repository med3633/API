# dynamic id 
go get github.com/gorilla/mux
# 0.19
# 0.16 integrat db
# command mysqlmysql> create database learning;
Query OK, 1 row affected (0.00 sec)

mysql> use learning;
Database changed
mysql> create TABLE data(id INT PRIMARY KEY , name VARCHAR(255));
Query OK, 0 rows affected (0.01 sec)

mysql> desc data;
+-------+--------------+------+-----+---------+-------+
| Field | Type         | Null | Key | Default | Extra |
+-------+--------------+------+-----+---------+-------+
| id    | int          | NO   | PRI | NULL    |       |
| name  | varchar(255) | YES  |     | NULL    |       |
+-------+--------------+------+-----+---------+-------+
2 rows in set (0.00 sec)

mysql> insert into data values(1,"amine");
Query OK, 1 row affected (0.00 sec)

mysql> insert into data values(2,"med");
Query OK, 1 row affected (0.01 sec)

mysql> desc data;
+-------+--------------+------+-----+---------+-------+
| Field | Type         | Null | Key | Default | Extra |
+-------+--------------+------+-----+---------+-------+
| id    | int          | NO   | PRI | NULL    |       |
| name  | varchar(255) | YES  |     | NULL    |       |
+-------+--------------+------+-----+---------+-------+
2 rows in set (0.01 sec)

mysql> select * from data;
+----+-------+
| id | name  |
+----+-------+
|  1 | amine |
|  2 | med   |
+----+-------+
2 rows in set (0.00 sec)


