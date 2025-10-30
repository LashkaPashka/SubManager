package workoption

import "github.com/jackc/pgx/v5"

func WorkDate(date string) (query string, args pgx.NamedArgs) {
	query = `SELECT COALESCE(SUM(price), 0) AS total
			 FROM subscriptions
			 WHERE TO_CHAR(start_date, 'MM-YYYY') = @date
			`
	
	args = pgx.NamedArgs{
		"date": date,
	}

	return query, args
}

func WorkUserID(date string, userID string) (query string, args pgx.NamedArgs) {
	query = `SELECT COALESCE(SUM(price), 0) AS total
			 FROM subscriptions
			 WHERE user_id = @user_id
				AND TO_CHAR(start_date, 'MM-YYYY') = @date
			`

	args = pgx.NamedArgs{
		"user_id": userID,
		"date": date,
	}

	return query, args
}

func WorkServiceName(date, serviceName string) (string, pgx.NamedArgs) {
	query := `SELECT COALESCE(SUM(price), 0) AS total
			  FROM subscriptions
			  WHERE service_name = @service_name
		  		 AND TO_CHAR(start_date, 'MM-YYYY') = @date
			 `

	args := pgx.NamedArgs{
		"service_name": serviceName,
		"date":         date,
	}

	return query, args
}
