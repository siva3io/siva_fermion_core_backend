package omnichannel

import (
	model_core "fermion/backend_core/internal/model/core"

	"github.com/lib/pq"
)

/*
Copyright (C) 2022 Eunimart Omnichannel Pvt Ltd. (www.eunimart.com)
All rights reserved.
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License v3.0 as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License v3.0 for more details.
You should have received a copy of the GNU Lesser General Public License v3.0
along with this program.  If not, see <https://www.gnu.org/licenses/lgpl-3.0.html/>.
*/
type OmnichannelFav struct {
	model_core.Model
	UserID              uint                 `json:"user_id"`
	User                model_core.CoreUsers `json:"user" gorm:"foreignKey:UserID;references:ID"`
	MarketplaceIds      pq.Int64Array        `json:"marketplace_ids" gorm:"type:int[]"`
	WebstoreIds         pq.Int64Array        `json:"webstore_ids" gorm:"type:int[]"`
	VirtualWarehouseIds pq.Int64Array        `json:"virtual_warehouse_ids" gorm:"type:int[]"`
}
