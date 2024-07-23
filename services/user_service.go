package services

import (
	"strings"

	"kplus.com/dto"
	"kplus.com/utils"
)

type UserService struct {
	db utils.Database
}

func (u UserService) GetUser(id string) (*dto.UserDto, error) {
	var data dto.UserDto
	if err := u.db.QueryRow(`
		SELECT u.id, u.name, u.email, u.role, ud.id, ud.full_name, ud.legal_name, ud.place_of_birth, ud.date_of_birth, ud.salary, ud.selfie, ud.selfie_with_national_id, ud.identity_number, ud.national_id_image FROM users u
		JOIN user_details ud ON u.id = ud.user_id
		WHERE id = ?`, id).Scan(
		&data.ID, &data.Email, &data.Role, &data.Details.ID, &data.Details.FullName, &data.Details.LegalName, &data.Details.PlaceOfBirth, &data.Details.DateOfBirth, &data.Details.Salary, &data.Details.Selfie, &data.Details.SelfieWithNationalID, &data.Details.Nik, &data.Details.NationalIdImage,
	); err != nil {
		return nil, err
	}
	return &data, nil
}

func (u UserService) CreateUserDetails(data dto.UserDetailsDto) (int64, error) {
	re, err := u.db.Exec(`
		INSERT INTO user_details
		(user_id, full_name, legal_name, place_of_birth, date_of_birth, salary, selfie, selfie_with_national_id, identity_number, national_id_image)
		VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, data.UserID, data.FullName, data.LegalName, data.PlaceOfBirth, data.DateOfBirth, data.Salary, data.Selfie, data.SelfieWithNationalID, data.Nik, data.NationalIdImage)

	if err != nil {
		return 0, err
	}

	return re.LastInsertId()
}

func (u UserService) UpdateUserDetails(data dto.UserDetailsDto, userID string) error {
	updateSql := `UPDATE user_details SET `
	setClause := []string{}
	args := []interface{}{}

	if data.FullName != nil {
		setClause = append(setClause, "full_name = ?")
		args = append(args, *data.FullName)
	}
	if data.LegalName != nil {
		setClause = append(setClause, "legal_name = ?")
		args = append(args, *data.LegalName)
	}
	if data.PlaceOfBirth != nil {
		setClause = append(setClause, "place_of_birth = ?")
		args = append(args, *data.PlaceOfBirth)
	}
	if data.DateOfBirth != nil {
		setClause = append(setClause, "date_of_birth = ?")
		args = append(args, *data.DateOfBirth)
	}
	if data.Salary != nil {
		setClause = append(setClause, "salary = ?")
		args = append(args, *data.Salary)
	}
	if data.Selfie != nil {
		setClause = append(setClause, "selfie = ?")
		args = append(args, *data.Selfie)
	}
	if data.SelfieWithNationalID != nil {
		setClause = append(setClause, "selfie_with_national_id = ?")
		args = append(args, *data.SelfieWithNationalID)
	}
	if data.Nik != nil {
		setClause = append(setClause, "identity_number = ?")
		args = append(args, *data.Nik)
	}
	if data.NationalIdImage != nil {
		setClause = append(setClause, "national_id_image = ?")
		args = append(args, *data.NationalIdImage)
	}
	updateSql += " " + strings.Join(setClause, ", ") + " WHERE user_id = ?"
	args = append(args, userID)
	if _, err := u.db.Exec(updateSql, args...); err != nil {
		return err
	} else {
		return nil
	}
}

func NewUserService(db utils.Database) UserService {
	return UserService{
		db: db,
	}
}
