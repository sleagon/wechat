// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package custom

import (
	"errors"
	"net/http"

	"github.com/sleagon/wechat/mp"
)

type Client struct {
	mp.WechatClient
}

// 创建一个新的 Client.
//  如果 HttpClient == nil 则默认用 http.DefaultClient
func NewClient(TokenServer mp.TokenServer, HttpClient *http.Client) *Client {
	if TokenServer == nil {
		panic("TokenServer == nil")
	}
	if HttpClient == nil {
		HttpClient = http.DefaultClient
	}

	return &Client{
		WechatClient: mp.WechatClient{
			TokenServer: TokenServer,
			HttpClient:  HttpClient,
		},
	}
}

// 发送客服消息, 文本.
func (clt *Client) SendText(msg *Text) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return clt.send(msg)
}

// 发送客服消息, 图片.
func (clt *Client) SendImage(msg *Image) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return clt.send(msg)
}

// 发送客服消息, 语音.
func (clt *Client) SendVoice(msg *Voice) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return clt.send(msg)
}

// 发送客服消息, 视频.
func (clt *Client) SendVideo(msg *Video) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return clt.send(msg)
}

// 发送客服消息, 音乐.
func (clt *Client) SendMusic(msg *Music) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return clt.send(msg)
}

// 发送客服消息, 图文.
func (clt *Client) SendNews(msg *News) (err error) {
	if msg == nil {
		return errors.New("msg == nil")
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return clt.send(msg)
}

func (clt *Client) send(msg interface{}) (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="
	if err = clt.PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
