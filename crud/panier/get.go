package crud

import (
	"encoding/json"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"

	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	db := utils.DbConnect()

	//execute the sql query and check errors
	row := db.QueryRow("SELECT product.id, product.name, product.price, product.description, picture.url FROM product INNER JOIN product_picture ON product_picture.id = product.id INNER JOIN picture ON picture.id = product_picture.id WHERE product.is_alive = true AND product.pending_validation = false AND product.id = ? LIMIT 1", 1)

	//parse the query
	//global vars
	var name, picture, description string
	var id int
	var price float64

	//retrieve the values and check errors
	err := row.Scan(&id, &name, &price, &description, &picture)
	utils.CheckErr(err)

	//add the values to the response
	// bestRated = append(bestRated, structures.BestRatedProduct{
	// 	Id_product: id,
	// 	Name:       name,
	// 	Price:      price,
	// 	Picture:    picture,
	// })

	//create the json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structures.Cart{
		AllItems: []structures.CartItem{
			{
				Id_product:  id,
				Name:        name,
				Price:       price,
				Description: description,
				Picture:     picture,
			},
		},
	})
}