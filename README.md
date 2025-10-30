# Subscription Manager

**Subscription Manager** — это REST API для управления подписками пользователей.
Позволяет создавать, получать, обновлять, удалять подписки и вычислять итоговую сумму по фильтрам.

---

## 📦 Версия

- API Version: 1.0
- Host: `localhost:8080`
- BasePath: `/`

---

## 🔧 Установка

1.  Клонируйте репозиторий:

    ```bash
    git clone https://github.com/LashkaPashka/SubManager.git
    cd SubManager
    ```

2.  Настройте переменные окружения в `.env`:

    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=yourpassword
    DB_NAME=submanager
    ```

3.  Соберите и запустите сервис:

    ```bash
    docker-compose up -d
    ```

API будет доступен по адресу: `http://localhost:8080`
Swagger API будет доступен по адресу: `http://localhost:8080/swagger/`

---

## 📚 Endpoints

### 1️⃣ Создать подписку

`POST /subscriptions`

Создаёт подписку в базе данных.

**Пример запроса:**

```json
{
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "07-2025",
  "end_date": "08-2025"
}
```

**Responses:**

-   `200`: `Subscription was created!`
-   `400`: `Invalid request. Please check the submitted data.`
-   `500`: `Internal server error`

### 2️⃣ Получить подписки

`GET /subscriptions`

Получает подписки конкретного пользователя.

**Параметры:**

| Name    | Description     | Required |
| :------ | :-------------- | :------- |
| user_id | ID пользователя | true     |

**Пример ответа:**

```json
{
  "subscriptions": [
    {
      "service_name": "Netflix",
      "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
      "price": 899,
      "start_date": "01-2024",
      "end_date": "01-2025"
    }
  ]
}
```

### 3️⃣ Удалить подписку

`DELETE /subscriptions`

Удаляет подписку по `sub_id` и `user_id`.

**Пример запроса:**

```json
{
  "subscription_id": "123e4567-e89b-12d3-a456-426614174000",
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba"
}
```

**Responses:**

-   `200`: `Subscription was deleted!`
-   `400`: `Invalid request. Please check the submitted data.`
-   `500`: `Internal server error`

### 4️⃣ Обновить подписку

`PATCH /subscriptions`

Обновляет цену подписки конкретного пользователя.

**Пример запроса:**

```json
{
  "sub_id": "123e4567-e89b-12d3-a456-426614174000",
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "price": 999
}
```

**Пример ответа:**

```json
{
  "sub_id": "123e4567-e89b-12d3-a456-426614174000",
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "service_name": "Netflix",
  "price": 999
}
```

**Responses:**

-   `200`: `OK`
-   `400`: `Invalid request. Please check the submitted data.`
-   `500`: `Internal server error`

### 5️⃣ Получить итоговую сумму

`GET /subscriptions/total`

Возвращает итоговую сумму по переданным параметрам: только по дате, по дате и `userID` или по дате и названию сервиса.

**Параметры:**

| Name         | Description                      | Required |
| :----------- | :------------------------------- | :------- |
| month        | Месяц и год (MM-YYYY)            | true     |
| user_id      | ID пользователя (UUID)           | false    |
| service_name | Название сервиса (например, Netflix) | false    |

**Пример ответа:**

```json
{
  "month": "09-2025",
  "total_sum": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "service_name": "Netflix"
}
```

**Responses:**

-   `200`: `OK`
-   `400`: `Invalid request. Please check the submitted data.`
-   `500`: `Internal server error`

---

## 🔖 Модели

### SubscriptionRequest

```json
{
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "07-2025",
  "end_date": "08-2025"
}
```

### SubscriptionsResponse

```json
{
  "subscriptions": [
    {
      "service_name": "Netflix",
      "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
      "price": 899,
      "start_date": "01-2024",
      "end_date": "01-2025"
    }
  ]
}
```

### ResponseDate

```json
{
  "month": "09-2025",
  "total_sum": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "service_name": "Netflix"
}
```

### SubscriptionResponse

```json
{
  "sub_id": "123e4567-e89b-12d3-a456-426614174000",
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "service_name": "Netflix",
  "price": 999
}
```

---

## ⚡ Технологии

-   Go (Golang)
-   PostgreSQL
-   Docker & Docker Compose
-   Swagger (OpenAPI)
