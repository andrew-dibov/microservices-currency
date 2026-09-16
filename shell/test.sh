#!/bin/bash
# set -e

PROTO_PATH="${PROTO_PATH:-./proto/currency/}"
PROTO_NAME="${PROTO_NAME:-currency.proto}"

HOST="${HOST:-localhost}"
PORT="${PORT:-50052}"

BASE_URL="$HOST:$PORT"

# ---

echo "test 01"
echo ""

grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"fromCurrency":"USD"}' "$BASE_URL" currency.Currency/Rate
grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"toCurrency":"EUR"}' "$BASE_URL" currency.Currency/Rate

grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"fromCurrency":"US","toCurrency":"EUR"}' "$BASE_URL" currency.Currency/Rate
grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"fromCurrency":"USD","toCurrency":"EU"}' "$BASE_URL" currency.Currency/Rate

grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"fromCurrency":"US1","toCurrency":"EUR"}' "$BASE_URL" currency.Currency/Rate
grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"fromCurrency":"USD","toCurrency":"EU1"}' "$BASE_URL" currency.Currency/Rate

# ---

echo ""
echo "test 02"
echo ""

grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{}' "$BASE_URL" currency.Currency/Rates
grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"baseCurrency":"US"}' "$BASE_URL" currency.Currency/Rates
grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"baseCurrency":"US1"}' "$BASE_URL" currency.Currency/Rates

# ---

echo ""
echo "test 03"
echo ""

grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"fromCurrency":"USD","toCurrency":"EUR"}' "$BASE_URL" currency.Currency/Rate
grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"fromCurrency":"EUR","toCurrency":"USD"}' "$BASE_URL" currency.Currency/Rate

grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"baseCurrency":"USD"}' "$BASE_URL" currency.Currency/Rates | head -n 10
grpcurl -plaintext -import-path "$PROTO_PATH" -proto "$PROTO_NAME" -d '{"baseCurrency":"EUR"}' "$BASE_URL" currency.Currency/Rates | head -n 10