#!/bin/sh
set -e

CONFIG_FILE="/home/goapp/app/config.env"

echo "Creating config.env from environment variables..."

# Buat config.env baru dari environment variables
cat > "$CONFIG_FILE" << EOF
DATABASE_DRIVER=${DATABASE_DRIVER:-postgres}
DATABASE_SOURCE=${DATABASE_SOURCE:-postgresql://postgres:qwer1234@127.0.0.1:5432/monitoring?sslmode=disable}
SERVER_ENVIRONMENT=${SERVER_ENVIRONMENT:-debug}
SERVER_ADDRESS=${SERVER_ADDRESS:-0.0.0.0:8080}
GMAIL_USER=${GMAIL_USER:-someemail@gmail.com}
GMAIL_PASS=${GMAIL_PASS:-somepassword}
EOF

echo "Config created successfully!"
echo "Starting application..."

# Jalankan aplikasi
exec "$@"
