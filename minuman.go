package model


type Minuman struct {
	ID_Makanan string `json:"id_minuman"`
	Id_toko string `json:"id_toko"`
	Nama_Makanan string `json:"nama_minuman"`
	Description_Makanan string `json:"description_minuman"`
	Harga_Makanan string `json:"harga_minuman"`
}