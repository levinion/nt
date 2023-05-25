package nt

type Ctx struct {
	foo      map[string]any
	err      error
	readOnly bool
}

// 为上下文设置变量，若存在则直接覆盖
func (c *Ctx) Set(key string, value any) {
	if c.readOnly {
		return
	}
	c.foo[key] = value
}

// 安全地设置变量，若变量已存在则不覆盖，返回false
func (c *Ctx) SafeSet(key string, value any) bool {
	if c.readOnly {
		return false
	}
	if _, ok := c.foo[key]; !ok {
		c.foo[key] = value
		return true
	}
	return false
}

// 从上下文中取值，若键不存在则返回nil
func (c *Ctx) Get(key string) any {
	return c.foo[key]
}

// 发生错误，终止剩余流程并返回错误
func (c *Ctx) Error(err error) {
	c.err = err
}

func (c *Ctx) add(cx *Ctx) {
	for k, v := range cx.foo {
		c.Set(k, v)
	}
}

func (c *Ctx) toggleStatus() {
	c.readOnly = !c.readOnly
}

type Template struct {
	ctx *Ctx
	f   []func(c *Ctx)
}

// 创建模板，若指定名称则可利用Find方法全局获取模板
func Create(name ...string) *Template {
	switch len(name) {
	case 0:
		return &Template{
			ctx: &Ctx{
				foo:      map[string]any{},
				err:      nil,
				readOnly: false,
			},
			f: make([]func(c *Ctx), 0),
		}
	case 1:
		t := &Template{
			ctx: &Ctx{
				foo:      map[string]any{},
				err:      nil,
				readOnly: false,
			},
			f: make([]func(c *Ctx), 0),
		}
		ntPool.pool[name[0]] = t
		return t
	default:
		return nil
	}
}

// 注册或追加模板函数
func (t *Template) Join(f func(c *Ctx)) {
	t.f = append(t.f, f)
}

// 调用模板执行函数；注意在此处使用的Set方法无效
func (t *Template) Call(f func(c *Ctx)) error {
	for i := range t.f {
		t.f[i](t.ctx)
		if t.ctx.err != nil {
			return t.ctx.err
		}
	}
	//开启和解除写锁
	t.ctx.toggleStatus()
	defer t.ctx.toggleStatus()
	f(t.ctx)
	if t.ctx.err != nil {
		return t.ctx.err
	}
	return nil
}

// 向上下文中添加变量
func (t *Template) Watch(flag string, foo any) {
	t.ctx.Set(flag, foo)
}

func (t *Template) WatchMany(foos map[string]any) {
	for k, v := range foos {
		t.ctx.Set(k, v)
	}
}

// 合并两个模板（包括上下文和函数）
func (t *Template) Concat(tx *Template) {
	t.ctx.add(tx.ctx)
	t.f = append(t.f, tx.f...)
}

type templatePool struct {
	pool map[string]*Template
}

var ntPool = &templatePool{
	pool: map[string]*Template{},
}

// 在全局查找模板，未找到返回nil
func Find(name string) *Template {
	return ntPool.pool[name]
}
