FROM nginx:latest

# Remove default Nginx configuration files
RUN rm /etc/nginx/conf.d/default.conf

# Copy custom configuration files
COPY nginx.conf /etc/nginx/conf.d/
COPY ssl/server.crt /etc/nginx/conf.d/ssl/
COPY ssl/server.key /etc/nginx/conf.d/ssl/
