package data

type Permission struct {
	Id   int64  `json:"id"`
	Name string `json:"name" validate:"required"`
}

func GetPermissions(id int64) ([]Permission, error) {
	rows, err := db.Query("SELECT permissionsr FROM users WHERE id = $1", id)
	defer rows.Close()
	if err != nil {
		return []Permission{}, err
	}
	p := []Permission{}
	for rows.Next() {
		permission := Permission{}
		rows.Scan(&permission.Id, &permission.Name)
		p = append(p, permission)
	}

	return p, nil
}
