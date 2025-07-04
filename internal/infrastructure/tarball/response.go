package tarball

type SaveResponse struct {
	Id string
}

type GetResponse struct {
	Tarball []byte
}