package structs

type ClienteNome string
type ClienteEmail string
type ClienteCPF string

type Cliente struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

func (cliente *Cliente) SetNome(nome string) {
	cliente.Nome = nome
}

type ClienteInternacional struct {
	Cliente   Cliente `json:"cliente"`
	Documento string  `json:"documento"`
}

func (c *ClienteInternacional) SetDocumento(documento string) {
	c.Documento = documento
}

func DeclareCliente() Cliente {
	cliente := Cliente{
		Nome:  "Fábio",
		Email: "fah_ds@live.com",
		CPF:   "123.456.789-00",
	}
	return cliente
}

func OtherDeclareCliente() Cliente {
	cliente := Cliente{
		"Fábio",
		"fah_ds@live.com",
		"123.456.789-00",
	}
	return cliente
}

func DeclareClienteInternacional() ClienteInternacional {
	cliente := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Fábio",
			Email: "fah_ds@live.com",
			CPF:   "123.456.789-00",
		},
		Documento: "12.25-XX",
	}
	cliente.Cliente.SetNome("Fábio dos Santos")
	cliente.SetDocumento("12.25-yy")
	return cliente
}
