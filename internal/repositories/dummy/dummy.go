package dummyRepository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/repositories/base_repository"
	"go-skeleton/pkg/database"
	"strings"
)

type DummyRepository struct {
	*base_repository.BaseRepo[dummy.Dummy]
}

func NewDummyRepo(mysql *database.MySql) *DummyRepository {
	return &DummyRepository{
		BaseRepo: base_repository.NewBaseRepository[dummy.Dummy](mysql),
	}
}

func (d *DummyRepository) CreateRaw(db *sqlx.DB) {
	args := map[string]interface{}{
		"id":   "sdsdsd",
		"name": "Samuel da Silva",
	}

	st := struct {
		Id   string `db:"id"`
		Name string `db:"name"`
	}{
		Id:   "sdsds",
		Name: "Samuel da Silva",
	}

	values := []string{}
	fields := []string{}
	for f, _ := range args {
		values = append(values, ":"+f)
		fields = append(fields, f)
	}

	f := strings.Join(fields, ", ")
	v := strings.Join(values, ", ")

	exec, err := db.NamedExec("insert into dummies ("+f+") values ("+v+")", st)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(exec)
	fmt.Println("funfou")
}
