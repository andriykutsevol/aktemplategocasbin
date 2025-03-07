#!/bin/bash
set -e

# Function to copy custom pg_hba.conf
copy_pg_hba_conf() {
  if [ ! -f /var/lib/postgresql/data/pg_hba.conf ]; then
    cp /tmp/pg_hba.conf /var/lib/postgresql/data/pg_hba.conf
    chown postgres:postgres /var/lib/postgresql/data/pg_hba.conf
  fi
}

# Check if the first argument is a flag (starts with -)
if [ "${1:0:1}" = '-' ]; then
    set -- postgres "$@"
fi

# If the first argument is postgres, initialize the database
if [ "$1" = 'postgres' ]; then
  # Copy custom pg_hba.conf as the postgres user
  su-exec postgres bash -c 'copy_pg_hba_conf'

  # Initialize the database if it doesn't exist
  if [ ! -s "$PGDATA/PG_VERSION" ]; then
    su-exec postgres initdb
  fi

  # Start PostgreSQL as the postgres user
  exec su-exec postgres "$@"
else
  # Otherwise, run whatever the user has specified as the command
  exec "$@"
fi




















# #!/bin/bash
# set -e

# # Check if the data directory is empty (i.e., PostgreSQL is being initialized)
# if [ -z "$(ls -A /var/lib/postgresql/data)" ]; then
#     echo "Data directory is empty. Initializing database..."
#     docker-entrypoint.sh postgres &

#     # Wait for the database initialization to complete
#     sleep 10

#     # Copy the custom pg_hba.conf file to the data directory
#     echo "Copying custom pg_hba.conf to the data directory..."
#     cp /tmp/pg_hba.conf /var/lib/postgresql/data/pg_hba.conf

#     # Stop the temporary PostgreSQL instance
#     pg_ctl -D /var/lib/postgresql/data -m fast -w stop
# fi

# # Start PostgreSQL
# echo "Starting PostgreSQL..."
# exec postgres