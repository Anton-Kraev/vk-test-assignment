#!/bin/bash

export $(cat .env | xargs)

goose -dir ./backend/migrations postgres "$DATABASE_URL_MIGRATIONS" status
goose -dir ./backend/migrations postgres "$DATABASE_URL_MIGRATIONS" up
