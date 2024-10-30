# The "local" environment represents our local testings.
env "dev" {
  src = "file://manifest/database/Pgsql/tables/"
  dev = "docker://postgres/latest/overflows"
  url = "postgres://overflower:Passw0rd@localhost:5432/overflows?sslmode=disable"
  migration {
    dir = "file://manifest/database/Pgsql/migrations/"
  }
}

# The "docker" environment represents our docker development, can only be used for migration apply
env "docker" {
  src = "file://manifest/database/Pgsql/tables/"
  dev = "docker://postgres/latest/overflows"
  url = "postgres://overflower:Passw0rd@database:5432/overflows?sslmode=disable"
  migration {
    dir = "file://manifest/database/Pgsql/migrations/"
  }
}

env "standalone" {
  src = "file://manifest/database/SQLite/tables/"
  dev = "sqlite://file?mode=memory"
  url = "sqlite://manifest/database/SQLite/overflows.db"
  migration {
    dir = "file://manifest/database/SQLite/migrations/"
  }
}