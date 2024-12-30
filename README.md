
# Jobs Service

**Jobs Service** es un servicio construido en Go que permite gestionar información sobre ofertas de trabajo, incluyendo funcionalidades para listar y crear nuevos trabajos. También incluye un endpoint para verificar el estado del servicio.

## Estructura del Proyecto

```plaintext
jobs-service/
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
│   ├── logger/           # Configuración de logging
│   └── utils/            # Funciones utilitarias
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
git clone https://github.com/poolcamacho/jobs-service.git
cd jobs-service
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
docker build -t jobs-service .
```

### 2. Ejecutar el contenedor

```bash
docker run -d --name jobs-service   -e DATABASE_URL=admin_db:password@tcp(localhost:3306)/auth_service_db   -e JWT_SECRET_KEY=tu-secreto-jwt   -p 3000:3000 jobs-service
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

### 2. **Registro de Trabajo**

**Descripción**: Agrega un nuevo trabajo.

**Endpoint**: `POST /jobs`

**Cuerpo de la Solicitud**:

```json
{
  "title": "Backend Developer",
  "description": "Develop backend services for applications.",
  "salary_range": "4000-6000"
}
```

**Ejemplo de Respuesta Exitosa**:

```json
{
  "message": "job created successfully"
}
```

**Ejemplo de Respuesta de Error**:

```json
{
  "error": "title and description are required"
}
```

---

### 3. **Listar Trabajos**

**Descripción**: Recupera una lista de todos los trabajos disponibles.

**Endpoint**: `GET /jobs`

**Ejemplo de Respuesta Exitosa**:

```json
[
  {
    "id": 1,
    "title": "Software Engineer",
    "description": "Develop and maintain software.",
    "salary_range": "4000-6000",
    "created_at": "2024-12-30T02:00:00Z",
    "updated_at": "2024-12-30T02:00:00Z"
  },
  {
    "id": 2,
    "title": "Product Manager",
    "description": "Oversee product lifecycle.",
    "salary_range": "5000-7000",
    "created_at": "2024-12-30T02:10:00Z",
    "updated_at": "2024-12-30T02:10:00Z"
  }
]
```
---

