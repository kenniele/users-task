# Time-tracker app

## Описание

Проект User Tasks предназначен для управления пользователями и их задачами. 

---


## Функционал

1. **Управление пользователями:**
    - Получение данных пользователей с фильтрацией и пагинацией.
    - Добавление нового пользователя.
    - Изменение данных пользователя.
    - Удаление пользователя.

2. **Управление задачами:**
    - Начало и окончание отсчета времени по задаче.
    - Получение трудозатрат по пользователю с сортировкой.

## Стек технологий

- **Backend**: Gin (Golang)
- **База данных**: PostgreSQL
- **Frontend**: HTML, CSS, JavaScript
- **Swagger**: Swaggo

## Установка и запуск

### Перед запуском

Важно, чтобы у вас был помещен файл .env в корневую папку проекта с вашими данными.

Пример содержимого `.env`:

    env
    DBUser=your_postgres_user
    DBPass=your_postgres_password
    DBHost=postgres
    DBPort=5432
    DBName=your_database_name

### Установка приложения

1. Клонируйте репозиторий:

    ```sh
    git clone https://github.com/kenniele/users-task-project.git
   ```

2. Перейдите в директорию проекта:

   ```sh
   cd users-task-project/backend
   ```

### Запуск приложения

1. Введите в консоли:

   ```sh
   go mod download
   ```

2. Запустите проект:

   ```sh
   go run main.go
   ```

3. Откройте браузер и перейдите по адресу:

   ```sh
   http://localhost:8080
   ```

   Вы также можете открыть Swagger UI для документации API по адресу:

    ```sh
    http://localhost:8080/swagger/index.html
    ```

## API

Документация API доступна через Swagger UI по адресу [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).

### Эндпоинты

- **Получение данных пользователей:**
   - `GET /` - Главная страница с пагинацией.
   - `GET /filter` - Фильтрация пользователей.

- **Управление пользователями:**
   - `POST /create/:Passport` - Добавление нового пользователя.
   - `PUT /update/:Passport` - Обновление данных пользователя.
   - `DELETE /delete/:Passport` - Удаление пользователя.

- **Управление задачами:**
   - `POST /man/begin/:Passport/:TaskName` - Начало задачи.
   - `POST /man/end/:Passport/:TaskID` - Завершение задачи.

- **Информация о пользователе:**
   - `GET /info/:PassportSerie/:PassportNumber` - Получение информации о пользователе по паспорту.

## Миграции

Миграции базы данных выполняются автоматически при запуске контейнера приложения. Убедитесь, что структура базы данных корректно определяется при старте контейнера.
