#!/bin/sh
set -e

until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER"; do
  echo "Aguardando banco de dados..."
  sleep 2
done

echo "Banco de dados está pronto! Iniciando a aplicação..."
exec "$@"
