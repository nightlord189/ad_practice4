package db

import (
	"errors"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/nightLord189/ad_practice4/model"
	"math/big"
)

func Create(session *gocql.Session, item *model.Almobglu) {
	err := session.Query(`INSERT INTO almoblgu1 (id, dolobimp, importk, importr, namek, namer, obimport, period, predperiod) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		item.ID, item.Dolobimp, item.Importk, item.Importr, item.Namek, item.Namer, item.Obimport, item.Period, item.Predperiod).Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Read(session *gocql.Session, id int) (*model.Almobglu, error) {
	m := map[string]interface{}{}
	q := `
		SELECT * FROM almoblgu1
			WHERE id = ?
		LIMIT 1
    	`
	itr := session.Query(q, id).Consistency(gocql.One).Iter()
	for itr.MapScan(m) {
		item := &model.Almobglu{}
		bigID := m["id"].(*big.Int)
		item.ID = int(bigID.Int64())
		item.Dolobimp = m["dolobimp"].(string)
		item.Importk = m["importk"].(string)
		item.Importr = m["importr"].(string)
		item.Namek = m["namek"].(string)
		item.Namer = m["namer"].(string)
		item.Obimport = m["obimport"].(string)
		item.Period = m["period"].(string)
		item.Predperiod = m["predperiod"].(string)
		return item, nil
	}

	return nil, errors.New("not found")
}

func Update(session *gocql.Session, id int, namer string) error {
	q := `
        	UPDATE almoblgu1
		SET namer = ?
		WHERE id = ?
    	`
	return session.Query(q, namer, id).Exec()
}

func Delete(session *gocql.Session, id int) error {
	q := `DELETE FROM almoblgu1
		WHERE id = ?
    	`
	return session.Query(q, id).Exec()
}
