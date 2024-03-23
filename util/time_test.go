package util

import "testing"

func TestStringToDate(t *testing.T) {
	var convertedTime = StringToTime("2024-03-22T23:11:00")

	if convertedTime.Year() != 2024 {
		t.Errorf("Ano esperado: 2024.")
	}

	if convertedTime.Month() != 03 {
		t.Errorf("Mês esperado: 03.")
	}

	if convertedTime.Day() != 22 {
		t.Errorf("Dia esperado: 22.")
	}

	if convertedTime.Hour() != 23 {
		t.Errorf("Hora esperada: 23.")
	}

	if convertedTime.Minute() != 11 {
		t.Errorf("Minuto esperado: 11.")
	}

	if convertedTime.Second() != 0 {
		t.Errorf("Segundo esperado: 0.")
	}
}
