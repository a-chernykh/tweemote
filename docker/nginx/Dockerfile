FROM nginx:1.13.0

RUN mkdir /etc/nginx/ssl
RUN chmod 700 /etc/nginx/ssl

COPY bin/entrypoint.sh /usr/local/bin/

COPY nginx.conf.tmpl /etc/nginx/nginx.conf.tmpl
COPY default.conf /etc/nginx/conf.d/default.conf

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
