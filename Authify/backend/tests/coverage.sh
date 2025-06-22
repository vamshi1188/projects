#!/bin/bash

echo "🔍 Running tests with coverage..."

cd ..
go test -v -coverprofile=coverage.out ./...

go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html

echo "✅ Coverage report generated at coverage.html"
