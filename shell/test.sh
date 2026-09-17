#!/bin/bash

source .env

BASE_URL="localhost:${HOST_PORT:-50052}"

# ---

echo "test 01"
echo ""

grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"fromCurrency":"USD"}' "$BASE_URL" currency.Currency/Rate
grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"toCurrency":"EUR"}' "$BASE_URL" currency.Currency/Rate
grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"fromCurrency":"US","toCurrency":"EUR"}' "$BASE_URL" currency.Currency/Rate
grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"fromCurrency":"USD","toCurrency":"EU"}' "$BASE_URL" currency.Currency/Rate
grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"fromCurrency":"US1","toCurrency":"EUR"}' "$BASE_URL" currency.Currency/Rate
grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"fromCurrency":"USD","toCurrency":"EU1"}' "$BASE_URL" currency.Currency/Rate

echo ""
echo "test 02"
echo ""

grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{}' "$BASE_URL" currency.Currency/Rates
grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"baseCurrency":"US"}' "$BASE_URL" currency.Currency/Rates
grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"baseCurrency":"US1"}' "$BASE_URL" currency.Currency/Rates

echo ""
echo "test 03"
echo ""

grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"fromCurrency":"USD","toCurrency":"EUR"}' "$BASE_URL" currency.Currency/Rate
grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"fromCurrency":"EUR","toCurrency":"USD"}' "$BASE_URL" currency.Currency/Rate
grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"baseCurrency":"EUR"}' "$BASE_URL" currency.Currency/Rates | head -n 10
grpcurl -plaintext -import-path proto/currency/ -proto currency.proto -d '{"baseCurrency":"USD"}' "$BASE_URL" currency.Currency/Rates | head -n 10