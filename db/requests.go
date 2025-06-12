package db

import (
	"TransportManagementService/proto"
	"fmt"
	"strings"
	"time"
)

func CreateTransportRequest(input *proto.CreateTransportRequest) (int64, error) {
	stmt, err := DB.Prepare(`
        INSERT INTO transports (
            number, type_id, is_active, current_driver_id
        ) VALUES (?, ?, ?, ?)
    `)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		input.Number,
		input.TypeId,
		input.IsActive,
		input.CurrentDriverId,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func UpdateTransportRequest(input *proto.UpdateTransportRequest) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var fields []string
	var args []interface{}

	if input.CurrentDriverId != nil {
		fields = append(fields, "current_driver_id = ?")
		args = append(args, input.CurrentDriverId.Value)
	}
	if input.IsActive != nil {
		fields = append(fields, "is_active = ?")
		args = append(args, input.IsActive.Value)
	}

	if len(fields) == 0 {
		return false, fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf("UPDATE transports SET %s WHERE id = ?", strings.Join(fields, ", "))
	args = append(args, input.TransportId.Value)

	stmt, err := tx.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return false, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if err := tx.Commit(); err != nil {
		return false, fmt.Errorf("ошибка при коммите транзакции: %w", err)
	}

	return affected > 0, nil
}

func GetTransportRequest(input *proto.GetTransportInfoRequest) ([]*proto.TransportInfo, error) {
	baseQuery := `
		SELECT
			t.id,
			t.number,
			tt.name AS type_name,
			t.is_active,
			t.current_driver_id
		FROM transports t
		JOIN transport_type tt ON t.type_id = tt.id

    `

	var args []interface{}
	var conditions []string

	if input.TransportId != nil {
		conditions = append(conditions, "t.id = ?")
		args = append(args, input.TransportId.Value)
	}

	if input.IsActive != nil {
		conditions = append(conditions, "t.is_active = ?")
		args = append(args, input.IsActive.Value)
	}
	if input.CurrentDriverId != nil {
		conditions = append(conditions, "t.current_driver_id = ?")
		args = append(args, input.CurrentDriverId.Value)
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + conditions[0]
		for i := 1; i < len(conditions); i++ {
			baseQuery += " AND " + conditions[i]
		}
	}

	rows, err := DB.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transports []*proto.TransportInfo

	for rows.Next() {
		var r proto.TransportInfo
		// var createdAt time.Time

		err := rows.Scan(
			&r.TransportId,
			&r.Number,
			&r.TransportType,
			&r.IsActive,
			&r.CurrentDriverId,
			//&createdAt,
		)
		if err != nil {
			return nil, err
		}

		// r.PreferredDate = preferredDate.Format(time.RFC3339)
		// r.CreatedAt = createdAt.Format(time.RFC3339)

		transports = append(transports, &r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transports, nil
}

func CreateTransportLogRequest(input *proto.CreateTransportLogRequest) (int64, error) {
	stmt, err := DB.Prepare(`
        INSERT INTO transport_logs (
            transport_id, service_type_id, description, service_date, mileage
        ) VALUES (?, ?, ?, ?, ?)
    `)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		input.TransportId,
		input.ServiceTypeId,
		input.Description,
		input.ServiceDate,
		input.Mileage,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func GetTransportLogsRequest(input *proto.GetTransportLogsInfoRequest) ([]*proto.TransportLogInfo, error) {
	baseQuery := `
		SELECT
			tl.id,
			tl.transport_id,
			st.name AS service_type,
			tl.description,
			tl.mileage,
			tl.service_date
		FROM transport_logs tl
		JOIN service_type st ON tl.service_type_id = st.id
    `

	var args []interface{}
	var conditions []string

	if input.TransportId != nil {
		conditions = append(conditions, "tl.transport_id = ?")
		args = append(args, input.TransportId.Value)
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + conditions[0]
		for i := 1; i < len(conditions); i++ {
			baseQuery += " AND " + conditions[i]
		}
	}

	rows, err := DB.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transportLogs []*proto.TransportLogInfo

	for rows.Next() {
		var r proto.TransportLogInfo
		var serviceDate time.Time

		err := rows.Scan(
			&r.Id,
			&r.TransportId,
			&r.ServiceType,
			&r.Description,
			&r.Mileage,
			&serviceDate,
			//&createdAt,
		)
		if err != nil {
			return nil, err
		}

		r.ServiceDate = serviceDate.Format(time.RFC3339)
		// r.CreatedAt = createdAt.Format(time.RFC3339)

		transportLogs = append(transportLogs, &r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transportLogs, nil
}
