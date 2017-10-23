// Copyright 2017 Lunny Xiao. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dingtalk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	webhook := NewWebhook("<your_access_token>")
	err := webhook.SendTextMsg("这是一个没有AT的文本消息", false)
	assert.NoError(t, err)

	err = webhook.SendTextMsg("这是一个AT所有人的文本消息", true)
	assert.NoError(t, err)

	err = webhook.SendLinkMsg("这是一个Link消息", "此内容链接到百度首页", "", "http://www.baidu.com")
	assert.NoError(t, err)

	err = webhook.SendMarkdownMsg("这是一个markdown消息不AT", `# 新闻简介
1. [第一条](http://baidu.com)
2. [第二条](http://baidu.com)
3. [第三条](http://baidu.com)`, false)
	assert.NoError(t, err)

	err = webhook.SendMarkdownMsg("这是一个markdown消息@所有人", `# 新闻简介
		* [第一条](http://baidu.com)
		* [第二条](http://baidu.com)
		* [第三条](http://baidu.com)`, true)
	assert.NoError(t, err)

	err = webhook.SendSingleActionCardMsg("这是一个single actionCard 消息", "链接到百度", "百度首页", "http://www.baidu.com", false, false)
	assert.NoError(t, err)

	err = webhook.SendActionCardMsg("这是一个actionCard消息", "链接到百度",
		[]string{"第一个链接", "第二个链接"},
		[]string{"http://baidu.com", "http://baidu.com"},
		false,
		false,
	)
	assert.NoError(t, err)

	err = webhook.SendLinkCardMsg([]LinkMsg{
		{
			Title:      "第一个消息",
			MessageURL: "http://baidu.com",
			PicURL:     "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png",
		},
		{
			Title:      "第二个消息",
			MessageURL: "http://baidu.com",
			PicURL:     "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png",
		},
		{
			Title:      "第三个消息",
			MessageURL: "http://baidu.com",
			PicURL:     "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png",
		},
	})
	assert.NoError(t, err)
}
