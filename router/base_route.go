package router

func (r router) BaseRouter() {
	transactionHandler := r.handler.TransactionHandler()

	v1Prefix := r.route.Group("/api")
	{
		transactionPrefix := v1Prefix.Group("transaction")
		{
			transactionPrefix.POST("/create-transaction", transactionHandler.TransactionRegister)
		}
	}
}
