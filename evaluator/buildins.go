package evaluator

import "monkey/object"

var buildins = map[string]*object.Buildin{
	"len": { // TODO: fix this build in type
		Fn: func(args ...object.Object) object.Object {

			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch args := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(args.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
					args.Type())
			}
		},
	},
}
