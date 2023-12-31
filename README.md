# Сервис, предоставляющий API по созданию сокращённых ссылок.

`config.yaml` - минимальный файл конфигурации для базы данных и приложения

## Запуск приложения

**Миграции в ручную накатывать не нужно, все происходит автоматически.**

**Сервис можно запустить командой:**

```bash
docker compose up 
```

**Сервис реализован через gRPC. Контракт сервиса можно найти в `./protos/proto/url_shortener/url_shortener.proto`**

**В `Dockerfile` можно заменить в команду `CMD ["./main", "postgres"]` на `CMD ["./main", "memory"]`. В таком режиме
сервис будет сохранять данные в памяти только компьютера, и при остановке программы все они потеряются.**

## Примечания:

- **Если сервис оставить работать на очень долгое время, в базе данных может накопиться слишком много данны. Предложил
  бы решить эту проблемы следующим образом - добавить в таблицу столбец `expired_at`, который бы показывал, через какое
  времени после последнего обращения к ссылке, если к ней никто больше не обратится, ее можно будет удалить из базы. За
  удаление записей из базы бы отвечал отдельный воркер, который бы работал в фоне и раз в какое-то время обходил бы всю
  таблицу.**