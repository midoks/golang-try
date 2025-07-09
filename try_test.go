package gtry

import (
	"context"
	"fmt"
	"os"
	// "reflect"
	"testing"
)

var (
	ctx = context.TODO()
)

// T is the testing unit case management object.
type T struct {
	*testing.T
}

// Assert checks `value` and `expect` EQUAL.
// func (t *T) Assert(value, expect interface{}) {
// 	Assert(value, expect)
// }

// Assert checks `value` and `expect` EQUAL.
// func Assert(value, expect interface{}) {
// 	rvExpect := reflect.ValueOf(expect)
// 	if empty.IsNil(value) {
// 		value = nil
// 	}
// 	if rvExpect.Kind() == reflect.Map {
// 		// if err := compareMap(value, expect); err != nil {
// 		// 	panic(err)
// 		// }
// 		return
// 	}
// 	var (
// 		strValue  = string(value)
// 		strExpect = string(expect)
// 	)
// 	if strValue != strExpect {
// 		panic(fmt.Sprintf(`[ASSERT] EXPECT %v == %v`, strValue, strExpect))
// 	}
// }

// C creates a unit testing case.
// The parameter `t` is the pointer to testing.T of stdlib (*testing.T).
// The parameter `f` is the closure function for unit testing case.
func C(t *testing.T, f func(t *T)) {
	defer func() {
		if err := recover(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			t.Fail()
		}
	}()
	f(&T{t})
}

func Test_Try(t *testing.T) {
	// C(t, func(t *T) {
	// 	s := `gutil Try test`
	// 	t.Assert(Try(ctx, func(ctx context.Context) {
	// 		panic(s)
	// 	}), s)
	// })
	// C(t, func(t *gtest.T) {
	// 	s := `gutil Try test`
	// 	t.Assert(Try(ctx, func(ctx context.Context) {
	// 		panic(gerror.New(s))
	// 	}), s)
	// })
}

// func Test_TryCatch(t *testing.T) {
// 	gtest.C(t, func(t *gtest.T) {
// 		TryCatch(ctx, func(ctx context.Context) {
// 			panic("gutil TryCatch test")
// 		}, nil)
// 	})

// 	gtest.C(t, func(t *gtest.T) {
// 		TryCatch(ctx, func(ctx context.Context) {
// 			panic("gutil TryCatch test")

// 		}, func(ctx context.Context, err error) {
// 			t.Assert(err, "gutil TryCatch test")
// 		})
// 	})

// 	gtest.C(t, func(t *gtest.T) {
// 		TryCatch(ctx, func(ctx context.Context) {
// 			panic(gerror.New("gutil TryCatch test"))

// 		}, func(ctx context.Context, err error) {
// 			t.Assert(err, "gutil TryCatch test")
// 		})
// 	})
// }
