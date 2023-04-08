package common

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"time"
)

func ReadDailySchedule1(dbPool *ConnPool) ([]byte, error) {
	db, err := dbPool.Get()
	if err != nil {
		return nil, err
	}
	defer dbPool.Release(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// sql查询语句，注意指定需要查询的列

	tsql := fmt.Sprintf(
		"SELECT Planid, Eid, Epos, W1 FROM [dbo].[PaiBan];",
	)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println("Error reading rows: ", err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err.Error())
		}
	}(rows)

	var jsonData []byte
	for rows.Next() {
		var planid, eid, epos, W1 string
		if err := rows.Scan(&planid, &eid, &epos, &W1); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		schedule := map[string]string{
			"Planid": planid,
			"Eid":    eid,
			"Epos":   epos,
			"W1":     W1,
		}
		jsonRow, err := json.Marshal(schedule)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			return nil, err
		}
		jsonData = append(jsonData, jsonRow...)
	}
	return jsonData, nil
}

func ReadDailySchedule2(dbPool *ConnPool) ([]byte, error) {
	db, err := dbPool.Get()
	if err != nil {
		return nil, err
	}
	defer dbPool.Release(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// sql查询语句，注意指定需要查询的列

	tsql := fmt.Sprintf(
		"SELECT Planid, Eid, Epos, W2 FROM [dbo].[PaiBan];",
	)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println("Error reading rows: ", err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err.Error())
		}
	}(rows)

	var jsonData []byte
	for rows.Next() {
		var planid, eid, epos, W2 string
		if err := rows.Scan(&planid, &eid, &epos, &W2); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		schedule := map[string]string{
			"Planid": planid,
			"Eid":    eid,
			"Epos":   epos,
			"W2":     W2,
		}
		jsonRow, err := json.Marshal(schedule)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			return nil, err
		}
		jsonData = append(jsonData, jsonRow...)
	}
	return jsonData, nil
}

func ReadDailySchedule3(dbPool *ConnPool) ([]byte, error) {
	db, err := dbPool.Get()
	if err != nil {
		return nil, err
	}
	defer dbPool.Release(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// sql查询语句，注意指定需要查询的列

	tsql := fmt.Sprintf(
		"SELECT Planid, Eid, Epos, W3 FROM [dbo].[PaiBan];",
	)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println("Error reading rows: ", err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err.Error())
		}
	}(rows)

	var jsonData []byte
	for rows.Next() {
		var planid, eid, epos, w3 string
		if err := rows.Scan(&planid, &eid, &epos, &w3); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		schedule := map[string]string{
			"Planid": planid,
			"Eid":    eid,
			"Epos":   epos,
			"W3":     w3,
		}
		jsonRow, err := json.Marshal(schedule)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			return nil, err
		}
		jsonData = append(jsonData, jsonRow...)
	}
	return jsonData, nil
}

func ReadDailySchedule4(dbPool *ConnPool) ([]byte, error) {
	db, err := dbPool.Get()
	if err != nil {
		return nil, err
	}
	defer dbPool.Release(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// sql查询语句，注意指定需要查询的列

	tsql := fmt.Sprintf(
		"SELECT Planid, Eid, Epos, W4 FROM [dbo].[PaiBan];",
	)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println("Error reading rows: ", err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err.Error())
		}
	}(rows)

	var jsonData []byte
	for rows.Next() {
		var planid, eid, epos, W4 string
		if err := rows.Scan(&planid, &eid, &epos, &W4); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		schedule := map[string]string{
			"Planid": planid,
			"Eid":    eid,
			"Epos":   epos,
			"W4":     W4,
		}
		jsonRow, err := json.Marshal(schedule)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			return nil, err
		}
		jsonData = append(jsonData, jsonRow...)
	}
	return jsonData, nil
}

func ReadDailySchedule5(dbPool *ConnPool) ([]byte, error) {
	db, err := dbPool.Get()
	if err != nil {
		return nil, err
	}
	defer dbPool.Release(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// sql查询语句，注意指定需要查询的列

	tsql := fmt.Sprintf(
		"SELECT Planid, Eid, Epos, W5 FROM [dbo].[PaiBan];",
	)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println("Error reading rows: ", err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err.Error())
		}
	}(rows)

	var jsonData []byte
	for rows.Next() {
		var planid, eid, epos, W5 string
		if err := rows.Scan(&planid, &eid, &epos, &W5); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		schedule := map[string]string{
			"Planid": planid,
			"Eid":    eid,
			"Epos":   epos,
			"W5":     W5,
		}
		jsonRow, err := json.Marshal(schedule)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			return nil, err
		}
		jsonData = append(jsonData, jsonRow...)
	}
	return jsonData, nil
}

func ReadDailySchedule6(dbPool *ConnPool) ([]byte, error) {
	db, err := dbPool.Get()
	if err != nil {
		return nil, err
	}
	defer dbPool.Release(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// sql查询语句，注意指定需要查询的列

	tsql := fmt.Sprintf(
		"SELECT Planid, Eid, Epos, W6 FROM [dbo].[PaiBan];",
	)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println("Error reading rows: ", err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err.Error())
		}
	}(rows)

	var jsonData []byte
	for rows.Next() {
		var planid, eid, epos, W6 string
		if err := rows.Scan(&planid, &eid, &epos, &W6); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		schedule := map[string]string{
			"Planid": planid,
			"Eid":    eid,
			"Epos":   epos,
			"W6":     W6,
		}
		jsonRow, err := json.Marshal(schedule)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			return nil, err
		}
		jsonData = append(jsonData, jsonRow...)
	}
	return jsonData, nil
}

func ReadDailySchedule7(dbPool *ConnPool) ([]byte, error) {
	db, err := dbPool.Get()
	if err != nil {
		return nil, err
	}
	defer dbPool.Release(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// sql查询语句，注意指定需要查询的列

	tsql := fmt.Sprintf(
		"SELECT Planid, Eid, Epos, W7 FROM [dbo].[PaiBan];",
	)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println("Error reading rows: ", err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err.Error())
		}
	}(rows)

	var jsonData []byte
	for rows.Next() {
		var planid, eid, epos, W7 string
		if err := rows.Scan(&planid, &eid, &epos, &W7); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		schedule := map[string]string{
			"Planid": planid,
			"Eid":    eid,
			"Epos":   epos,
			"W7":     W7,
		}
		jsonRow, err := json.Marshal(schedule)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			return nil, err
		}
		jsonData = append(jsonData, jsonRow...)
	}
	return jsonData, nil
}
