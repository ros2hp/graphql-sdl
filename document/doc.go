package document

import (
	"github.com/ros2hp/graphql-sdl/internal/db"
)

var NoItemFoundErr = db.NoItemFoundErr

func GetDocument() string {
	return db.GetDocument()
}

func SetDocument(doc string) {
	db.SetDocument(doc)
}

func DeleteType(obj string) error {
	return db.DeleteType(obj)
}

func SetDefaultDoc(doc string) {
	db.SetDefaultDoc(doc)
}
