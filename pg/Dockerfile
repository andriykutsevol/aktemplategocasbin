# Dockerfile

FROM postgres:16

# Add your initialization scripts
COPY ./init_db /docker-entrypoint-initdb.d

COPY ./config/pg_hba.conf /tmp/pg_hba.conf

COPY ./docker-entrypoint.sh /usr/local/bin

# # Create an entrypoint script
# COPY ./config/docker-entrypoint.sh /docker-entrypoint.sh
# RUN chmod +x /docker-entrypoint.sh

# # Override the default entrypoint with our custom script
# ENTRYPOINT ["/docker-entrypoint.sh"]