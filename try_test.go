package gtry

import (
	"context"
	"fmt"
	"testing"

	"github.com/midoks/golang-try/test"
)

var (
	ctx = context.TODO()
)

func Test_Try(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := `gutil Try test`
		t.Assert(Try(ctx, func(ctx context.Context) {
			panic(s)
		}), s)
	})
	gtest.C(t, func(t *gtest.T) {
		s := `gutil Try test`
		t.Assert(Try(ctx, func(ctx context.Context) {
			panic(gerror.New(s))
		}), s)
	})
}

func Test_TryCatch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		TryCatch(ctx, func(ctx context.Context) {
			panic("gutil TryCatch test")
		}, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		TryCatch(ctx, func(ctx context.Context) {
			panic("gutil TryCatch test")

		}, func(ctx context.Context, err error) {
			t.Assert(err, "gutil TryCatch test")
		})
	})

	gtest.C(t, func(t *gtest.T) {
		TryCatch(ctx, func(ctx context.Context) {
			panic(gerror.New("gutil TryCatch test"))

		}, func(ctx context.Context, err error) {
			t.Assert(err, "gutil TryCatch test")
		})
	})
}
