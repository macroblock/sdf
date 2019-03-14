package chainer

import (
	"fmt"
	"unicode/utf8"

	"github.com/macroblock/imed/pkg/ptool"
	"github.com/macroblock/imed/pkg/zlog/zlog"
)

var (
	log = zlog.Instance("chainer")
)

var rule = `
entry       = '' { '<' keyExpr '>' | @rune } $;

keyExpr    = {@mod ('+'|'-')} (@key|@rune);

rune        = !'<'!'>!'-'!'+' \x21..\xfe;

key         = 'enter'|'esc'|'f1'|'space';
mod         = 'shift'|'alt'|'ctrl';

    = { ' ' | \x09 | \x0a | \x0d };
`

var parser *ptool.TParser

func init() {
	p, err := ptool.NewBuilder().FromString(rule).Entries("entry").Build()
	if err != nil {
		fmt.Println("\nparser error: ", err)
		panic("")
	}
	parser = p
}

func keychainToBinary(keychain string) ([]BinaryKey, error) {
	tree, err := parser.Parse(keychain)
	if err != nil {
		return nil, err
	}
	data := []BinaryKey{}
	var traverse func(*ptool.TNode) error
	traverse = func(tree *ptool.TNode) error {
		for _, node := range tree.Links {
			nodeType := parser.ByID(node.Type)
			switch nodeType {
			default:
				return fmt.Errorf("unsupported node type %q", nodeType)
			case "groupSet", "keySet":
				err = traverse(node)
			case "rune":
				r, _ := utf8.DecodeRuneInString(node.Value)
				data = append(data, MakeBinaryKey(int(r), 0, 0, 0))
			case "range":
				r1, _ := utf8.DecodeRuneInString(node.Links[0].Value)
				r2, _ := utf8.DecodeRuneInString(node.Links[1].Value)
				if r1 > r2 {
					r1, r2 = r2, r1
				}
				for i := r1; i < r2; i++ {
					data = append(data, MakeBinaryKey(int(i), 0, 0, 0)) // error
				}
			case "mod", "key":
			}
		}
		return nil
	}

	err = traverse(tree)
	if err != nil {
		return nil, err
	}
	return data, nil
}
