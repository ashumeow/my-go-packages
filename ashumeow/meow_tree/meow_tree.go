package meow_tree

type meowTree interface {
	meowInsert(interface{})
	meowExists(interface{}) bool
	meowEmpty(interface{}) bool
	meowSearch(interface{}) interface{}
}