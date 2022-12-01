package acloud

import (
	"database/sql"
	"errors"
	"fmt"
	"goiq/iq"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	fmt.Println("NewSQLiteRepository")
	return &SQLiteRepository{
		db: db,
	}
}


/*
Title   string
	Date    string
	Author  string
	Content string
	Budget  string
	Label   string
*/
func (r *SQLiteRepository) CreateTable() error {
	fmt.Println("Migrate")
	query := `
	CREATE TABLE IF NOT EXISTS Requests(
		id 			INTEGER PRIMARY KEY AUTOINCREMENT,
		Author 		TEXT NOT NULL UNIQUE,
		Date 		TEXT NOT NULL,
		Content		TEXT NOT NULL,
		Budget 		TEXT NOT NULL,
		Label 		TEXT NOT NULL
	)
	`
	fmt.Println("r.db.Exec(query) :")
	_, err := r.db.Exec(query)
	fmt.Println("err  :", err)
	return err
}

func (r *SQLiteRepository) CreateRecord(request iq.IqRequest) (*iq.IqRequest, error) {
	fmt.Println("CreateRecord : ", request)
	fmt.Println("r: ", r)
	res, err := r.db.Exec("INSERT INTO Requests(Author 	
		,Date 	
		,Content	
		,Budget 	
		,Label 	) values(?,?,?,?,?)",
		creds.User, creds.Password, creds.URL, creds.KeyID, creds.AccessKey)
	fmt.Println("res : ", res)
	cli.PrintIfErr(err)
	// fmt.Println("err : ", err)
	// if err != nil {
	// 	var sqliteErr sqlite3.Error
	// 	if errors.As(err, &sqliteErr) {
	// 		if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
	// 			return nil, ErrDuplicate
	// 		}
	// 	}
	// 	return nil, err
	// }

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	creds.ID = id
	return &creds, nil
}

func (r *SQLiteRepository) Last() (*SandboxCredential, error) {

	row := r.db.QueryRow("SELECT * FROM SandboxCredentials ORDER BY User DESC LIMIT 1;")

	var SandboxCredential SandboxCredential
	if err := row.Scan(&SandboxCredential.ID, &SandboxCredential.User, &SandboxCredential.Password, &SandboxCredential.URL, &SandboxCredential.KeyID, &SandboxCredential.AccessKey); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &SandboxCredential, nil
}

func (r *SQLiteRepository) All() ([]SandboxCredential, error) {
	rows, err := r.db.Query("SELECT * FROM SandboxCredentials")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []SandboxCredential
	for rows.Next() {
		var SandboxCredential SandboxCredential
		if err := rows.Scan(&SandboxCredential.ID, &SandboxCredential.User, &SandboxCredential.Password, &SandboxCredential.URL, &SandboxCredential.KeyID, &SandboxCredential.AccessKey); err != nil {
			return nil, err
		}
		all = append(all, SandboxCredential)
	}
	return all, nil
}

func (r *SQLiteRepository) GetByName(user string) (*SandboxCredential, error) {
	row := r.db.QueryRow("SELECT * FROM SandboxCredentials WHERE User = ?", user)

	var SandboxCredential SandboxCredential
	if err := row.Scan(&SandboxCredential.ID, &SandboxCredential.User, &SandboxCredential.Password, &SandboxCredential.URL, &SandboxCredential.KeyID, &SandboxCredential.AccessKey); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &SandboxCredential, nil
}

func (r *SQLiteRepository) Update(id int64, updated SandboxCredential) (*SandboxCredential, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	res, err := r.db.Exec("UPDATE SandboxCredentials SET User=?,Password=?,URL=?,KeyID=?,AccessKey=? WHERE ID = ?", updated.User, updated.Password, updated.URL, updated.KeyID, updated.AccessKey, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *SQLiteRepository) Delete(id int64) error {
	res, err := r.db.Exec("DELETE FROM SandboxCredentials WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
