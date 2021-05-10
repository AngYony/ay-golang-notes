package wyfile

type File struct {
	fd int
}

//打开文件
func OpenFile(name string) (f *File, err error) {
	return nil, nil
}

//关闭文件
func (f *File) Close() error {
	return nil
}

//读文件数据
func (f *File) Read(offset int64, data []byte) int {
	return 0
}
