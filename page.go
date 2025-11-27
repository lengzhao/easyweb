package easyweb

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lengzhao/easyweb/util"
)

var _ Page = &easyPage{}

func newPage() *easyPage {
	var page easyPage
	page.callback = make(map[string]eventMsgData)
	return &page
}

func (p *easyPage) Title(title string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "title", title}
	p.updatePageData(msg)
	return p
}

func (p *easyPage) AddJs(js string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "js", js}
	p.updatePageData(msg)
	return p
}

func (p *easyPage) RunJs(js string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "eval", js}
	p.updatePageData(msg)
	return p
}

func (p *easyPage) GetPeer() string {
	return ""
}

func (p *easyPage) AddCss(css string) Page {
	id := util.GetCallerID(util.LevelParent)
	msg := toClientMsgData{id, "css", css}
	p.updatePageData(msg)
	return p
}

// e 将作为子元素放到div内，如<div id=rand_id>e</div>
func (p *easyPage) Write(e any) string {
	var id string
	gid, ok := e.(IContainerID)
	if ok {
		id = gid.ContainerID()
	}
	if id == "" {
		id = util.GetCallerID(util.LevelParent)
	}
	return p.WriteWithID(id, e)
}

// e 将作为子元素放到id所属的元素内，如<div id=id>e</div>
func (p *easyPage) WriteWithID(id string, e any) string {
	if id == "" {
		id = util.GetCallerID(util.LevelParent)
	}

	msg := toClientMsgData{id, "", fmt.Sprint(e)}
	p.updatePageData(msg)

	if e, ok := e.(IAfterLoaded); ok {
		e.AfterLoaded(p)
	}
	return id
}

func (p *easyPage) Replace(e IGetID) string {
	id := e.GetID()
	if id == "" {
		return ""
	}
	msg := toClientMsgData{id, "replace", fmt.Sprint(e)}
	p.updatePageData(msg)
	return id
}

func (p *easyPage) Delete(id string) {
	msg := toClientMsgData{id, "", ""}
	p.updatePageData(msg)
}

type attrInfo struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// 修改指定id的元素的属性
func (p *easyPage) SetAttr(id, key, value string) string {
	info := attrInfo{Key: key, Value: value}
	data, _ := json.Marshal(info)
	msg := toClientMsgData{id, "attr", string(data)}
	p.updatePageData(msg)
	return id
}

// 注册指定id的前端事件，typ是前端事件名，cb是回调还是，如果cb为空，则是取消事件回调
func (p *easyPage) RegistEvent(id, typ string, cb IMessageCb) {
	if id == "" || typ == "" {
		return
	}
	log.Println("page.RegistEvent", id, typ, cb)
	p.callback[id] = eventMsgData{ID: id, Event: cb}
}

// 更新页面数据
func (p *easyPage) updatePageData(msg any) {
	p.mu.Lock()
	defer p.mu.Unlock()

	log.Println("updatePageData", msg)
	// 根据消息类型更新页面数据
	switch m := msg.(type) {
	case toClientMsgData:
		// 处理客户端消息数据
		// 将更新操作添加到pageData数组中
		p.elements = append(p.elements, m)
	case eventMsgData:
		p.callback[m.ID] = m
	}
}

func (p *easyPage) MessageCallbackFromFramwork(session Session, id string, dataType CbDataType, data []byte) bool {
	// 查找对应的事件处理器
	p.mu.Lock()
	cb, exists := p.callback[id]
	p.mu.Unlock()

	log.Println("Page.MessageCallbackFromFramwork", id, exists, dataType, string(data))
	// 如果找到了事件处理器且不为nil，调用它
	if exists && cb.Event != nil {
		return cb.Event.MessageCallbackFromFramwork(session, id, dataType, data)
	}

	// 没有找到匹配的事件处理器
	return false
}

func (p *easyPage) PageLoaded(session Session) {
	for _, msg := range p.elements {
		session.Write(msg)
	}
}
