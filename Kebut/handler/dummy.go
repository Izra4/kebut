package handler

import (
	"gorm.io/gorm"
	"kebut/database"
	"kebut/entity"
	"log"
)

func PDummy() {
	p1 := entity.Penyakit{
		Model:  gorm.Model{},
		Nama:   "Rebah Kecambah",
		Gejala: "1. Daun pucat.;2. Daun layu.;3. Daun menguning.;4. Batang berlekuk.;5. Batang membusuk.;6. Batang menggelap.",
		Penanganan: "1. Sanitasi.;2. Pengolahan tanah untuk pembibitan sebanyak 3-4x dengan selang waktu 5-7 hari.;3. Penjarahan bibit dan pengaturan atap pembibitan untuk mengurangi kelembaban.;" +
			"4. Mendisinfeksi tanah sebelum penaburan benih dengan larutan cuSO4, kapur tohor, fungisida metalaksil.;5. Penyemprotan pembibitan dengan fungisida antara lain metalaksil, mankozeb, benomil, propaniokarb hidroklorida.",
		Links: nil,
	}

	p2 := entity.Penyakit{
		Model:  gorm.Model{},
		Nama:   "Busuk Hitam / Black Shenk",
		Gejala: "1. Daun layu.;2. Daun menggelap.;3. Lesi tengah daun berwarna hitam.;4. Bercak daun meluas.;5. Batang menggelap.;6. Lesi batang gelap.;7. Batang membengkak.",
		Penanganan: "1. Pengelolaan drainase, bersihkan drainase dari kotoran.;2. Rotasi tanaman yang sakit dengan yang baru yang sehat.;3. Pertahankan kebersihan kebun.;" +
			"4. Pengaturan irigasi yang baik.;5. Aplikasikan Fungisida secara teratur setiap 1 minggu.;6. Lakukan pemantauan rutin.;7. Praktikkan Agronomi yang baik dengan pemupukan dan pemeliharaan tanaman secara rutin.",
		Links: nil,
	}

	p3 := entity.Penyakit{
		Model:  gorm.Model{},
		Nama:   "Lanas",
		Gejala: "1. Daun layu.;2. Daun menguning.;3. Bercak daun meluas.;4. Batang membusuk.;5. Lesi batang gelap.;6. Batang berjamur.",
		Penanganan: "1. Sanitasi, mencabut tanaman yang sakit kemudian dibakar.;2. Pengolahan tanah untuk pembibitan sebanyak 3-4 kali dengan selang waktu 7-15 hari.;" +
			"3. Menggunakan pupuk kandang yang telah masak.;4. Pembuatan guludan yang tinggi sehingga drainase lebih baik;5. Mendisinfeksi tanah sebelum penaburan benih dengan larutan cuSO4, kapur tohor, fungisida metalaksil.",
		Links: nil,
	}

	p5 := entity.Penyakit{
		Model:  gorm.Model{},
		Nama:   "Kerupuk",
		Gejala: "1. Daun melengkung.;2. Tulang daun bengkok.;3. Tulang daun menebal.",
		Penanganan: "1. Sanitasi.;2. Mencabut tanaman sakit maupun sisa-sisa pertanaman dan gulma kemudian dikumpulkan dan dimusnahkan.;" +
			"3. Pengendalian vektor lalat putih B. tabaci dengan insektisida seperti yang digunakan untuk mengendalikan Myzus persicae.",
		Links: nil,
	}

	p4 := entity.Penyakit{
		Model:      gorm.Model{},
		Nama:       "Mosaik / TMV",
		Gejala:     "1. Daun berpola mosaik.;2. Daun menggulung.;3. Daun kering.;4. Batang membusuk.;5. Batang berpola mosaik.;6. Batang mudah rapuh.",
		Penanganan: "1. Sanitasi, mencabut tanaman yang sakit kemudian dibakar.;2. Mendesinfeksi tangan para pekerja dengan sabun trinatrium fosfat.;3. Menggunakan varietas tahan seperti coker 176, nc 567.",
		Links:      nil,
	}

	p6 := entity.Penyakit{
		Model:  gorm.Model{},
		Nama:   "Kerah Busuk / Colar Cot",
		Gejala: "1. Daun layu.;2. Daun menguning.;3. Daun kering.;4. Daun bintik gelap.;5. Batang membusuk.;6. Batang menggelap.;7. Batang mudah rapuh.",
		Penanganan: "1. Sirkulasi udara yang memadai, yang memadai durasi kebasahan daun, akan berkurang infeksi awal dan penyebarannya.;2. Petani harus mengatur suhu sedemikian rupa yang berguna untuk meningkatkan kesehatan tanaman dan meminimalkan cedera.;" +
			"3. Pertahankan zona bebas gulma di sekitar bedengan pelampung.;4. Kesuburan harus dijaga sekitar 100 bagian per juta(ppm) nitrogen",
		Links: nil,
	}

	p7 := entity.Penyakit{
		Model:  gorm.Model{},
		Nama:   "Akar Busuk / Root Rot",
		Gejala: "1. Daun layu.;2. Daun menguning.;3. Daun rontok.;4. Batang membusuk.;5. Batang mudah rapuh.;6. Batang berlendir.",
		Penanganan: "1. Sanitasi.;2. Aplikasikan fungisida secara tepat waktu, hal ini efektif dalam mengendalikan penyebaran penyakit dan mengurangi dampaknya terhadap tanaman tembakau.;" +
			"3. Melakukan metode Biocontrol, metode ini melibatkan penggunaan mikroorganisme atau senyawa alami untuk menekan pertumbuhan patogen penyebab akar busuk.;" +
			"4. Kelola tingkat kelembaban tanah dengan baik, karena dapat membantu menciptakan lingkungan yang kurang kondusif bagi perkembangan dan penyebaran patogen penyebab akar busuk.",
		Links: nil,
	}

	p8 := entity.Penyakit{
		Model:  gorm.Model{},
		Nama:   "Kerdil Tembakau",
		Gejala: "1. Daun pucat.;2. Daun menguning.;3. Bercak daun meluas.;4. Daun menggulung.;5. Daun mengecil.;6. Batang menggelap.;7. Batang pucat.;8. Batang kerdil.",
		Penanganan: "1. Menggunakan benih yang sehat dan bebas dari virus.;2. Menanam tembakau di lahan yang bebas dari kutu daun.;3. Melakukan rotasi tanaman secara berkala.;" +
			"4. Mengontrol populasi kutu daun dengan menggunakan pestisida.;5. Pemilihan varietas yang baik.;6. Cabut dan bakar tanaman yang terserang penyakit segera mungkin.;" +
			"7. Gunakan pestisida secara bijak dan sesuai dengan petunjuk penggunaan.;8. Periksa tanaman secara rutin untuk mendeteksi gejala penyakit.",
		Links: nil,
	}

	p9 := entity.Penyakit{
		Model:      gorm.Model{},
		Nama:       "Antraknosa",
		Gejala:     "1. Lesi tengah daun berwarna hitam.;2. Bercak daun meluas.;3. Daun kering.;4. Daun berjamur.;5. Batang menggelap.;6. Lesi batang gelap.;7. Batang berjamur berwarna gelap.",
		Penanganan: "1. Memotong daun yang sakit.;2. Mengurangi kelembaban di pembibitan.;3. Aplikasikan fungisida secara tepat waktu, hal ini efektif dalam mengendalikan penyebaran penyakit dan mengurangi dampaknya terhadap tanaman tembakau.",
		Links:      nil,
	}

	if err := database.DB.Create(&p1).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&p2).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&p3).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&p4).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&p5).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&p6).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&p7).Error; err != nil {
		log.Fatalln("Failed to create: ")
	}
	if err := database.DB.Create(&p8).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&p9).Error; err != nil {
		log.Fatalln("Failed to create")
	}
}

func LDummy() {
	l1 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://sinta.ditjenbun.pertanian.go.id/rebah-kecambah/",
		PenyakitId: 1,
	}
	l2 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://ephytia.inra.fr/en/C/10840/Tobacco-Damping-off",
		PenyakitId: 1,
	}
	l3 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://www.apsnet.org/edcenter/disandpath/oomycete/pdlessons/Pages/BlackShank.aspx",
		PenyakitId: 2,
	}
	l4 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://content.ces.ncsu.edu/black-shank",
		PenyakitId: 2,
	}
	l5 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://sinta.ditjenbun.pertanian.go.id/lanas/",
		PenyakitId: 3,
	}
	l6 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://ephytia.inra.fr/en/C/10981/Tobacco-Phytophthora-nicotianae",
		PenyakitId: 3,
	}
	l7 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://distan.bulelengkab.go.id/informasi/detail/artikel/upaya-penanggulangan-virus-mosaik-tembkau-57#:~:text=Virus%20mosaik%20tembakau%20",
		PenyakitId: 4,
	}
	l8 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://www.apsnet.org/edcenter/disandpath/viral/pdlessons/Pages/TobaccoMosaic.aspx",
		PenyakitId: 4,
	}
	l9 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://sinta.ditjenbun.pertanian.go.id/kerupuk/",
		PenyakitId: 5,
	}
	l10 := entity.Link{
		Model:      gorm.Model{},
		Link:       "http://www.eagri.org/eagri50/PATH272/lecture08/009.html",
		PenyakitId: 5,
	}
	l11 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://ephytia.inra.fr/en/C/10797/Tobacco-Collar-rot-Sclerotinia-sclerotiorum",
		PenyakitId: 6,
	}
	l12 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://ctri.icar.gov.in/for_controlDisease.php",
		PenyakitId: 6,
	}

	l13 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://jateng.antaranews.com/berita/80799/tanaman-tembakau-busuk-akar",
		PenyakitId: 7,
	}

	l14 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://plantix.net/id/library/plant-diseases/200074/tobacco-leaf-curl-disease/",
		PenyakitId: 8,
	}

	l15 := entity.Link{
		Model:      gorm.Model{},
		Link:       "https://sinta.ditjenbun.pertanian.go.id/antraknosa/",
		PenyakitId: 9,
	}

	l16 := entity.Link{
		Model:      gorm.Model{},
		Link:       "http://eagri.org/eagri50/PATH272/lecture08/006.html#:~:text=Initially%2C%20infection%20starts%20on%20lower,to%20form%20large%20necrotic%20lesions",
		PenyakitId: 9,
	}
	if err := database.DB.Create(&l1).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l2).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l3).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l4).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l5).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l6).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l7).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l8).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l9).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l10).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l11).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l12).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l13).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l14).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l15).Error; err != nil {
		log.Fatalln("Failed to create")
	}
	if err := database.DB.Create(&l16).Error; err != nil {
		log.Fatalln("Failed to create")
	}

}
