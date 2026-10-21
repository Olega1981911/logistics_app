# Logistics CRM

Backend-система для управления логистическими операциями на Go.
Платформа для малого и среднего бизнеса: диспетчерская, отслеживание
груза от точки до клиента, ИИ-помощники для каждой роли, SaaS-подписка.

---

# 🚀 Features

### Core (MVP)
- управление клиентами (физлица и компании)
- управление заказами на доставку
- управление автопарком (машины, водители)
- управление перевозчиками

### Tracking
- приём GPS-координат от мобильного приложения водителя
- live-отображение позиции машины на карте (WebSocket)
- история перемещений (time-series, партиционирование по месяцам)
- события статуса доставки (picked_up, in_transit, delivered, exception)

### Routing
- построение маршрута от точки A до точки B (OSRM / GraphHopper)
- оптимизация порядка объезда точек (TSP / VRP)
- учёт высоты над уровнем моря и уклона (экономия топлива)
- учёт погодных условий (осадки, ветер, наводнения)
- динамический пересчёт ETA при отклонении от маршрута

### AI Agents
- **Диспетчер** — автоматическое распределение заказов,
  перестроение маршрутов при сбоях, коммуникация с водителями
- **Помощник водителя** — голосовая навигация, напоминания о режиме
  труда и отдыха, поиск дешёвого топлива и парковок
- **Бухгалтер** — аудит счетов, сверка с договорами, выявление расхождений
- **Экономист** — расчёт себестоимости рейса, прогноз расхода топлива,
  обнаружение аномалий, анализ маржинальности
- **Комплаенс** — контроль сроков страховок, техосмотров, лицензий

### Finance
- счета, платежи, расходы
- заправки (учёт топлива, расхода, аномалий)
- финансовая аналитика и дашборды

### Documents
- хранение документов со сроками действия (полиморфно: машина/водитель/клиент)
- генерация документов (в будущем)

### SaaS
- мультитенантность (изоляция данных компаний-подписчиков)
- тарифы Starter / Professional / Business
- биллинг по количеству активных машин
- подписки, периоды оплаты, интеграция с платёжными системами

---

# 🏗 Architecture

Проект использует **layered architecture**.

```text
HTTP Request / WebSocket
     ↓
Handler
     ↓
Service (бизнес-логика + вызовы AI-агентов)
     ↓
Repository
     ↓
Database (SQLite / PostgreSQL)
```

Дополнительно:
- **AI-слой** (`internal/ai/`) — отдельные агенты, общающиеся с LLM
  (локальные через Ollama или облачные через Pollinations).
- **Routing-слой** — обёртка над OSRM/GraphHopper для построения маршрутов.
- **Hub** — WebSocket-хаб для live-трекинга (горутина на водителя,
  горутина на подписчика дашборда).

### Структура проекта

```text
cmd/
    main/                # точка входа приложения
    simulator/           # симулятор GPS для демо

internal/
    config/             # загрузка конфигурации (cleanenv)
    models/             # доменные структуры (без логики)
    storage/            # подключение к БД
    repository/         # доступ к данным (SQL)
    service/            # бизнес-логика
    handler/            # HTTP-обработчики
    middleware/         # tenant, auth
    router/             # сборка chi-роутов
    ai/                 # ИИ-агенты (dispatcher, driver_assistant, ...)

pkg/                    # переиспользуемые пакеты

migrations/             # SQL миграции (golang-migrate)

configs/                # YAML-конфигурация

tests/                  # интеграционные тесты
```

### Модели данных

**Существующие:** `Client`, `Address`, `Cargo`, `Route`, `Vehicle`,
`User`, `Invoice`, `Payment`, `Expense`.

**Новые (запланированы):**
- `Tenant`, `Plan`, `Subscription` — SaaS
- `Order` — заказ на доставку
- `Driver` — водитель (отдельно от `User`)
- `LocationPoint` — GPS-точка (time-series)
- `TrackingEvent` — событие статуса
- `FuelRecord` — заправка
- `Document` — документы со сроками
- `AIAgentAction` — лог действий ИИ

---

# 🧰 Tech Stack

| Technology         | Description                                  |
| ------------------ | -------------------------------------------- |
| Go                 | Backend язык                                 |
| SQLite             | БД для разработки                            |
| PostgreSQL         | Production БД (партиционирование для GPS)    |
| chi                | HTTP router                                  |
| nhooyr/websocket   | WebSocket для live-трекинга                  |
| slog               | Структурированное логирование                |
| golang-migrate     | Миграции                                     |
| cleanenv           | Загрузка конфигурации из YAML + ENV          |
| Redis              | Кеш last-known позиций + pub/sub             |
| OSRM / GraphHopper | Маршрутизация (self-hosted, бесплатно)       |
| Ollama             | Локальные LLM для ИИ-агентов                 |

---

# ⚙️ Getting Started

## Clone repository

```bash
git clone https://github.com/Olega1981911/logistics_app
cd logistics_app
```

## Install dependencies

```bash
go mod tidy
```

## Run application

```bash
go run ./cmd/main
```

## Build binary

```bash
go build -o bin/logistics ./cmd/main
```

## Run tests

```bash
go test ./...
```

---

# 🗄 Database

Миграции управляются через **golang-migrate**.

### Применить миграции

```bash
migrate -path migrations -database "sqlite3://logistics.db" up
```

### Откатить миграцию

```bash
migrate -path migrations -database "sqlite3://logistics.db" down 1
```

### Схема (MVP)
- `clients`, `addresses` — клиенты и их адреса
- `vehicles`, `drivers` — автопарк
- `orders` — заказы на доставку
- `location_points` — GPS-точки (партиционирование по `recorded_at`)
- `tracking_events` — статусы доставки

### Схема (следующие этапы)
- `tenants`, `plans`, `subscriptions` — SaaS
- `routes`, `route_points` — маршруты
- `fuel_records`, `invoices`, `payments`, `expenses` — финансы
- `documents` — комплаенс
- `ai_agent_actions` — аудит ИИ

---

# 🔌 API (MVP)

### Customers
```text
POST   /customers
GET    /customers
GET    /customers/{id}
PUT    /customers/{id}
DELETE /customers/{id}
```

### Vehicles
```text
POST   /vehicles
GET    /vehicles
GET    /vehicles/{id}
PUT    /vehicles/{id}
DELETE /vehicles/{id}
```

### Tracking (следующий этап)
```text
POST   /api/v1/locations     # приём координат от водителя
WS     /ws/tracking          # live-обновления для дашборда
GET    /orders/{id}/events   # история событий по заказу
```

### Routing (следующий этап)
```text
POST   /routes/build         # построить маршрут
POST   /routes/{id}/recalc   # пересчитать ETA
```

Документация API будет добавлена позже (Swagger/OpenAPI).

---

# 🤖 AI Agents

Каждый агент — отдельный сервис в `internal/ai/`, общающийся с LLM
через HTTP. Может работать локально (Ollama) или через облако
(Pollinations, OpenAI).

| Агент             | Роль                                                |
| ----------------- | --------------------------------------------------- |
| `dispatcher`      | распределение заказов, перестроение маршрутов        |
| `driver_assistant`| голосовая навигация, режим труда и отдыха            |
| `accountant`      | аудит счетов, сверка с договорами                    |
| `economist`       | себестоимость, прогноз топлива, маржинальность       |
| `compliance`      | контроль сроков документов                           |

Каждое действие агента логируется в `ai_agent_actions`
(для аудита, отладки и обучения).

---

# 📱 Clients

- **Web-дашборд** — диспетчер, бухгалтер, экономист
- **Мобильное приложение водителя** — навигация, статусы, фото документов
- **Мобильное приложение клиента** — трекинг груза, ETA, уведомления

---

# 🛣 Roadmap

### MVP (текущий этап)
- [ ] HTTP сервер (chi)
- [ ] база данных (SQLite) + миграции
- [ ] CRUD клиентов (repository → service → handler)
- [ ] CRUD машин
- [ ] CRUD заказов

### Phase 2 — Tracking
- [ ] приём GPS-координат (HTTP POST)
- [ ] WebSocket-хаб для live-дашборда
- [ ] мобильное приложение водителя (React Native / Flutter)
- [ ] история перемещений (PostgreSQL, партиционирование)

### Phase 3 — Routing
- [ ] OSRM / GraphHopper в Docker
- [ ] построение маршрута A→B
- [ ] оптимизация порядка точек (TSP)
- [ ] учёт высоты/уклона (экономия топлива)
- [ ] учёт погоды (OpenWeatherMap)
- [ ] динамический пересчёт ETA

### Phase 4 — AI Agents
- [ ] интеграция с Ollama / Pollinations
- [ ] диспетчер-агент
- [ ] помощник водителя
- [ ] бухгалтер-агент (аудит счетов)
- [ ] экономист-агент (прогноз топлива)

### Phase 5 — SaaS
- [ ] мультитенантность (`tenant_id` во всех таблицах)
- [ ] тарифы и подписки
- [ ] биллинг (Stripe / Paddle)
- [ ] приложение клиента (трекинг + ETA)

---

# 🧪 Testing

Планируется:
- unit tests (service, models)
- integration tests (handler + repository + БД)
- repository tests (SQLite in-memory)

Запуск:

```bash
go test ./...
```

---

# 📦 Development

Рекомендуемые инструменты:
- `golangci-lint` — линтер
- `air` — hot reload
- `migrate` — управление миграциями
- `Docker` — для OSRM, Redis, PostgreSQL
- `wscat` — для тестирования WebSocket

---

# 📄 License

Планируется модель: **Open Core + коммерческие модули**.

Базовые модули (Core, Tracking) — открытые.
SaaS-функции, AI-агенты для бизнеса, интеграции с платёжными
системами — коммерческие.

---

# 👨‍💻 Author

Backend Developer

GitHub: [Olega1981911](https://github.com/Olega1981911)