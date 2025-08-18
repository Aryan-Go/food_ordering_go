The 3 curl requests that I used
```
curl -X POST "http://127.0.0.1:8000/customer/food_items_added" -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImdveWFsLmFyeWFuQGdtYWlsLmNvbSIsImV4cCI6MTc1NTUyODMwMiwicm9sZSI6ImN1c3RvbWVyIn0.DD5pD47SZeXoYICcvG8crdkUML3MbrfYuer9tYlxddg" -d @data_send.json
```
```
curl -o /dev/null -s -w "Time: %{time_total}s\n" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoZWZAZ21haWwuY29tIiwiZXhwIjoxNzU1NTI1MzQ0LCJyb2xlIjoiY2hlZiJ9.MdjpqtmVHNVh2cgS2JTu5EvGfiHmChjowXQuh4GEzkM" http://127.0.0.1:8000/chef/render_order
```
```
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNoZWZAZ21haWwuY29tIiwiZXhwIjoxNzU1NTI1MzQ0LCJyb2xlIjoiY2hlZiJ9.MdjpqtmVHNVh2cgS2JTu5EvGfiHmChjowXQuh4GEzkM" http://127.0.0.1:8000/chef/render_order
```