package models

type User struct {
	Id       string `bson:"_id,omitempty"`
	Username string `bson:"username,omitempty" required:"true"`
	Email    string `bson:"email,omitempty" unique:"true" required:"true"`
	Roll     string `bson:"roll,omitempty" default:"user"`
	Password string `bson:"password"`
	Created  string `bson:"created,omitempty" default:"now"`
}

type Category struct {
	Id       string `json:"id" bson:"_id,omitempty"`
	Category string `json:"category" bson:"category"`
	Created  string `bson:"created,omitempty" default:"now"`
}

type Note struct {
	Id       string `bson:"_id,omitempty"`
	Title    string `bson:"title" required:"true"`
	Image    string `bson:"image" required:"true"`
	Content  string `bson:"content" required:"true"`
	Category string `bson:"category" required:"true"`
	Writer   string `bson:"writer" required:"true"`
	Created  string `bson:"created" default:"now"`
}