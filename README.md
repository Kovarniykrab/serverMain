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
1. развернуть сервера баз данных
REQ: vps/vds на Linux Ubuntu 22  RAM 2 GB  HDD 30 GB
    docker, https://www.docker.com/

2. настраиваем внешние службы 
    a: S3 совместимое хранилище (https://cloud.yandex.ru) для хранения файлов. По тикету происходит запрос на хранилища 
    если там появились файлы происходит выгрузка новых данных и обновления их в бд.
    
    b: личный кабинет https://go1.unisender.ru для почтовых рассылок.

3. развернуть сервер приложений
REQ: vps/vds на Linux Ubuntu 22  RAM 2 GB  HDD 30 GB
    docker  https://www.docker.com/
    настраиваем с помощью dns и https://letsencrypt.org/ru/ https сертификаты сервера
    a: создаем директорию mkdir /var/mayak
    b: загружаем в нее файы:
        .env  - пример файла окружения для приложения там в том числе устанавливаются креды для баз данных
        pull.sh - файл размещенный на целевом сервере служит для автоматизации выкатки новых версий
    c: в script/pull.sh   (скрипт работает только с реджистри гитлаба):
        user= -имя пользователя гитлаб
        token= gpg токен этого пользователя с правами на чтение (glpat........)
    d: настраиваем в кроне сервера запуск скрипта раз в минтуту
    e: вызываем вручную скрипт убеждаемся что процесс появился в docker ps
```

##  Документация
Документация описана в swagger и доступна по ссылке:
[swagger](https://gitlab.com/kovarniykrab/servermain/-/blob/main/docs/swagger.json?ref_type=heads)

##  Пример конфигурационного файла .env.
???
##  Пример скрипта для сервера.
???
###  Dockerfile
Готовый образ доступен в [Docker](https://gitlab.com/kovarniykrab/servermain/-/blob/main/Dockerfile?ref_type=heads)

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
При каждом пуше в ветки `main` или `firstStage` GitLab Runner:
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