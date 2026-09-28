# gocommerce-tools

Repositório compartilhado contendo **apenas contratos, modelos de dados, migrações de banco e eventos de mensageria** para o ecossistema de microsserviços do **GoCommerce**.

> ⚠️ **Atenção:** Este repositório é uma biblioteca compartilhada de tipos e migrações. Não contém rotas HTTP, APIs, handlers ou lógica de servidor.

---

## 📁 Estrutura do Pacote

```
gocommerce-tools/
├── migrations/             # Scripts SQL de migrações do banco de dados compartilhado
│   ├── 000001_create_users_table.up.sql
│   ├── 000001_create_users_table.down.sql
│   ├── 000002_create_rbac_tables.up.sql
│   └── 000002_create_rbac_tables.down.sql
└── pkg/
    ├── models/             # Entidades puras de banco de dados (User, Role, Permission)
    │   └── user.go
    ├── events/             # Contratos de eventos para filas (RabbitMQ, Kafka, SQS)
    │   └── user_events.go
    └── constants/          # Constantes globais (Departamentos, Roles, Nomes de Filas)
        └── constants.go
```

---

## 🚀 Como importar em seus microsserviços

```go
import (
    "github.com/matheusbanqueiro/gocommerce-tools/pkg/models"
    "github.com/matheusbanqueiro/gocommerce-tools/pkg/events"
    "github.com/matheusbanqueiro/gocommerce-tools/pkg/constants"
)
```

Para trabalhar localmente sem precisar dar push no GitHub durante o desenvolvimento:
```powershell
go mod edit -replace github.com/matheusbanqueiro/gocommerce-tools=../gocommerce-tools
```
