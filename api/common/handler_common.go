/*
 * Copyright (c) 2023 OceanBase
 * OBAgent is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

package common

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/oceanbase/obagent/lib/http"
	"github.com/oceanbase/obagent/log"
)

// keys stored in gin.Context
const (
	OcpAgentResponseKey = "ocpAgentResponse"
	TraceIdKey          = "traceId"
	OcpServerIpKey      = "ocpServerIp"
)

func NewContextWithTraceId(c *gin.Context) context.Context {
	traceId := ""
	if t, ok := c.Get(TraceIdKey); ok {
		if ts, ok := t.(string); ok {
			traceId = ts
		}
	}
	return context.WithValue(context.Background(), log.TraceIdKey{}, traceId)
}

func SendResponse(c *gin.Context, data interface{}, err error) {
	resp := http.BuildResponse(data, err)
	c.Set(OcpAgentResponseKey, resp)
}
