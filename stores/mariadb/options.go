package mariadb

//Options ...
type Options struct {
	Offset int64
	Limit  int64
	Column string
	Sort   string
}

//Option ...
type Option func(*Options)

//WithOffset ...
func WithOffset(offset int64) Option {
	return func(args *Options) {
		args.Offset = offset
	}
}

//WithLimit ...
func WithLimit(limit int64) Option {
	return func(args *Options) {
		args.Limit = limit
	}
}

//WithColumn ...
func WithColumn(column string) Option {
	return func(args *Options) {
		args.Column = column
	}
}

//WithSort ...
func WithSort(sort string) Option {
	return func(args *Options) {
		args.Sort = sort
	}
}
