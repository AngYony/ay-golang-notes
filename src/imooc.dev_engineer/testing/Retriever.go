package testing

type Retriever struct {

}

func (Retriever) Get(url string) string{
	return "这是一条测试信息"
}
