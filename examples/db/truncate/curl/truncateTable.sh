curl "https://api.m3o.com/v1/db/Truncate" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $MICRO_API_TOKEN" \
-d '{
  "table": "users"
}'