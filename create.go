package Toko

func createToko(w http.ResponseWriter, r *hhtp.Request){	

	db , err := ConncetionToDatabase()

	if err != nil {
		log.Fatal(err)
	}

	queryText := fmt.Sprintf("INSERT INTO toko (id_toko , nama_toko, alamat_toko, picture_toko , rating_toko) values( %v ,'%v','%v','%v','%v','%v','%v')",
		toko.ID,
		toko.Restaurant_id,
		toko.Name,
		toko.Description,
		toko.PictureID,
		toko.City,
		toko.Rating,
	)
}