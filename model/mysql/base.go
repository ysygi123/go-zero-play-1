package mysql

import "gorm.io/gorm"

func GeneralSql(queryFunc func(tx *gorm.DB) *gorm.DB) (string, []interface{}) {
	tx := queryFunc(&gorm.DB{
		Config: &gorm.Config{
			SkipDefaultTransaction:                   false,
			NamingStrategy:                           nil,
			FullSaveAssociations:                     false,
			Logger:                                   nil,
			NowFunc:                                  nil,
			DryRun:                                   true,
			PrepareStmt:                              false,
			DisableAutomaticPing:                     false,
			DisableForeignKeyConstraintWhenMigrating: false,
			DisableNestedTransaction:                 false,
			AllowGlobalUpdate:                        false,
			QueryFields:                              false,
			CreateBatchSize:                          0,
			ClauseBuilders:                           nil,
			ConnPool:                                 nil,
			Dialector:                                nil,
			Plugins:                                  nil,
		},
		Error:        nil,
		RowsAffected: 0,
		Statement:    nil,
	})

	return tx.Statement.SQL.String(), tx.Statement.Vars
}
