FROM postgres:9.6.2

COPY healthcheck.sh /usr/local/bin/healthcheck.sh

# I suspect the line below is the reason of the following error: https://forums.docker.com/t/container-not-found-error/32303
# HEALTHCHECK CMD ["/usr/local/bin/healthcheck.sh"]
