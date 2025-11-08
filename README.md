#  OrderFlow – Sistema de Pedidos Distribuido con Observabilidad

![Docker](https://img.shields.io/badge/Docker-✓-2496ED?style=flat&logo=docker&logoColor=white)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Replica%20Set-336791?style=flat&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-Cache%20%26%20Queue-DC382D?style=flat&logo=redis&logoColor=white)
![Grafana](https://img.shields.io/badge/Observability-Grafana%20%7C%20Prometheus-F46800?style=flat&logo=grafana&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green)

---

##  Descripción

**OrderFlow** es un sistema de pedidos **distribuido y escalable**, diseñado para simular un entorno de microservicios backend con:
- Replicación de base de datos (lecturas distribuidas)
- Cache y colas con Redis
- Observabilidad completa (métricas, logs, traces)
- Escalabilidad horizontal mediante contenedores Docker

> ⚙️ Este proyecto no es multi-tenant (por ahora), pero está diseñado para crecer hacia ello en el futuro.

---

## ️ Arquitectura General

```plaintext
+----------------------------------------------------------+
|                        Nodo Principal                    |
|                                                          |
|  +---------------------+     +-------------------------+ |
|  |  Backend Services   |     |   PostgreSQL Cluster    | |
|  |---------------------|     |-------------------------| |
|  |  [Docker] Backend 1 |     |  DB Main               | |
|  |  [Docker] Backend 2 |     |  Replica 1             | |
|  |  [Docker] Backend 3 |     |  Replica 2             | |
|  +---------------------+     +-------------------------+ |
|                                                          |
|                   +-------------------+                  |
|                   |      Redis        |                  |
|                   +-------------------+                  |
+----------------------------------------------------------+
