package models

type Block struct {
	Number       string         `json:"number"`
	Hash         string         `json:"hash"`
	Nonce        string         `json:"nonce"`
	GasLimit     string         `json:"gasLimit"`
	GasUsed      string         `json:"gasUsed"`
	Transactions []*Transaction `json:"transactions"`
	Next         *Block
	Prev         *Block
}

type ListBlock struct {
	Head *Block
	Tail *Block
	List map[string]*Block
}

func (l *ListBlock) Insert(block *Block) {
	prevTail := l.Tail.Prev

	prevTail.Next = block
	block.Prev = prevTail
	block.Next = l.Tail
	l.Tail.Prev = block

	l.List[block.Number] = block
}

func (l *ListBlock) RemoveHead() {
	next := l.Head.Next

	l.Head.Next = next.Next
	next.Next.Prev = l.Head

	delete(l.List, next.Number)
}

type Transaction struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Hash     string `json:"hash"`
	Nonce    string `json:"nonce"`
}
