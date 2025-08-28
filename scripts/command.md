**короткий справочник по Docker Compose для  проекта** с Go и Postgres, включая миграции.

---

## **1. Поднять контейнеры (Postgres и API)**

```bash
# Поднять все сервисы из docker-compose.yml в фоне
docker-compose up -d
```

* `-d` — запуск в фоне (detached mode).
* Контейнеры будут созданы, и Postgres начнёт слушать порт.

---

## **2. Посмотреть логи контейнеров**

```bash
# Логи конкретного контейнера
docker logs -f restapi_postgres
docker logs -f restapi_api
```

* `-f` — следить за логами в реальном времени.

---

## **3. Подключиться к базе данных (Postgres)**

```bash
docker exec -it restapi_postgres psql -U admin -d restapi_db
```

* `exec -it` — интерактивный режим.
* `psql -U admin -d restapi_db` — подключение к базе `restapi_db` пользователем `admin`.

---

## **4. Создать/применить миграции**

Если сервис `migrate` в `docker-compose.yml`:

```bash
# Применить все миграции
docker-compose run migrate
```

или

```bash
docker-compose run migrate -path /migrations -database "postgres://admin:12345@db:5432/restapi_db?sslmode=disable" up
```

* `up` — применяет миграции (создаёт таблицы).

---

## **5. Откат миграций**

```bash
docker-compose run migrate -path /migrations -database "postgres://admin:12345@db:5432/restapi_db?sslmode=disable" down 1
```

* `down 1` — откат последней миграции (удаляется только таблица/шаг, который был последним).
* `down` — откат всех миграций.
* `drop -f` — полностью удалить все таблицы.

---

## **6. Проверка статуса миграций**

```bash
docker-compose run migrate -path /migrations -database "postgres://admin:12345@db:5432/restapi_db?sslmode=disable" version
```

* Показывает последнюю применённую миграцию.

---

## **7. Остановка контейнеров**

```bash
docker-compose down
```

* Останавливает и удаляет контейнеры, но **данные в volume сохраняются**

---

## **8. Список всех контейнеров и томов**

```bash
docker ps -a           # список всех контейнеров
docker volume ls       # список всех volume
docker volume inspect restapi_db_data   # посмотреть путь volume
```

---

## **9. Пересборка Go API**

Если поменял код Go:

```bash
docker-compose build api   # пересобрать только API
docker-compose up -d       # поднять контейнер
```





``` go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	dsn := "postgres://admin:12345@localhost:5434/restapi_db?sslmode=disable"
	fmt.Println("Using DSN:", dsn)

	pool, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(), "SELECT id, login, role_id FROM users")
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}

	log.Print("Users:")

	for rows.Next() {
		var id int
		var login string
		var role_id int
		err := rows.Scan(&id, &login, &role_id)
		if err != nil {
			log.Fatalf("Row scan failed: %v\n", err)
		}

		fmt.Printf("ID: %d | Login: %s | Role_id: %d \n", id, login, role_id)
	}

	if rows.Err() != nil {
		log.Fatalf("Row iteration error: %v\n", rows.Err())
	}

	defer rows.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```