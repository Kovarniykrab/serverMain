# testServer

```
Первый тестовый сервер 
```

## libs
```
https://github.com/swaggo/swag  swagger

https://github.com/pressly/goose  for migrate
```

## installation

```

```

##  Документация
Документация описана в swagger и доступна по ссылке:
[swagger](https://github.com/Kovarniykrab/serverMain/blob/main/docs/swagger.json)

##  Пример конфигурационного файла .env.
???
##  Пример скрипта для сервера.
???
###  Dockerfile
Готовый образ доступен в [Docker](https://github.com/Kovarniykrab/serverMain/blob/main/Dockerfile)

**Основные модули:**
- Регистрация
- Авторизация
- JWT-аутентификация
- Cистема уведомлений(WebSocket)
- Email-рассылки (Unisender)
- Отправка файлов

**Интеграции:**
- PostgreSQL - основное хранилище


## Деплой и CI/CD

Система использует GitLab CI/CD для автоматизированной сборки и деплоя. Процесс делится на два этапа:

# 1. Сборка образа (CI)
При каждом пуше в ветки `main`  GitHub Runner:
- Авторизуется в Docker Registry
- Собирает Docker-образ с тегом, соответствующим ветке 
- Пушит образ в GitLab Container Registry

# 2. Деплой на сервер
- На сервере по cron запускается скрипт pull.sh :
- Проверяет обновления в registry
- Скачивает новую версию образа
- Перезапускает контейнер.

## Миграции
Используется goose:
Автоматическое применение при старте
Миграции в resources/store/psql/migrations/