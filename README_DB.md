## Migrations

Create the environment file:
```
cp env-example.list .env
```

Build and start all containers:
```
docker compose up -d
```

Access pgAdmin (wait a couple of minutes for it to start):
```
http://localhost:5400/login?next=/
```

### Register the server in pgAdmin
After logging in using the credentials from your .env(PGADMIN_DEFAULT_EMAIL/PGADMIN_DEFAULT_PASSWORD), add the PostgreSQL server:
```
Click Servers → Register → Server…
```
On the General tab:
```
Name (example: local-db).
```
On the Connection tab:
```
Host name/address: db
Port: 5432
Maintenance DB: postgres (or your DB_NAME from .env)
Username: from .env (POSTGRES_USER)
Password: from .env (POSTGRES_PASSWORD)
Click Save — the server will appear in the tree and be ready to use.
```

### Create, apply, or rollback migrations
Create a new file migration:
```
make migrate-create name=add_test_table
```

Apply all pending migrations:
```
make migrate-up
```

Rollback the last migration:
```
make migrate-down
```