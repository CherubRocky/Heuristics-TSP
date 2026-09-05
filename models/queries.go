package models

import(
	"strings"
	"github.com/CherubRocky/Heuristics-TSP/tsp"
)

func (db *DB) QueryConnections(cit []int) ([]tsp.Connection, error) {
	if len(cit) == 0 {
		return []tsp.Connection{}, nil
	}
	
	query := getPlaceHolders(cit)
	args := make([]any, 0, 2 * len(cit))
	for _, val := range cit {
		args = append(args, val)
	}
	for _, val := range cit {
		args = append(args, val)
	}
	
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	connections := make([]tsp.Connection, 0, len(cit) * (len(cit) - 1) / 2)
	for rows.Next() {
		var conn tsp.Connection
		if err := rows.Scan(&conn.Cit1.Id, &conn.Cit2.Id, &conn.Distance); err != nil {
			return nil, err
		}
		connections = append(connections, conn)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return connections, nil
}

func getPlaceHolders(arr []int) string {
	placeholders := strings.Repeat("?,", len(arr))
	placeholders = strings.TrimRight(placeholders, ",")

	var b strings.Builder
	b.WriteString(`SELECT
id_city_1,
id_city_2,
distance
FROM connections
WHERE id_city_1 IN (`)
	b.WriteString(placeholders)
	b.WriteString(") AND id_city_2 IN (")
	b.WriteString(placeholders)
	b.WriteRune(')')
	
	return b.String()
}

func (db *DB) QueryCities(cits []int) ([]tsp.City, error) {
	if len(cits) == 0 {
		return []tsp.City{}, nil
	}
	
	placeholders := strings.Repeat("?,", len(cits))
	placeholders = strings.TrimRight(placeholders, ",")

	var b strings.Builder
	b.WriteString(`SELECT id, latitude, longitude FROM cities WHERE id IN (`)
	b.WriteString(placeholders)
	b.WriteString(")")
	
	query := b.String()
	args := make([]any, 0, len(cits))
	for _, val := range cits {
		args = append(args, val)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cities := make([]tsp.City, 0, len(cits))
	for rows.Next() {
		var city tsp.City
		var lat, lon float64
		if err := rows.Scan(&city.Id, &lat, &lon); err != nil {
			return nil, err
		}
		city.Coords = tsp.NewCoordinates(lat, lon)
		cities = append(cities, city)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cities, nil
}
