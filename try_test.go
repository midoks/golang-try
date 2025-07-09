package gtry

import (
	"context"
	"fmt"
	"testing"
)

var (
	ctx = context.TODO()
)

func Test_Try(t *testing.T) {
	s := `gutil Try test`
	err := Try(ctx, func(ctx context.Context) {
		panic(s)
	})

	fmt.Println(err)
}

func Test_TryCatch(t *testing.T) {
	TryCatch(ctx, func(ctx context.Context) {
		panic("gutil TryCatch test")
	}, func(ctx context.Context, err error) {
		fmt.Println("Test_TryCatch:", err)
		// t.Assert(err, "gutil TryCatch test")
	})
}
