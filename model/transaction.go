package model

type (
	Customer struct {
		ID           int    `json:"id"`
		CustomerName string `json:"customer_name"`
	}

	CustomerAddress struct {
		ID         int      `json:"id"`
		CustomerID Customer `json:"customer_id"`
	}

	Product struct {
		ID    int    `json:"id"`
		Name  string `josn:"name"`
		Price int    `json:"price"`
	}

	PaymentMethod struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		IsActive int    `json:"is_active"`
	}

	Transaction struct {
		ID                int `json:"id"`
		CustomerAddressID int `json:"customer_address_id" db:"customer_address_id"`
		Product           int `json:"product_id" db:"product_id"`
		PaymentMethod     int `json:"payment_method_id" db:"payment_method_id"`
		ProductQty        int `json:"product_qty" db:"product_qty"`
	}
)
