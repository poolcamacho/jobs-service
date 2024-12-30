
# Auth Service

**Auth Service** es un servicio de autenticación construido en Go, que incluye funcionalidades como registro de usuarios, inicio de sesión con JWT y un endpoint para verificar el estado del servicio.

## Estructura del Proyecto

```plaintext
auth-service/
├── .github/              # Configuración para GitHub Actions
├── cmd/
│   └── main.go           # Punto de entrada principal de la aplicación
├── docs/                 # Documentación Swagger generada
├── internal/
│   ├── domain/           # Definiciones de modelos y estructuras
│   ├── repository/       # Interacción con la base de datos
│   ├── service/          # Lógica de negocio
│   └── transport/        # Handlers de HTTP (controladores)
├── pkg/
│   ├── config/           # Configuración de la aplicación
│   ├── db/               # Conexión a la base de datos
│   ├── jwt/              # Utilidades para manejo de JWT
│   ├── logger/           # Configuración de logging
│   └── utils/            # Funciones utilitarias (e.g., hashing de contraseñas)
├── .env                  # Variables de entorno (no incluir en producción)
├── .gitignore            # Archivos y carpetas ignoradas por Git
├── coverage.out          # Archivo de cobertura de pruebas
├── Dockerfile            # Archivo para construir la imagen de Docker
├── go.mod                # Gestión de dependencias de Go
├── Makefile              # Tareas comunes (compilación, pruebas, Swagger, etc.)
```

---

## Requisitos Previos

- **Go** 1.20 o superior
- **Docker** (opcional, para contenedores)
- **MySQL** como base de datos

---

## Cómo Probar en Local

### 1. Clonar el repositorio

```bash
git clone https://github.com/poolcamacho/auth-service.git
cd auth-service
```

### 2. Configurar las variables de entorno

Crea un archivo `.env` en la raíz del proyecto con las siguientes variables:

```env
DATABASE_URL=admin_db:password@tcp(localhost:3306)/talent_management_db
JWT_SECRET_KEY=tu-secreto-jwt
PORT=3000
```

### 3. Ejecutar la aplicación localmente

Ejecuta el siguiente comando para iniciar el servicio:

```bash
make run
```

Accede al servicio en `http://localhost:3000`.

### 4. Generar la documentación Swagger

```bash
make swagger
```

La documentación estará disponible en `http://localhost:3000/swagger/index.html`.

### 5. Ejecutar pruebas

```bash
make test
```

Esto generará un archivo `coverage.out` con la cobertura de pruebas.

---

## Uso de Docker

### 1. Construir la imagen de Docker

```bash
docker build -t auth-service .
```

### 2. Ejecutar el contenedor

```bash
docker run -d --name auth-service   -e DATABASE_URL=admin_db:password@tcp(localhost:3306)/auth_service_db   -e JWT_SECRET_KEY=tu-secreto-jwt   -p 3000:3000 auth-service
```

El servicio estará disponible en `http://localhost:3000`.

---

## Endpoints

### 1. **Health Check**

**Descripción**: Verifica el estado del servicio.

**Endpoint**: `GET /health`

**Ejemplo de respuesta**:

```json
{
  "status": "healthy"
}
```

---

### 2. **Registro de Usuario**

**Descripción**: Registra un nuevo usuario.

**Endpoint**: `POST /register`

**Cuerpo de la Solicitud**:

```json
{
  "username": "testuser",
  "email": "testuser@example.com",
  "password": "password123"
}
```

**Ejemplo de Respuesta Exitosa**:

```json
{
  "message": "user registered successfully"
}
```

---

### 3. **Inicio de Sesión**

**Descripción**: Autentica un usuario y genera un token JWT.

**Endpoint**: `POST /login`

**Cuerpo de la Solicitud**:

```json
{
  "email": "testuser@example.com",
  "password": "password123"
}
```

**Ejemplo de Respuesta Exitosa**:

```json
{
  "message": "login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Ejemplo de Respuesta de Error**:

```json
{
  "error": "invalid credentials"
}
```

---

