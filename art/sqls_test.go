package art

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_ParseSql(t *testing.T) {

	//file := "../demo/sql/tree/test.sql"
	//file := "../demo/sql/tree/tree.sql"
	file := "../demo/sql/init/2.data.sql"

	bytes, err := ioutil.ReadFile(file)
	panicIfErr(err)

	sqls := ParseSqls(pref, &FileEntity{file, string(bytes)})

	fmt.Println("segs------")
	for _, x := range sqls {
		fmt.Printf("%#v\n", x)
	}
}

func Test_DepairQuote(t *testing.T) {

	q2 := "`'12345'`"
	fmt.Printf("%s\n", q2)

	cnt := countQuotePair(q2)
	fmt.Printf("%d\n", cnt)
}