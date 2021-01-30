package database

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"reflect"
	"strings"
	"text/template"
)

type Database struct {
	GormDB      *gorm.DB
	neo4j       neo4j.Driver
	neo4jKeys   []string
	neo4jValues []string
	Driver      string
}

func Open(driver string, dsn string) (*Database, error) {
	var err error
	gormLogger := logger.Default.LogMode(logger.Error)
	r := Database{
		Driver: driver,
	}
	switch driver {
	case "sqlite":
		r.GormDB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
		if err != nil {
			return nil, err
		}
		r.GormDB.Exec("PRAGMA foreign_keys = ON")
		sqlDB, err := r.GormDB.DB()
		if err != nil {
			return nil, err
		}
		sqlDB.SetMaxOpenConns(1)
	case "postgresql":
		r.GormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
	case "mysql":
		r.GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
	case "sqlserver":
		r.GormDB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
	case "neo4j":
		u, err := url.Parse(dsn)
		if err != nil {
			return nil, err
		}
		password, _ := u.User.Password()
		r.neo4j, err = neo4j.NewDriver(dsn, neo4j.BasicAuth(u.User.Username(), password, ""))
	default:
		return nil, fmt.Errorf("database driver only supports one of sqlite,postgresql,mysql,sqlserver")
	}

	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (d *Database) AutoMigrate(dst ...interface{}) error {
	if d.GormDB != nil {
		return d.GormDB.AutoMigrate(dst...)
	}
	return nil
}

func getNodeName(v reflect.Value) string {
	NodeName := v.Type().Name()
	m := v.MethodByName("TableName")
	if m.IsValid() {
		r := m.Call([]reflect.Value{})
		NodeName = r[0].String()
	}
	return strcase.ToCamel(NodeName)
}

func (d *Database) Delete(dst ...interface{}) {
	if d.GormDB != nil {
		d.GormDB.Delete(dst[0])
		return
	}

	cypherTpl := `
MATCH (n: {{ .Node }} {
	{{- $FieldsLen := len .Fields }}
	{{- $FieldsLen = minus $FieldsLen 1 }}
	{{- range $index, $element := .Fields }}
		{{- if eq $index $FieldsLen }}
			{{ $element }}: ${{ $element }}
		{{- else }}
			{{ $element }}: ${{ $element }},
		{{- end }}
	{{- end }}	
})
DETACH DELETE (n)
`

	session := d.neo4j.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	params := map[string]interface{}{}
	for i, key := range d.neo4jKeys {
		params[key] = d.neo4jValues[i]
	}

	for _, nodeStruct := range dst {
		var s strings.Builder
		nodeName := getNodeName(reflect.ValueOf(nodeStruct).Elem())
		t := template.Must(template.New("").Funcs(funcMap).Parse(cypherTpl))
		err := t.Execute(&s, map[string]interface{}{
			"Node":   nodeName,
			"Fields": d.neo4jKeys,
		})
		if err != nil {
			log.Fatal(err)
		}

		_, err = session.Run(s.String(), params)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func (d *Database) Where(key string, value string) *Database {
	if d.GormDB != nil {
		return &Database{
			GormDB:      d.GormDB.Where(key+" = ?", value),
			neo4j:       d.neo4j,
			neo4jKeys:   d.neo4jKeys,
			neo4jValues: d.neo4jValues,
		}
	}

	return &Database{
		GormDB:      d.GormDB,
		neo4j:       d.neo4j,
		neo4jKeys:   append(d.neo4jKeys, key),
		neo4jValues: append(d.neo4jValues, value),
	}
}

const chunkSize = 100

func (d *Database) InsertOne(value interface{}) {
	if d.GormDB != nil {
		d.GormDB.Create(reflect.ValueOf(value).Interface())
		return
	}

	session := d.neo4j.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	d.neo4jInsertOne(reflect.ValueOf(value).Elem(), session, nil, "")
}

func (d *Database) ChunkedCreate(value interface{}) {
	arr := reflect.ValueOf(value)
	if d.GormDB != nil {
		for i := 0; i < arr.Len(); i += chunkSize {
			end := i + chunkSize
			if i+chunkSize > arr.Len() {
				end = arr.Len()
			}
			d.GormDB.Create(arr.Slice(i, end).Interface())
		}
		return
	}

	session := d.neo4j.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	d.neo4jInsertMany(arr, session, nil, "")
}

func (d *Database) ChunkedUpsert(value interface{}) {
	arr := reflect.ValueOf(value)
	if d.GormDB != nil {
		for i := 0; i < arr.Len(); i += chunkSize {
			end := i + chunkSize
			if i+chunkSize > arr.Len() {
				end = arr.Len()
			}
			d.GormDB.Clauses(clause.OnConflict{DoNothing: true}).Create(arr.Slice(i, end).Interface())
		}
		return
	}

	session := d.neo4j.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	d.neo4jInsertMany(arr, session, nil, "")
}