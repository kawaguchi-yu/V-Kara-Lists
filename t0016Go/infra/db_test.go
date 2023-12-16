package infra

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sharin-sushi/0016go_next_relation/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	// skip transaction for sqlmock testing
	gormdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return nil, nil, err
	}
	return gormdb, mock, nil
}

func testMigrate(db *gorm.DB, mock sqlmock.Sqlmock) error {
	db.AutoMigrate(
		// User
		domain.Listener{},
		// Like Relatoin
		domain.Favorite{}, domain.Follow{},
		// Vtuber Contents
		domain.Karaoke{}, domain.Movie{}, domain.Vtuber{}, domain.OriginalSong{},
	)
	return nil
}

///////////以下、クエリ生成のテストコード///////////

func TestSqlHandler_Create(t *testing.T) {
	userBody := getCreateUser()
	now := time.Now()

	userBody.UpdatedAt = now
	userBody.DeletedAt.Valid = false

	db, mock, err := getMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	r := SqlHandler{Conn: db}
	query := "INSERT"
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(userBody.ListenerName, userBody.Email, userBody.Password, userBody.UpdatedAt, userBody.DeletedAt).
		// WithArgs("test", userBody.Email, userBody.Password, userBody.UpdatedAt, userBody.DeletedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err != nil {
		t.Log(err)
		t.Log("mockdb is down")
		t.Fail()
	}
	if err := r.Create(userBody).Error; err != nil {
		t.Log(err)
		t.Log("create is down")
		t.Fail()
	}
}

func getCreateUser() *domain.Listener {
	var timeTime time.Time
	var gormDeletedAt gorm.DeletedAt
	body := &domain.Listener{
		ListenerName: "shari",
		Email:        "shari@gmail.com",
		Password:     "$5G$aaabbbccc",
		CreatedAt:    timeTime,
		UpdatedAt:    timeTime,
		DeletedAt:    gormDeletedAt,
	}
	return body
}

// gorm.DeletedAtが有れば論理削除、無ければ物理削除
// 便宜的に論理削除を1、物理削除を2としてテスト関数を名付けた。
// func TestSqlHandler_Delete1(t *testing.T) {
// 	db, mock, err := getMockDB()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	r := SqlHandler{Conn: db}

// 	logicallBody := getLogicalDeleteUser()
// 	now := time.Now()
// 	logicallBody.DeletedAt = gorm.DeletedAt{Time: now, Valid: true}

// 	logilalyQuery := "UPDATE `listeners` SET `deleted_at`=? WHERE listener_id = ? AND `listeners`.`deleted_at` IS NULL"
// 	mock.ExpectExec(regexp.QuoteMeta(logilalyQuery)).
// 		WithArgs(logicallBody.DeletedAt.Time, logicallBody.ListenerId).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	if err != nil {
// 		t.Log(err)
// 		t.Log("mockdb is down")
// 		t.Fail()
// 	}
// 	if err := r.Delete(&domain.Listener{}, "listener_id = ?", logicallBody.ListenerId).Error; err != nil {
// 		t.Log(err)
// 		t.Log("logicaly delete is down")
// 		t.Fail()
// 	}
// }

func TestSqlHandler_Delete1(t *testing.T) {
	db, mock, err := getMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	r := SqlHandler{Conn: db}

	logicallBody := getLogicalDeleteUser()
	now := time.Now()
	logicallBody.DeletedAt = gorm.DeletedAt{Time: now, Valid: true}

	logilalyQuery := "UPDATE `listeners` SET `deleted_at`=? WHERE listener_id = ? AND `listeners`.`deleted_at` IS NULL"
	mock.ExpectExec(regexp.QuoteMeta(logilalyQuery)).
		WithArgs(sqlmock.AnyArg(), logicallBody.ListenerId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err != nil {
		t.Log(err)
		t.Log("mockdb is down")
		t.Fail()
	}
	if err := r.Delete(&domain.Listener{}, "listener_id = ?", logicallBody.ListenerId).Error; err != nil {
		t.Log(err)
		t.Log("logicaly delete is down")
		t.Fail()
	}
}

func getLogicalDeleteUser() *domain.Listener {
	// var timeTime time.Time
	var gormDeletedAt gorm.DeletedAt
	body := &domain.Listener{
		ListenerId: 1,
		// ListenerName: "shari",
		// Email:        "shari@gmail.com",
		// Password:     "$5G$aaabbbccc",
		// CreatedAt:    timeTime,
		// UpdatedAt: timeTime,
		DeletedAt: gormDeletedAt,
	}
	return body
}

func TestSqlHandler_Delete2(t *testing.T) {
	db, mock, err := getMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	r := SqlHandler{Conn: db}

	phycialBody := getPhycicalDeleteMovie()

	phycialQuery := "DELETE FROM `movies` WHERE movie_url = ?"
	mock.ExpectExec(regexp.QuoteMeta(phycialQuery)).
		WithArgs(phycialBody.MovieUrl).
		WillReturnResult(sqlmock.NewResult(1, 1))
	if err != nil {
		t.Log(err)
		t.Log("mockdb is down")
		t.Fail()
	}
	if err := r.Delete(&domain.Movie{}, "movie_url = ?", phycialBody.MovieUrl).Error; err != nil {
		t.Log(err)
		t.Log("phygical delete is down")
		t.Fail()
	}
}

func getPhycicalDeleteMovie() *domain.Movie {
	body := &domain.Movie{
		MovieUrl: "url1",
		VtuberId: 1,
	}
	return body
}