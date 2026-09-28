package constants

// Departamentos do sistema
const (
	DeptSystem  = "system"
	DeptFinance = "finance"
	DeptStock   = "stock"
	DeptClient  = "client"
)

// Cargos/Roles padrão
const (
	RoleAdmin          = "admin"
	RoleFinanceManager = "finance_manager"
	RoleFinanceAnalyst = "finance_analyst"
	RoleStockManager   = "stock_manager"
	RoleStockist       = "stockist"
	RoleClient         = "client"
)

// Filas SQS
const (
	QueuePasswordReset = "password-reset-queue"
)
