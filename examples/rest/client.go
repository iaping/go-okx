package rest

import rc "github.com/iaping/go-okx/rest"

// 敏感信息申请的模拟盘KEY，不确定何时会删除
var TestClient = rc.New("", "1d1f3552-5355-4205-954e-565102283aaa", "F15537181073B9654F260153B093A8EA", "Abc%123456", true, nil)
