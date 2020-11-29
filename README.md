# gofuel

An attempt at building a Golang powered Fuel tracking system for motorists. Or anybody else really.

## End points

|endpoint|parameters|notes|
|-|-|-|
|`/health-check`||Returns system status|

### `/health-check`

Returns:

```json
{
    "alive": true,
    "version": "v0.0.2"
}
```

## Environment variables

A number of enviroment variables are required in production. For the benefit of development and test, you can create `.env` and `.test.env` files based on the `example.env` file in the repository. Or create the requesite files in your environment.

## Generate a new migration

Migrations are created using the `migrate` tool. You generate the migration files, then need to edit the files to add the actual migration.

```bash
migrate create -ext sql -dir internal/pkg/migrations/files create_users_table
ls internal/pkg/migrations/files
# 20191119133115_create_users_table.down.sql  20191119133115_create_users_table.up.sql

```

Edit the files (this example adds basic user details):

```bash
# internal/pkg/migrations/files/20191119110032_create_table.down.sql
DROP TABLE IF EXISTS users;
```

```bash
# internal/pkg/migrations/files/20191119110032_create_table.up.sql
CREATE TABLE IF NOT EXISTS users (
    user_id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(300),
    email VARCHAR(300) UNIQUE NOT NULL,
    password VARCHAR(500),
    updated_at BIGINT
);
```

## Stolen from a stack overflow answer

Long answer: Ruby on Rails is a framework, Go is a language. Making a "universal" authentication system for Go would be a huge task and/or would have to be very opinionated by design, as most auth systems rely on a session store and/or database. Rails can do this (with libs. like Devise) because parts of the framework like ActiveRecord and ActionController provide an abstract API that Devise can talk to.

For the most part, you'll need to tie together a few libraries to get what you need w/ Go. "Gluing" things together is a commonly preferred way to do things, and monolithic/kitchen sink frameworks aren't typically favoured.

Some suggestions for libraries:

* An [OAuth/OAuth2 package](https://github.com/markbates/goth)
* [user sessions](http://www.gorillatoolkit.org/pkg/sessions)
* [password hashing](https://godoc.org/golang.org/x/crypto/bcrypt)
* [a micro-framework [think: Sinatra-level] w/ a solid middleware API](https://goji.io/)
* [a convenience wrapper for Go's standard database/sql package](https://github.com/jmoiron/sqlx)
* [a POST form to Go struct library](http://www.gorillatoolkit.org/pkg/schema)
