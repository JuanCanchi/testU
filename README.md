# Plataforma de Microblogging - Ejercicio Técnico

Este proyecto implementa una versión simplificada de una plataforma de microblogging, con características básicas similares a Twitter, como publicar tweets, seguir usuarios y ver un timeline.

## Funcionalidades

1. **Publicación de Tweets**  
   Los usuarios pueden publicar mensajes cortos de hasta 280 caracteres.

2. **Seguir Usuarios**  
   Los usuarios pueden seguir a otros para acceder a sus publicaciones en el timeline.

3. **Timeline**  
   Muestra los tweets más recientes de los usuarios seguidos.

---

## Tecnologías Utilizadas

- **Lenguaje**: Go (Golang)
- **Base de Datos**: MongoDB
- **Arquitectura**: Hexagonal / Clean Architecture
- **Contenedores**: Docker y Docker Compose

---

## Levantar el Proyecto

1. **Tener instalado Docker**  
   Es necesario tener instalado Docker para poder levantar el proyecto

2. **Inicializar el Proyecto**  
   Con Docker instalado y el proyecto descargado ubicarse el la raiz del mismo y ejecutar ```sudo docker compose up --build -d```
   con eso levantamos proyecto

3. **Crear Uusarios de Prueba**  
   Buscamos el id container de la app ```docker ps``` 
   Ingresamo al container ```docker exec -it idContainer bash```
   Corremos el binario usersCreate ```cmd/users/createUsers```



## Probar endpoints
Puedes usar el siguiente comando `curl` para crear un nuevo tweet:


   ```bash
   curl -X POST http://localhost:8080/tweets \
   -H "Content-Type: application/json" \
   -d '{
     "user_id": "user2",
     "content": "Este es un tweet de prueba3"
   }'
   ```

Seguir a un nuevo usuario:
   ```bash 
   curl -X POST http://localhost:8080/follow \
   -H "Content-Type: application/json" \
   -d '{
     "user_id": "user1",
     "follow_id": "user2"
   }'
   ```
Obtener el timeline del usuario:
 ```bash
 curl -X GET "http://localhost:8080/timeline?user_id=user1"
 ```

