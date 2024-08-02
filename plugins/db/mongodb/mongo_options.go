/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package mongodb

import (
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/framework/plugins/db"
	"github.com/elliotchance/pie/v2"
)

type MongoDBOptions struct {
	DBInfos []db.DBInfo
}

var With _Option

type _Option struct{}

func (_Option) Default() option.Setting[MongoDBOptions] {
	return func(options *MongoDBOptions) {
		With.DBInfos().Apply(options)
	}
}

func (_Option) DBInfos(infos ...db.DBInfo) option.Setting[MongoDBOptions] {
	return func(options *MongoDBOptions) {
		options.DBInfos = pie.Filter(infos, func(info db.DBInfo) bool {
			switch info.Type {
			case db.MongoDB:
				return true
			}
			return false
		})
	}
}
