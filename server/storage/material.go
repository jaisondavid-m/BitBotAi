package storage

import (
	"library/config"
)


func SaveMaterial(text string) error {
	_,err:=config.DB.Exec("INSERT INTO study_material (content) VALUES (?)",text)
	return err
}

func GetMaterial() ([]string,error){
	rows,err := config.DB.Query("SELECT content FROM study_material")

	if err!=nil{
		return nil,err
	}

	defer rows.Close()

	var materials []string
	for rows.Next() {
		var m string
		if err := rows.Scan(&m); err != nil {
			return nil, err
		}
		materials = append(materials, m)
	}
	return materials,err
}