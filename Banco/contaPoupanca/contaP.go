package contaPoupanca

import "contaCorrente/Banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (conta *ContaPoupanca) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= conta.saldo
	if podeSacar {
		conta.saldo -= valorDoSaque
		return "Saque Realizado com Sucesso!\nsaldo atual:", conta.saldo
	} else {
		return "saldo insuficiente\nsaldo atual:", conta.saldo
	}
}

func (c *ContaPoupanca) Depositar(ValorDoDeposito float64) string {
	if ValorDoDeposito > 0 {
		c.saldo += ValorDoDeposito
		return "Deposito realizado com sucesso!"
	} else {
		return "Não foi possível realizar o depósito!"
	}

}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) PagarBoletoContaPoupanca(valorDoBoleto float64) string {
	if valorDoBoleto <= c.saldo {
		c.saldo -= valorDoBoleto
		return "Boleto pago com sucesso!"
	} else {
		return "Saldo insuficiente"
	}
}
