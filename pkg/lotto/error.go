package lotto

type LottoError string

func (e LottoError) Error() string { return string(e) }
