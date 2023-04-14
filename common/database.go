package common

//连接sql sever
import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github/Intelligent-scheduling-system-back-end/model"
	"log"
	"time"
)

func ReadDailySchedule(dbPool *ConnPool) ([]byte, error) {
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
		var planid, eid, epos, w1 string
		if err := rows.Scan(&planid, &eid, &epos, &w1); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		schedule := map[string]string{
			"Planid": planid,
			"Eid":    eid,
			"Epos":   epos,
			"W1":     w1,
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

func ReadWeeklySchedule(dbPool *ConnPool) ([]string, error) {
	db, err := dbPool.Get()
	if err != nil {
		return nil, err
	}
	defer dbPool.Release(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// SQL查询语句，注意指定需要查询的列
	var tsql = `SELECT Planid, Eid, Epos, W1, W2, W3, W4, W5, W6, W7 FROM PaiBan;`

	// 准备SQL查询语句
	stmt, err := db.PrepareContext(ctx, tsql)
	if err != nil {
		log.Printf("\nPrepare failed:%T %+v\n", err, err)
		return nil, err
	}
	defer stmt.Close()

	// 执行SQL查询
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Println("Error executing SQL query:", err)
		return nil, err
	}
	defer rows.Close()

	// 将查询结果转换为JSON格式
	var jsonData []string
	for rows.Next() {
		var planid, eid, epos, w1, w2, w3, w4, w5, w6, w7 string
		if err := rows.Scan(&planid, &eid, &epos, &w1, &w2, &w3, &w4, &w5, &w6, &w7); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		schedule := map[string]string{
			"Planid": planid,
			"Eid":    eid,
			"Epos":   epos,
			"W1":     w1,
			"W2":     w2,
			"W3":     w3,
			"W4":     w4,
			"W5":     w5,
			"W6":     w6,
			"W7":     w7,
		}
		jsonRow, err := json.Marshal(schedule)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			return nil, err
		}
		var jsonRow0 = string(jsonRow)
		jsonData = append(jsonData, jsonRow0)
	}
	return jsonData, nil
}

func InsertEmployee(dbPool *ConnPool, employee model.Employee, emplike model.Emplike) (string, error) {
	db, err := dbPool.Get()
	if err != nil {
		return "nil", err
	}
	defer dbPool.Release(db)
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	var tsql1 = fmt.Sprintf("INSERT INTO [dbo].[Employee] VALUES ('%s', '%s', '%s', '%s', '%s');",
		employee.Eid, employee.Ename, employee.Eemail, employee.Epos, employee.Shid)
	_, err = db.Exec(tsql1)
	if err != nil {
		fmt.Println("Error reading rows in Employee: ", err.Error())
		return "nil", err
	}

	var tsql2 = fmt.Sprintf("INSERT INTO [dbo].[Emplike] VALUES ('%s', '%s', '%s');",
		emplike.Etype, emplike.Eid, emplike.Elike)
	_, err = db.Exec(tsql2)
	if err != nil {
		fmt.Println("Error reading rows in emplike: ", err.Error())
		return "nil", err
	}

	return "success", nil
}
