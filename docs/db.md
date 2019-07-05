# create table users.
```
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
)  ENGINE=INNODB;
```

# create table categories

```
CREATE TABLE IF NOT EXISTS categories (
    id INT AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
)  ENGINE=INNODB;
```

# create table articles.
```
CREATE TABLE IF NOT EXISTS articles (
    id INT AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    user_id int NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)

)  ENGINE=INNODB;
```

# create table category_article
```

CREATE TABLE IF NOT EXISTS category_article (
    category_id int NOT NULL,
    article_id int NOT NULL,
    PRIMARY KEY (category_id, article_id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (article_id) REFERENCES articles(id)
)  ENGINE=INNODB;

```

# select all article with user name.

```
select articles.*, users.name as user_name from articles inner join users on articles.user_id = users.id;
+----+-----------+---------+-------------+
| id | title     | user_id | user_name   |
+----+-----------+---------+-------------+
|  3 | Title1    |      16 | Edited Name |
|  4 | Title1 2  |      16 | Edited Name |
|  5 | Title1 32 |      16 | Edited Name |
+----+-----------+---------+-------------+

```

# select single article with user name

```
mysql> select articles.*, users.name as user_name from articles inner join users on articles.user_id = users.id where articles.id=3;
+----+--------+---------+-------------+
| id | title  | user_id | user_name   |
+----+--------+---------+-------------+
|  3 | Title1 |      16 | Edited Name |
+----+--------+---------+-------------+
```

# add collumn to table

```
 alter table users add email varchar(255) not null;

```
# Update all collums 
```
 update users set email="acb@gmail.com";

```