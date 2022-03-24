package CRUD

import (
	databaseTools "inprinte/backend/database"
	"inprinte/backend/utils"
)

func updateAddress(street, city, state, country, zip_code string, id_user int) {
	db := databaseTools.DbConnect()
	_, err := db.Exec(`UPDATE address SET street = ?, city = ?, state = ?, country = ?, zip_code = ? WHERE id = (SELECT user.id_address FROM user WHERE user.id = ?)`, street, city, state, country, zip_code, id_user)
	utils.CheckErr(err)
}

func updateName(first_name, last_name string, id_user int) {
	db := databaseTools.DbConnect()
	_, err := db.Exec(`UPDATE user SET first_name = ?, last_name = ? WHERE user.id = ?`, first_name, last_name, id_user)
	utils.CheckErr(err)
}

func updateEmail(email string, id_user int) {
	db := databaseTools.DbConnect()
	_, err := db.Exec(`UPDATE user SET email = ? WHERE user.id = ?`, email, id_user)
	utils.CheckErr(err)
}

func updatePassword(password string, id_user int) {
	db := databaseTools.DbConnect()
	_, err := db.Exec(`UPDATE user SET password = ? WHERE user.id = ?`, password, id_user)
	utils.CheckErr(err)
}

func updatePhone(phone string, id_user int) {
	db := databaseTools.DbConnect()
	_, err := db.Exec(`UPDATE user SET phone = ? WHERE user.id = ?`, phone, id_user)
	utils.CheckErr(err)
}

func disableAccount(id_user int) {
	db := databaseTools.DbConnect()
	_, err := db.Exec(`UPDATE user SET is_alive = false WHERE user.id = ?`, id_user)
	utils.CheckErr(err)
}
