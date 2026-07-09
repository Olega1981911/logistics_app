# Logistics CRM

Logistics CRM — это backend система для управления логистическими операциями, разработанная на Go.


---

# 🚀 Features

Планируемая функциональность:

### Core

* управление клиентами
* управление заказами перевозок
* управление автопарком
* управление водителями
* управление перевозчиками

### Finance

* счета и платежи
* расходы
* финансовая аналитика

### Documents

* генерация документов
* хранение файлов

### Future

* интеграции с GPS трекерами
* оптимизация маршрутов
* аналитика перевозок
* SaaS версия

---

# 🏗 Architecture

Проект использует **layered architecture**.

```text
HTTP Request
     ↓
Handler
     ↓
Service
     ↓
Repository
     ↓
Database
```

### Структура проекта

```text
cmd/
    main/                # точка входа приложения

internal/
    domain/             # бизнес модели
    repository/         # интерфейсы доступа к данным
    service/            # бизнес логика
    handler/            # HTTP обработчики
    middleware/         # HTTP middleware
    storage/            # реализация репозиториев

pkg/                    # переиспользуемые пакеты

migrations/             # SQL миграции

configs/                # конфигурация

tests/                  # интеграционные тесты
```

---

# 🧰 Tech Stack

| Technology     | Description                |
| -------------- | -------------------------- |
| Go             | Backend язык               |
| SQLite         | База данных для разработки |
| PostgreSQL     | Production база            |
| chi / net/http | HTTP router                |
| slog           | Логирование                |
| golang-migrate | Миграции                   |
| YAML           | Конфигурация               |

---

# ⚙️ Getting Started

## Clone repository

```bash
git clone https://github.com/yourname/logistics-crm
cd logistics-crm
```

---

## Install dependencies

```bash
go mod tidy
```

---

## Run application

```bash
make run
```

---

## Build binary

```bash
make build
```

---

## Run tests

```bash
make test
```

---

# 🗄 Database

Миграции управляются через **golang-migrate**.

### Применить миграции

```bash
make migrate-up
```

### Откатить миграцию

```bash
make migrate-down
```

---

# 🔌 API

Будущие endpoints:

```text
POST   /customers
GET    /customers
GET    /customers/{id}
PUT    /customers/{id}
DELETE /customers/{id}
```

Документация API будет добавлена позже.

---

# 📦 Development

Рекомендуемые инструменты:

* golangci-lint
* air (hot reload)
* migrate
* Docker (в будущем)

---

# 🧪 Testing

Проект будет включать:

* unit tests
* integration tests
* repository tests

Запуск:

```bash
make test
```

---

# 🛣 Roadmap

### MVP

* HTTP сервер
* база данных
* CRUD клиентов
* CRUD заказов

### Phase 2

* управление автопарком
* управление водителями
* финансовые операции

### Phase 3

* интеграции
* SaaS
* аналитика

---

# 📄 License

Лицензия будет определена позже.

Планируется модель:

Open Core + коммерческие модули.

---

# 👨‍💻 Author

Backend Developer

GitHub: https://github.com/yourname
