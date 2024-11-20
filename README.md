## Endpoint'ы

- `/api/v1/create` - создать песню
- `/api/v1/get/{id}?limit=,offset=` - получить текст песни по id и по строкам пагинация limit - сколько строк, offset - начальный id
- `/api/v1/get?limit=,offset=` - получить список песен, limit - сколько песен, offset - начальный id
- `/api/v1/update/{id}` - обновить песню
- `/api/v1/delete/{id}` - удалить песню
- `/swagger` - swagger-ui

## Инструкция по запуску проекта

### Без контейнера
```bash
go run main.go
```

### Запустить docker-compose

```bash
make docker_up
```

### Остановить контейнеры

```bash
make docker_stop
```

### Остановить и удалить контейнеры

```bash
make docker_down
```