package gtry


import (
	"fmt"
	"context"
)

// HasStack checks and reports whether `err` implemented interface `IStack`.
func HasStack(err error) bool {
	_, ok := err.(IStack)
	return ok
}

// Throw throws out an exception, which can be caught be TryCatch or recover.
func Throw(exception interface{}) {
	panic(exception)
}


func Try(ctx context.Context, try func(ctx context.Context)) (err error) {
	if try == nil {
		return
	}

	defer func() {
        if exception := recover(); exception != nil {
        	if v, ok := exception.(error); ok && HasStack(v){
				err = v
			} else {
				 err = fmt.Errorf("panic: %v", exception) 
			}
        }
    }()
	try(ctx)
	return
}

func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error)) {
	if try == nil {
		return
	}
	if exception := Try(ctx, try); exception != nil && catch != nil {
		catch(ctx, exception)
	}
}