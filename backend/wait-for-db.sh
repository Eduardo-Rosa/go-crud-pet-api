#!/bin/sh
set -e

echo "Aguardando o banco de dados ficar pronto..."
until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER"; do
  sleep 2
done

echo "Banco de dados está pronto! Iniciando a aplicação..."
exec "$@"