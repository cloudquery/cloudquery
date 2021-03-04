package resources

import "gorm.io/gorm/schema"

// List of tables created by Okta Provider
var ResourceTables = []schema.Tabler{
	&User{},
	&UserType{},
	&Group{},
	&Application{},
	&ApplicationUser{},
	&ApplicationGroup{},
}

