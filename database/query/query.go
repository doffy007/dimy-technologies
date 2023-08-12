package query

var (
	CreateTransaction = `
	INSERT INTO
	transaction (
		customer_address_id, 
		product_id, 
		payment_method_id, 
		product_qty
	)VALUES(
		:customer_address_id, 
		:product_id, 
		:payment_method_id, 
		:product_qty
		);
  `
)
