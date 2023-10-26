package db

import (
	"database/sql"
	"reflect"
)

func MapQueryResultToStruct(rows *sql.Rows, resultStruct interface{}) ([]interface{}, error) {
	defer rows.Close()

	var resultSlice []interface{}

	// Get the type of the resultStruct
	resultType := reflect.TypeOf(resultStruct)

	for rows.Next() {
		// Create a new instance of the result struct
		result := reflect.New(resultType).Interface()

		// Get the value of the result
		resultValue := reflect.Indirect(reflect.ValueOf(result))

		// Iterate through the fields of the result struct
		for i := 0; i < resultType.Elem().NumField(); i++ {
			field := resultType.Elem().Field(i)
			tag := field.Tag.Get("db") // Get the "db" tag value

			// If the tag is not empty, set the value from the query result
			if tag != "" {
				resultValue.Field(i).Set(reflect.ValueOf(tag))
			}
		}

		// Append the result to the resultSlice
		resultSlice = append(resultSlice, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resultSlice, nil
}
