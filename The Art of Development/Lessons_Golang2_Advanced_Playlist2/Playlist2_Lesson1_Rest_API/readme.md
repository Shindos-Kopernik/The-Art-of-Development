# rest-api-tutoriol

# user-service

# REST API

1. Поддержка кэширования
2. Энтемподентность операций
3. Построен поверх http

GET /users -- list of users -- 200, 404, 500
GET /users/:id -- user by id -- 200, 404, 500
POST /users/:id -- create users -- 204, 4xx Header Location: url
PUT /users/:id -- fully update user -- 204/200 404, 400, 500
PATCH /users/:id -- partially update user -- 204/200, 404, 400, 500
DELETE /users/:id -- delete user by id -- 204, 404, 400