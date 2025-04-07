#!/bin/sh

set -e

host="$1"
shift
cmd="$@"

echo "⏳ Waiting for $host to be available..."

until curl -s "$host" > /dev/null; do
  sleep 2
done

echo "✅ $host is ready, starting app..."
exec $cmd

