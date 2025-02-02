package infraestructure

import (
	"fmt"
	"log"
	core "nombre-del-proyecto/src/config"
	"nombre-del-proyecto/src/products/domain/entities"

	_ "github.com/go-sql-driver/mysql"
	// el _ sirve para aquellas importaciones que no van a ser usadas.
)

type MySQL struct {
	conn *core.Conn_MySQL
}
func NewMySql() *MySQL{
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql * MySQL) CreateProduct(product entities.Product) (*entities.Product, error) {
	query := "INSERT INTO product (name, description, stock, price) VALUES (?, ?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Description, product.Stock, product.Price)
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
            product.ID = int32(lastInsertID)
		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Ha habido un error en la consulta (ningún resultado).")
	}
	return &product, nil
}

func (mysql *MySQL) GetProducts() (*[]entities.Product, error) {
	query := "SELECT * FROM product"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	//.Next() se usa para iterar sobre los resultados de una consulta SQL fila por fila.
	for rows.Next() {
		var id int32
		var name string
		var price float32
		var description string
		var stock int32
		if err := rows.Scan(&id, &name, &description, &stock, &price); err != nil {
			fmt.Println("error al escanear la fila: %w", err)
		}
		product := entities.Product{
			ID:          id,
			Name:        name,
			Price:       price,
			Description: description,
			Stock:       int(stock),
		}
		products = append(products, product)
		
		fmt.Printf("ID: %d, Nombre: %s, Precio: %f, Descripcion: %s, Cantidad: %d \n", id, name, price, description, stock)
	}
	return &products, nil
}

func (mysql *MySQL) EditProduct(id int, product entities.UpdateProduct) (*entities.UpdateProduct,error) {
	query := "UPDATE product SET name=?, description=?, stock=?, price=? WHERE id=?"
    result, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Description, product.Stock, product.Price, id)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
	if result != nil {
		rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 1 {
            fmt.Printf("Producto actualizado con éxito.\n")
        } else {
            fmt.Printf("Ningún producto fue actualizado.\n")
        }
    } else {
        fmt.Printf("Ha habido un error en la consulta (ningún resultado).\n")
    }
    return &product, nil
}

func (mysql *MySQL) DeleteProduct(id int) (bool, error) {
	query := "DELETE FROM Product WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err!= nil {
        fmt.Println(err)
        return false, err
    }
	if result!= nil {
		rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 1 {
            fmt.Printf("Producto eliminado con éxito.\n")
            return true, nil
        } else {
            fmt.Printf("Ningún producto fue eliminado.\n")
            return false, nil
        }
    } else {
        fmt.Printf("Ha habido un error en la consulta (ningún resultado).\n")
        return false, nil
	}
	
}