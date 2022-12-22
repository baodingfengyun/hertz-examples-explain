/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// 修改配置
func main() {
	// The default listening port is 8888.
	// You can modify it with server.WithHostPorts().
	// 传递不定数的参数
	h := server.Default(
		// 设置地址和端口
		server.WithHostPorts("127.0.0.1:8080"),
		// 设置最大请求体大小
		server.WithMaxRequestBodySize(20<<20),
		// 设置传输层:标准库
		server.WithTransport(standard.NewTransporter),
	)

	// 设置路由
	h.GET("/hello", helloHandler)

	h.Spin()
}

func helloHandler(ctx context.Context, c *app.RequestContext) {
	c.String(consts.StatusOK, "Hello hertz!")
}
