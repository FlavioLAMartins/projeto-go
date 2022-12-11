package entity

type Product struct { //NOME DE VARIAVEL TEM Q SER MAISUCULA POR CAUSA DO JSON
	ID        uint32 `json:"id"`
	Name      string `json:"name_prod"`
	Price     string `json:"price_prod"`
	Code      string `json:"code_prod"`
	CreatedAt string `json:"created_at,omitempty"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAdmin() *User {
	return &User{
		Username: "admin",
		Password: "supersenha",
	}
}

type Token struct {
	Token string `json:"token"`
}

const USER_TOKEN = "fake-WzD5fqrlaAXLv26bpI0hxvAhDp7T1Bac"
