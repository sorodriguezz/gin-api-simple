# Gin API Simple

API desarrollada en Go utilizando el framework Gin. Contiene un CRUD de usuarios con persistencia de datos en PostgreSQL.

## Iniciar el Proyecto

1. Asegúrate de crear el archivo `.env` en la raíz del proyecto con las variables de conexión a PostgreSQL. Un ejemplo de `.env`:

    ```
    DATABASE_URL=host=localhost user=usuario password=contraseña dbname=nombre_base_datos port=5432 sslmode=disable
    ```

2. Luego, inicia el servicio de PostgreSQL usando Docker:

    ```bash
    docker compose up -d
    ```

3. Instala las dependencias del proyecto ejecutando:

    ```bash
    go mod tidy
    ```

4. Ejecuta la aplicación en la raíz del proyecto con:

    ```bash
    go run main.go
    ```

La API estará disponible en `http://localhost:8080`.

## cURLs del Proyecto

### Obtener Usuarios
```bash
curl --location 'http://localhost:8080/usuarios'
```

### Crear Usuario
```bash
curl --location 'http://localhost:8080/usuarios' \
--header 'Content-Type: application/json' \
--data-raw '{
    "nombre": "nombre_prueba",
    "email": "test@test.cl"
}'
```

### Obtener Usuario por ID
```bash
curl --location 'http://localhost:8080/usuarios/1'
```

### Modificar Usuario por ID
```bash
curl --location --request PUT 'http://localhost:8080/usuarios/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "nombre": "nombre_actualizado",
    "email": "actualizado@test.cl"
}'
```

### Eliminar Usuario por ID
```bash
curl --location --request DELETE 'http://localhost:8080/usuarios/1'
```
