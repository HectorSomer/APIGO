package infraestructure

import (
	"fmt"
	"log"
	core "api-hexagonal/src/config"
	"api-hexagonal/src/sells/domain/entities"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
    conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}
func (mysql * MySQL) CreateSell(sell entities.Sell) (*entities.Sell, error) {
	query := "INSERT INTO sell (concept, total_price, date) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, sell.Concept, sell.Total_Price, sell.Date)
	if err != nil {
		fmt.Println(err)
		//solo retornar el error
		return nil, err
	}

	if result != nil {
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 1 {
			lastInsertID, err := result.LastInsertId()
            if err != nil {
                fmt.Println(err)
                return nil,err
            }
            sell.ID = int32(lastInsertID)
		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Ha habido un error en la consulta (ningún resultado).")
	}
	return &sell, nil
}

func (mysql *MySQL) GetAllSells() (*[]entities.Sell, error) {
	query := "SELECT * FROM sell"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sells []entities.Sell
	//.Next() se usa para iterar sobre los resultados de una consulta SQL fila por fila.
	for rows.Next() {
		var id int32
		var concept string
		var total_price float32
		var date string
		if err := rows.Scan(&id, &concept, &total_price, &date); err != nil {
			fmt.Println("error al escanear la fila: %w", err)
		}
		sell := entities.Sell{
			ID:          id,
            Concept:     concept,
            Total_Price:  total_price,
			Date:        date,
		}
		sells = append(sells, sell)
		
	}
	return &sells, nil
}

func (mysql *MySQL) EditSell(id int, sell entities.UpdatedSell) (*entities.UpdatedSell,error) {
	query := "UPDATE sell SET concept=?, date=?, total_price=? WHERE id=?"
    result, err := mysql.conn.ExecutePreparedQuery(query, sell.Concept, sell.Date, sell.TotalPrice, id)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
	if result != nil {
		rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 1 {
            fmt.Printf("Venta actualizada con éxito.\n")
        } else {
            fmt.Printf("Ninguna venta fue actualizada, intente de nuevo, por favor.\n")
        }
    } else {
        fmt.Printf("Ha habido un error en la consulta (ningún resultado).\n")
    }
    return &sell, nil
}

func (mysql *MySQL) DeleteSell(id int) (bool, error) {
	query := "DELETE FROM sell WHERE id=?"
    result, err := mysql.conn.ExecutePreparedQuery(query, id)
    if err != nil {
        fmt.Println(err)
        return false, err
    }
    if result != nil {
        rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 1 {
            fmt.Printf("Venta eliminada con éxito.\n")
            return true, nil
        } else {
            fmt.Printf("Ninguna venta fue eliminada, intente de nuevo, por favor.\n")
            return false, nil
        }
    } else {
        fmt.Printf("Ha habido un error en la consulta (ningún resultado).\n")
        return false, nil
    }
}