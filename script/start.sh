echo Start Server

(cd ../docker-compose && docker-compose --env-file .env up --build -d)