# Use the official PostgreSQL image as the base
FROM postgres:latest

# Set environment variables for PostgreSQL
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD password

# Copy SQL script to create database and table
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the default PostgreSQL port (5432)
EXPOSE 5432
