# Mini-e-commerce Project

# Database Design Project – E-Commerce Schema (DBML → SQL)

This project demonstrates a complete workflow for designing a relational database schema using **DBML**, exporting it to **PostgreSQL SQL**, and validating it inside a running **Postgres container via Docker**.

The project includes:

- A full **e-commerce relational schema**
- The original **DBML source (`schema.dbml`)**
- **Generated PostgreSQL DDL (`schema_postgres.sql`) using dbdiagram.io**
- A runnable **Postgres environment using Docker**
- Example queries to confirm schema correctness

---

## 🧩 Schema Overview

### 1. users
Stores customer account data.

| Field      | Type          | Notes                     |
|------------|---------------|---------------------------|
| id         | SERIAL PK     |                           |
| email      | varchar(255)  | unique, required          |
| full_name  | varchar(255)  | optional                  |
| created_at | timestamp     | defaults to now()         |

---

### 2. products

| Field        | Type            | Notes               |
|--------------|-----------------|---------------------|
| id           | SERIAL PK       |                     |
| sku          | varchar(64)     | unique, required    |
| name         | varchar(255)    | required            |
| description  | text            | optional            |
| price        | decimal(10,2)   | required            |
| available    | boolean         | defaults to true    |
| created_at   | timestamp       | defaults to now()   |

---

### 3. categories

| Field | Type          | Notes            |
|-------|---------------|------------------|
| id    | SERIAL PK     |                  |
| name  | varchar(100)  | unique, required |

---

### 4. product_categories (junction table)

| Field        | Type      | Notes                                    |
|--------------|-----------|-------------------------------------------|
| product_id   | int FK    | references products(id)                   |
| category_id  | int FK    | references categories(id)                 |
| PRIMARY KEY (product_id, category_id) | Composite PK                 |

---

### 5. orders

| Field        | Type            | Notes                    |
|--------------|-----------------|---------------------------|
| id           | SERIAL PK       |                           |
| user_id      | int FK          | references users(id)      |
| status       | varchar(50)     | defaults to 'pending'     |
| total        | decimal(12,2)   |                           |
| created_at   | timestamp       | defaults to now()         |

---

### 6. order_items

| Field        | Type            | Notes                      |
|--------------|-----------------|----------------------------|
| id           | SERIAL PK       |                            |
| order_id     | int FK          | references orders(id)      |
| product_id   | int FK          | references products(id)    |
| quantity     | int             | defaults to 1              |
| unit_price   | decimal(10,2)   | required                   |

---

## 🛠 Running PostgreSQL Using Docker

This project uses Postgres running in Docker, mapped as:

```
8040 -> 5432  #more secure to use a different port than the default
```

### Start the database:

```bash
docker compose up -d
```

Verify:

```bash
docker ps
```

---

## 🗄 Apply the Schema

Install the PostgreSQL client (WSL):

```bash
sudo apt install postgresql-client -y
```

Load the schema:

```bash
psql "postgresql://demo:demo@localhost:8040/demo" -f schema_postgres.sql
```

---

## ✔️ Test the Schema

```bash
psql "postgresql://demo:demo@localhost:8040/demo"   -c "INSERT INTO users (email, full_name) VALUES ('tia@gimmy.com','Alice'); SELECT * FROM users;"
```

---

## 🔄 Resetting the Schema

```bash
psql "postgresql://demo:demo@localhost:8040/demo"   -c "DROP TABLE IF EXISTS order_items, orders, product_categories, products, categories, users CASCADE;"
```

Then reapply:

```bash
docker compose up -d
psql "postgresql://demo:demo@localhost:8040/demo" -f schema_postgres.sql
psql "postgresql://demo:demo@localhost:8040/demo" -c "INSERT INTO users (email, full_name) VALUES ('kia.gimmy@gmail.com','Alice'); SELECT * FROM users;"
[+] Running 1/0
 ✔ Container dbdesignproject-db-1  Running                                                                                           0.0s 
psql:schema_postgres.sql:6: ERROR:  relation "users" already exists
psql:schema_postgres.sql:16: ERROR:  relation "products" already exists
psql:schema_postgres.sql:21: ERROR:  relation "categories" already exists
psql:schema_postgres.sql:27: ERROR:  relation "product_categories" already exists
psql:schema_postgres.sql:35: ERROR:  relation "orders" already exists
psql:schema_postgres.sql:43: ERROR:  relation "order_items" already exists
ALTER TABLE
ALTER TABLE
ALTER TABLE
ALTER TABLE
ALTER TABLE
 id |        email        | full_name |         created_at         
----+---------------------+-----------+----------------------------
  1 | tia.gimmy@gmail.com | Alice     | 2025-12-12 20:53:09.73588
  2 | kia.gimmy@gmail.com | Alice     | 2025-12-12 20:58:55.843224
(2 rows)
```

---

## 📝 Files in This Project

```
schema.dbml
schema_postgres.sql
docker-compose.yml
README.md
```

---

## 📌 Purpose

This project demonstrates:

- DBML schema design  
- SQL DDL generation  
- Dockerized Postgres setup  
- Real schema validation  



