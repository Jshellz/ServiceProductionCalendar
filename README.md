# Сервис производственного календаря (Service production calendar)

## Для запуска сервиса календаря надо:
1. Установить Go -> https://go.dev/doc/install
2. Установить GoLang -> https://www.jetbrains.com/go/
3. Установить демон запуска -> https://github.com/Jshellz/CompileDaemon
4. Установить cloudDB -> https://tableplus.com/ и https://console.neon.tech/
5. Установить Postman -> https://www.postman.com/
6. В рабочей директории в терминате вставте -> CompileDaemon -command="./ServiceProductionCalendar", запуститься локально
   ![Вот такое отображать должно](/image/photo_comp_demon.png)
7. После откройте устновленый https://tableplus.com/, и вставте DB_URL который находиться в .env либо сосздайте свой тут -> https://console.neon.tech/ и поменяйте DB_URL в .env на свой
8. Запросы создавайте в https://www.postman.com/

