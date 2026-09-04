# Gator CLI

Gator is a command-line RSS feed aggregator written in Go. It allows users to register and manage accounts, follow RSS feeds, aggregate posts from feeds, and browse the collected posts.

## Prerequisites

Before installing Gator, make sure you have the following installed:

* [Go](https://go.dev/)
* [PostgreSQL](https://www.postgresql.org/)

Gator requires PostgreSQL to store users, feeds, feed follows, and posts.

## Installation

Clone the repository:

```bash
git clone https://github.com/alexandre-j95/gator_cli.git
cd gator_cli
```

Install the Gator CLI with `go install`:

```bash
go install github.com/alexandre-j95/gator_cli@latest
```

After installation, the `gator` executable should be available in your `$PATH`.

You can verify the installation with:

```bash
gator
```

> `go run .` can be used while developing the project, but `gator` is the intended command to use after installation.

## Database Setup

Gator uses PostgreSQL as its database.

First, create a database named `gator`:

```bash
createdb gator
```

The default connection string used in the configuration is:

```text
postgres://postgres:postgres@localhost:5432/gator?sslmode=disable
```

Make sure your PostgreSQL server is running and that the `postgres` user can connect using the credentials above.

### Create the database schema

The database schema is provided in the `sql/schema` directory.

Apply the schema files in order:

```bash
psql "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" -f sql/schema/001_users.sql
psql "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" -f sql/schema/002_feeds.sql
psql "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" -f sql/schema/003_feedfollows.sql
psql "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" -f sql/schema/004_feeds_lastfetched.sql
psql "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" -f sql/schema/005_posts.sql
```

## Configuration

Gator uses a configuration file located at:

```text
~/.gatorconfig.json
```

Create the file with the following contents:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

The `current_user_name` field is updated when you log in or register a user.

## Getting Started

### 1. Register a user

Create your first Gator user:

```bash
gator register john
```

The newly registered user will become the current user.

You can also log in as an existing user:

```bash
gator login john
```

### 2. Add an RSS feed

Add a feed by providing its name and URL:

```bash
gator addfeed "Hacker News" "https://news.ycombinator.com/rss"
```

Adding a feed also causes the current user to follow that feed.

### 3. View available feeds

List the feeds currently stored in the database:

```bash
gator feeds
```

### 4. Aggregate feed posts

Fetch posts from the configured feeds:

```bash
gator agg 1m
```

The argument is a time interval. For example:

```bash
gator agg 30s
gator agg 5m
gator agg 1h
```

The aggregator will periodically fetch feeds and store their posts in the database.

### 5. Browse posts

View posts that have been collected (defaults to only 2 posts):

```bash
gator browse
```

You can optionally specify how many posts to display:

```bash
gator browse 10
```

## Commands

| Command                | Description                                      |
| ---------------------- | ------------------------------------------------ |
| `register <name>`      | Register a new user                              |
| `login <name>`         | Log in as an existing user                       |
| `reset`                | Reset the application data                       |
| `users`                | List registered users                            |
| `addfeed <name> <url>` | Add a new RSS feed                               |
| `feeds`                | List available feeds                             |
| `follow <url>`         | Follow an existing feed                          |
| `following`            | List feeds followed by the current user          |
| `unfollow <url>`       | Stop following a feed                            |
| `agg <duration>`       | Fetch posts from feeds at the specified interval |
| `browse [limit]`       | Browse collected posts                           |

## Typical Workflow

A typical session might look like this:

```bash
# Register and log in
gator register john

# Add an RSS feed
gator addfeed "Hacker News" "https://news.ycombinator.com/rss"

# Fetch posts
gator agg 1m

# Browse the collected posts
gator browse 10

# See available feeds
gator feeds

# See the feeds you follow
gator following
```

## Development

While developing the application, you can run it directly from the project directory:

```bash
go run .
```

For example:

```bash
go run . register john
```

For normal usage after installation, use the compiled `gator` executable instead:

```bash
gator register john
```

Go programs are statically compiled, so after running `go install` or building the program, the resulting executable can be used without requiring the Go toolchain to run the application.

## Project Structure

```text
gator_cli/
├── internal/
│   ├── config/       # Configuration handling
│   └── database/     # Database access generated by sqlc
├── sql/
│   ├── queries/      # SQL queries
│   └── schema/       # Database schema
├── main.go           # CLI entry point
├── go.mod            # Go module definition
└── sqlc.yaml         # sqlc configuration
```

## Technologies

* Go
* PostgreSQL
* SQL
* sqlc
* RSS/Atom feeds

## Repository

The source code for Gator is available on GitHub:

https://github.com/alexandre-j95/gator_cli
