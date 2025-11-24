package storage

import (
	"library/config"
)


func SaveMaterial(text string) error {
	_,err:=config.DB.Exec("INSERT INTO study_material (content) VALUES (?)",text)
	return err
}

func GetMaterial() (string,error){
	row := config.DB.QueryRow("SELECT content FROM study_material ORDER BY id DESC LIMIT 1")

	var material string
	err := row.Scan(&material)
	return material,err
}