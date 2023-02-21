package pos

import (
	"fmt"

	"fermion/backend_core/internal/model/pagination"
	"fermion/backend_core/pkg/util/cache"
	"fermion/backend_core/pkg/util/helpers"
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


func UpdatePosLinkInCache(token_id string, access_template_id string) {
	p := new(pagination.Paginatevalue)
	searchquery := map[string]interface{}{}
	data, _ := NewService().GetPosLinkList(searchquery, p, token_id, access_template_id, "LIST")
	key := fmt.Sprintf("%v:%v", token_id, cache.POS_LINK)
	fmt.Println("Updating Key:", key)
	cache.SetCacheKey(key, data)
	key = fmt.Sprintf("%v:%v:%v", token_id, cache.POS_LINK, cache.PAGINATION)
	fmt.Println("Updating Key:", key)
	cache.SetCacheKey(key, p)
}

func GetPosLinkFromCache(user_id string) (interface{}, pagination.Paginatevalue) {
	p := new(pagination.Paginatevalue)
	key := fmt.Sprintf("%v:%v", user_id, cache.POS_LINK)
	data := cache.GetCacheKey(key)
	key = fmt.Sprintf("%v:%v:%v", user_id, cache.POS_LINK, cache.PAGINATION)
	pagination := cache.GetCacheKey(key)
	helpers.JsonMarshaller(pagination, p)
	return data, *p
}
