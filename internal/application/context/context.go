package context

type Data interface {
	SetClient(client string)
}

type PrepareContext struct {
	Tenant string
}

func NewPrepareContext(tenant string) *PrepareContext {
	return &PrepareContext{
		Tenant: tenant,
	}
}

func (ctx *PrepareContext) SetContext(data Data) {
	data.SetClient(ctx.Tenant)
}
