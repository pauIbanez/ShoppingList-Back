package handlers

import "shoppinglist/src/data"

func CreateItemHandler() *ItemHandlers {
	return &ItemHandlers{
		store: map[string]data.Item{
			"1655570749194813500": {
				Id:       "1655570749194813500",
				Name:     "Carrots",
				Quantity: 10,
        Checked: false,
			},
      "1655570749194813600": {
				Id:       "1655570749194813600",
				Name:     "Bananas",
				Quantity: 6,
        Checked: false,
			},
		},
	}
}



