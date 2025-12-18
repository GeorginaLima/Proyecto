# Prueba Técnica 
Este repositorio contiene la implementación de una **prueba técnica backend** basada en una arquitectura de microservicios utilizando **Go**, siguiendo el patrón **CQRS (Command Query Responsibility Segregation)**.

---
El sistema está dividido en dos microservicios independientes:

#Reservation Command MS

Responsable de **modificar el estado** del sistema.

**Endpoints:**

* `POST /create-reservation` → Crear una reserva
* `PUT /update-reservation/{id}` → Actualizar una reserva
* `DELETE /delete-reservation/{id}` → Eliminar una reserva

**Características:**

* Autenticación mediante **JWT**
* Manejo de datos en memoria (para fines de la prueba)
* Arquitectura en capas: `handler`, `service`, `model`

---

# Reservation Query MS

Responsable de **consultar información**.

**Endpoints:**

* `GET /reservation` → Retorna las reservas en formato JSON

**Características:**

* Servicio de solo lectura
* Respuesta en JSON
* Separación clara respecto al Command MS (enfoque CQRS)

---

#Seguridad – JWT

El Command MS está protegido mediante **JSON Web Tokens (JWT)**.
Las solicitudes deben incluir el header:

```
Authorization: Bearer <token>
```

El token es validado mediante firma **HS256** usando una clave secreta definida en el servicio.

---


#Ejecución del proyecto

### Command MS

```bash
cd reservation-command-ms
go run ./cmd
```

### Query MS

```bash
cd reservation-query-ms
go run ./cmd
```

---

*Ejemplo de Request*

```bash
curl -X POST http://localhost:8081/create-reservation \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <TOKEN>" \
-d '{
  "customer_name": "Georgina",
  "date": "2025-12-20T00:00:00Z"
}'
`

