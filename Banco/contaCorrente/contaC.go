package contacorrente

import "contaCorrente/Banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float32
}

func (conta *ContaCorrente) Sacar(valorDoSaque float32) (string, float32) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= conta.saldo
	if podeSacar {
		conta.saldo -= valorDoSaque
		return "Saque Realizado com Sucesso!\nsaldo atual:", conta.saldo
	} else {
		return "saldo insuficiente\nsaldo atual:", conta.saldo
	}
}

func (c *ContaCorrente) Depositar(ValorDoDeposito float32) (string, float32) {
	if ValorDoDeposito > 0 {
		c.saldo += ValorDoDeposito
		return "Deposito realizado com sucesso!\nsaldo atual:", c.saldo
	} else {
		return "Não foi possível realizar o depósito!\nsaldo atual:", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valorDaTrasferencia float32, contaDeDestino *ContaCorrente) bool {
	if valorDaTrasferencia <= c.saldo {
		c.saldo -= valorDaTrasferencia
		contaDeDestino.Depositar(valorDaTrasferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float32 {
	return c.saldo
}

func (c *ContaCorrente) PagarBoletoContaPoupanca(valorDoBoleto float32) bool {
	if valorDoBoleto <= c.saldo {
		c.saldo -= valorDoBoleto
		return true
	} else {
		return false
	}
}
