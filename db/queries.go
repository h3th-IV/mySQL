package db

import "log"

func Insert() {
	query := `insert into user(name, password) values('mAcBobby', 'password12');`
	_, err := dB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}
