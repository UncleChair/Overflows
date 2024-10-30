# The "local" environment represents our local testings.
env "local" {
  src = "file://manifest/database/tables/"
  dev = "docker://postgres/latest/overflows"
  url = "postgres://overflower:Passw0rd@localhost:5432/overflows?sslmode=disable"
  migration {
    dir = "file://manifest/database/migrations/"
  }
}

# The "cloud" environment represents our local testings, but could use atlas cloud service
env "cloud" {
  src = "file://manifest/database/tables/"
  dev = "docker://postgres/latest/overflows"
  url = "postgres://overflower:Passw0rd@localhost:5432/overflows?sslmode=disable"
  migration {
    dir = "atlas://overflows"
  }
}

# The "docker" environment represents our docker development, can only be used for migration apply
env "docker" {
  src = "file://manifest/database/tables/"
  dev = "docker://postgres/latest/overflows"
  url = "postgres://overflower:Passw0rd@database:5432/overflows?sslmode=disable"
  migration {
    dir = "file://manifest/database/migrations/"
  }
}