# Golang images gallery

## Description (RU)

Данный репозиторий содержит исходный код онлайн-галереи, созданной в ходе курса *[Web Delvelopment with Go by Jon Calhoun](https://www.usegolang.com/)*

На момент написания данной заметки (08 May 2021) с работой галлереи можно ознакомиться по [ссылке](https://138.68.132.2/galleries/6) (при предупреждении безопасности всё равно можно перейти)

База данных - **PostgreSQL** с использованием **[GORM](https://gorm.io/)**

Внешний вид - **[Bootstrap](https://getbootstrap.com/)**

Формирование .html страниц - стандартная библиотека **html/template**

Cookie - стандартная библиотека **net/http**

Router, CSRF-protection, POST form parsing - **[Gorilla](https://www.gorillatoolkit.org/)**

Запущена на Digital Ocean с web-сервером **[Caddy](https://caddyserver.com/)**

## Launch

I hope that works :)

```bash

# postgres is running on localhost:5432 with 'postgres' password and 'learnwebdev' DB 

go mod tidy

go run .

```
