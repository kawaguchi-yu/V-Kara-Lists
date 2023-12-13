package database

import (
	"fmt"

	"github.com/sharin-sushi/0016go_next_relation/domain"
)

type FavoriteRepository struct {
	SqlHandler
}

func (db *FavoriteRepository) CountMovieFavorites() ([]domain.TransmitMovie, error) {
	fmt.Print("interfaces/database/favorite.go\n")
	// 期待するクエリ
	// SELECT movie_url COUNT(*) AS count FROM favorite_posts WHERE karaoke_id != 0 GROUP BY movie_url;
	var fav domain.Favorite
	var favCnt []domain.TransmitMovie
	err := db.Model(&fav).Select("movie_url").Where("karaoke_id = 0").Group("movie_url").
		Find(&favCnt).Error

	// 動作確認できたら削除するメモ
	// var favs []domain.Favorite
	// var moviesFavCount []domain.MovieFavoriteCount
	// var cnt int64
	// err := db.Model(&favs).Select().Group("movie_url").Count(&cnt).Error

	if err != nil {
		return favCnt, err
	}
	return favCnt, err
}

func (db *FavoriteRepository) CountKaraokeFavorites() ([]domain.TransmitKaraoke, error) {
	fmt.Print("interfaces/database/favorite.go \n")
	// 期待するクエリ
	// SELECT karaoke_id, COUNT(*) AS count FROM favorite_posts WHERE karaoke_id != 0 GROUP BY karaoke_id;
	var fav domain.Favorite
	var favCnt []domain.TransmitKaraoke
	err := db.Model(&fav).Select("karaoke_id").Where("where karaoke_id != 0").Group("karoke_list_id").
		Find(&favCnt).Error
	if err != nil {
		return nil, err
	}
	return favCnt, err
}

func (db *FavoriteRepository) DeleteMovieFavorite(fav domain.Favorite) error {
	fmt.Print("interfaces/database/favorite.go\n")
	whereQu := fmt.Sprintf("listener_id = %v AND movie_url = '%v' AND karaoke_id = 0", fav.ListenerId, fav.MovieUrl)
	err := db.Where(whereQu).Delete(&fav).Error
	fmt.Print("check\n")
	if err != nil {
		return err
	}
	fmt.Print("check2 \n")

	return err
}

func (db *FavoriteRepository) DeleteKaraokeFavorite(fav domain.Favorite) error {
	fmt.Print("interfaces/database/favorite.go\n")
	whereQu := fmt.Sprintf("listener_id = %v AND movie_url = '%v' AND karaoke_id = %v", fav.ListenerId, fav.MovieUrl, fav.KaraokeId)
	err := db.Where(whereQu).Delete(&fav).Error
	if err != nil {
		return err
	}
	return err
}

// 使ってないと思う
func (db *FavoriteRepository) GetAllFavContensByListenerId(favs []domain.Favorite) ([]domain.VtuberMovie, []domain.VtuberMovieKaraoke, error) {
	fmt.Print("interfaces/database/favorite.go \n")
	var Mos []domain.Movie
	var VtsMos []domain.VtuberMovie
	var Kas []domain.Karaoke
	var errs []error
	var err error
	joinsQOfVtsMos := "LEFT JOIN vtubers USING(vtuber_id)"
	whereOfVtsMos := fmt.Sprintf("where movies.inputter_listener_id = %v", favs[0].ListenerId)
	err = db.Model(Mos).Where(whereOfVtsMos).Joins(joinsQOfVtsMos).Scan(&VtsMos).Error
	if err != nil {
		errs = append(errs, err)
	}

	var VtsMosKas []domain.VtuberMovieKaraoke
	joinsQOfVtsMosKas := "LEFT JOIN movies USING(movie_url) LEFT JOIN vtubers USING(vtuber_id)"
	whereOfVtsMosKas := fmt.Sprintf("where karaoke_lists.inputter_listener_id = %v", favs[0].ListenerId)
	err = db.Model(Kas).Where(whereOfVtsMosKas).Joins(joinsQOfVtsMosKas).Scan(&VtsMosKas).Error
	if err != nil {
		errs = append(errs, err)
	}
	return VtsMos, VtsMosKas, err
}

func (db *FavoriteRepository) GetVtubersMoviesWithFavCnts() ([]domain.TransmitMovie, error) {
	fmt.Print("interfaces/database/favorite.go \n")
	var TmMos []domain.TransmitMovie
	var err error

	var vt domain.Vtuber
	selectQu1 := "vtubers.vtuber_id, vtubers.vtuber_name, vtubers.vtuber_kana, vtubers.intro_movie_url, vtubers.vtuber_inputter_id"
	selectQu2 := "m.movie_url, m.movie_title, m.movie_inputter_id"
	selectQu3 := "COUNT(favorites.movie_url) AS count "
	joinQu1 := "LEFT JOIN movies as m USING(vtuber_id)"
	joinQu2 := "LEFT JOIN favorites ON m.movie_url = favorites.movie_url AND favorites.karaoke_id = 0 "
	joinQu := fmt.Sprint(joinQu1, joinQu2)
	whereQu := "m.movie_url IS NOT NULL "
	groupQu := "m.movie_url"
	err = db.Model(vt).Select(selectQu1, selectQu2, selectQu3).
		Joins(joinQu).Where(whereQu).Group(groupQu).
		Scan(&TmMos).Error
	if err != nil {
		return nil, err
	}
	return TmMos, nil
}

func (db *FavoriteRepository) GetVtubersMoviesKaraokesWithFavCnts() ([]domain.TransmitKaraoke, error) {
	fmt.Print("interfaces/database/favorite.go \n")
	var TmKas []domain.TransmitKaraoke
	var err error

	var vt domain.Vtuber
	selectQu1 := "vtubers.vtuber_id, vtubers.vtuber_name, vtubers.vtuber_kana, vtubers.intro_movie_url, vtubers.vtuber_inputter_id "
	selectQu2 := "m.movie_url, m.movie_title, m.movie_inputter_id "
	selectQu3 := "k.karaoke_id, k.sing_start, k.song_name, k.karaoke_inputter_id "
	selectQu4 := "COUNT(f.karaoke_id) AS count"
	joinQu1 := "LEFT JOIN movies as m USING(vtuber_id) "
	joinQu2 := "LEFT JOIN karaokes as k ON m.movie_url = k.movie_url "
	joinQu3 := "LEFT JOIN favorites as f ON k.karaoke_id = f.karaoke_id AND f.karaoke_id != 0 "
	joinQu := fmt.Sprint(joinQu1, joinQu2, joinQu3)
	whereQu := "m.movie_url IS NOT NULL AND k.karaoke_id != 0 "
	groupQu := "k.karaoke_id"
	err = db.Model(vt).Select(selectQu1, selectQu2, selectQu3, selectQu4).
		Joins(joinQu).Where(whereQu).Group(groupQu).
		Scan(&TmKas).Error
	if err != nil {
		return nil, err
	}
	return TmKas, nil
}
func (db *FavoriteRepository) FindFavMoviesByListenerId(favs []domain.Favorite) ([]domain.VtuberMovie, error) {
	fmt.Print("interfaces/database/favorite.go \n")
	var Mos []domain.Movie
	var VtsMos []domain.VtuberMovie
	var errs []error
	var err error
	joinsQOfVtsMos := "LEFT JOIN vtubers USING(vtuber_id)"
	whereOfVtsMos := fmt.Sprintf("where movies.inputter_listener_id = %v", favs[0].ListenerId)
	err = db.Model(Mos).Where(whereOfVtsMos).Joins(joinsQOfVtsMos).Scan(&VtsMos).Error
	if err != nil {
		errs = append(errs, err)
	}
	return VtsMos, err
}

func (db *FavoriteRepository) FindFavKaraokesByListenerId(favs []domain.Favorite) ([]domain.VtuberMovieKaraoke, error) {
	fmt.Print("interfaces/database/favorite.go \n")
	var Kas []domain.Karaoke
	var errs []error
	var err error
	var VtsMosKas []domain.VtuberMovieKaraoke
	joinsQOfVtsMosKas := "LEFT JOIN movies USING(movie_url) LEFT JOIN vtubers USING(vtuber_id)"
	whereOfVtsMosKas := fmt.Sprintf("where karaoke_lists.inputter_listener_id = %v", favs[0].ListenerId)
	err = db.Model(Kas).Where(whereOfVtsMosKas).Joins(joinsQOfVtsMosKas).Scan(&VtsMosKas).Error
	if err != nil {
		errs = append(errs, err)
	}
	return VtsMosKas, err
}

func (db *FavoriteRepository) FindFavsOfUser(Lid domain.ListenerId) ([]domain.Favorite, error) {
	fmt.Print("interfaces/database/favorite.go \n")
	// 期待するクエリ
	// Select * FROM favorites Where listener_id = ?
	var favsOfUser []domain.Favorite
	whereQuery := fmt.Sprintf("listener_id = %v", Lid)
	err := db.Where(whereQuery).Find(&favsOfUser).Error
	if err != nil {
		return nil, err
	}
	return favsOfUser, err
}

func (db *FavoriteRepository) FindFavsByListenerId(lid domain.ListenerId, fav domain.Favorite) (domain.Favorite, error) {
	fmt.Print("interfaces/database/favorite_db.go \n")
	result := db.Where("listener_id=?", lid).Find(fav)
	return fav, result.Error
}

func (db *FavoriteRepository) FindFavoriteUnscopedByFavOrUnfavRegistry(fav domain.Favorite) domain.Favorite {
	fmt.Print("interfaces/database/favorite.go\n")
	whereQu := fmt.Sprintf("listener_id = '%v' AND movie_url = '%v' AND karaoke_id = '%v'", fav.ListenerId, fav.MovieUrl, fav.KaraokeId)
	err := db.Unscoped().Where(whereQu).First(&fav).Error
	fmt.Printf("FindFavoriteUnscopedByFavOrUnfavRegistry got err=%v\n", err)
	return fav
}

func (db *FavoriteRepository) CreateMovieFavorite(fav domain.Favorite) error {
	fmt.Print("interfaces/database/favorite.go\n")
	err := db.Create(&fav).Error
	if err != nil {
		return err
	}
	return err
}

func (db *FavoriteRepository) CreateKaraokeFavorite(fav domain.Favorite) error {
	fmt.Print("interfaces/database/favorite.go\n")
	err := db.Create(&fav).Error
	if err != nil {
		return err
	}
	return err
}

func (db *FavoriteRepository) UpdateMovieFavorite(fav domain.Favorite) error {
	fmt.Print("interfaces/database/favorite.go\n")
	err := db.Unscoped().Model(fav).Update("deleted_at", nil).Error
	if err != nil {
		return err
	}
	return err
}

func (db *FavoriteRepository) UpdateKaraokeFavorite(fav domain.Favorite) error {
	fmt.Print("interfaces/database/favorite.go\n")
	whereQu := fmt.Sprintf("listener_id = '%v' AND movie_url = '%v' AND karaoke_id = %v", fav.ListenerId, fav.MovieUrl, fav.KaraokeId)
	err := db.Unscoped().Model(fav).Where(whereQu).Update("deleted_at", nil).Error
	if err != nil {
		return err
	}
	return err
}

// /////////////////////////////////////////////
// /////////////////////////////////////////////
// /////////////////////////////////////////////
// /////////////////////////////////////////////
// /////////////////////////////////////////////

// 以下、開発中

func (db *FavoriteRepository) FindVtubersThisListenerIdCreated(lid domain.ListenerId) ([]domain.Vtuber, error) {
	fmt.Print("interfaces/database/favorite.go \n")
	var vts []domain.Vtuber
	err := db.Where("vtuber_inputter_id = ?", lid).Find(&vts).Error
	if err != nil {
		return vts, err
	}
	return vts, nil
}

func (db *FavoriteRepository) FindMoviesThisListenerIdCreated(lid domain.ListenerId) ([]domain.TransmitMovie, error) {
	fmt.Print("interfaces/database/favorite.go\n")
	var vt domain.Vtuber
	var tmMos []domain.TransmitMovie
	err := db.Model(vt).Where("movie_inputter_id = ?", lid).Joins("LEFT JOIN movies as m USING(vtuber_id)").Find(&tmMos).Error

	// selectQu1 := "vtubers.vtuber_id, vtubers.vtuber_name, vtubers.vtuber_kana, vtubers.intro_movie_url, vtubers.vtuber_inputter_id"
	// selectQu2 := "m.movie_url, m.movie_title, m.movie_inputter_id"
	// selectQu3 := "COUNT(favorites.movie_url) AS count "
	// joinQu1 := "LEFT JOIN movies as m USING(vtuber_id)"
	// joinQu2 := "LEFT JOIN favorites ON m.movie_url = favorites.movie_url AND favorites.karaoke_id = 0 "
	// joinQu := fmt.Sprint(joinQu1, joinQu2)
	// whereQu := "m.movie_url IS NOT NULL "
	// groupQu := "m.movie_url"
	// err = db.Model(vt).Select(selectQu1, selectQu2, selectQu3).
	// 	Joins(joinQu).Where(whereQu).Group(groupQu).
	// 	Scan(&tmMos).Error
	if err != nil {
		return tmMos, err
	}
	return tmMos, err
}

// まだーーーー
func (db *FavoriteRepository) FindKaraokesThisListenerIdCreated(lid domain.ListenerId) ([]domain.TransmitKaraoke, error) {
	fmt.Print("interfaces/database/favorite.go\n")
	var err error
	var tmKas []domain.TransmitKaraoke
	return tmKas, err
}
func (db *FavoriteRepository) FindMoviesThisListenerIdFavorited(lid domain.ListenerId) ([]domain.TransmitMovie, error) {
	fmt.Print("interfaces/database/favorite.go\n")
	var err error
	var tmMos []domain.TransmitMovie
	return tmMos, err
}
func (db *FavoriteRepository) FindKaraokesThisListenerIdFavorited(lid domain.ListenerId) ([]domain.TransmitKaraoke, error) {
	fmt.Print("interfaces/database/favorite.go\n")
	var err error
	var tmKas []domain.TransmitKaraoke
	return tmKas, err
}
