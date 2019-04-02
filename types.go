package vectorizing

import "time"

type TT struct {
	IDTT                       string    `json:"idtt,omitempty"`
	IDTTnonriscontrato         string    `json:"idttnonriscontrato,omitempty"`
	APPOENTITA                 string    `json:"appoentita,omitempty"`
	CODICERISCONTRODIS         string    `json:"codiceriscontrodis,omitempty"`
	DATAORAINIZIOSEGNALAZIONE  time.Time `json:"@timestamp,omitempty"`
	DATAORASEGNALAZIONERECLAMO time.Time `json:"timestampsegnalazionereclamo,omitempty"`
	DATAORAINIZIOINTERVENTO    time.Time `json:"timestampiniziointervento,omitempty"`
	DATAORAFINEINTERVENTO      time.Time `json:"timestampfineintervento,omitempty"`
	DATAORAPRESAINCARICOTT     time.Time `json:"timestamppresaincaricott,omitempty"`
	DATAORACREAZIONETT         time.Time `json:"timestampcreazionett,omitempty"`
	DATAORACHIUSURATT          time.Time `json:"timestampchiusura,omitempty"`
	DESCDIAGNOSI               string    `json:"diagnosi,omitempty"`
	DESCRIZIONEEVENTO          string    `json:"descrizioneevento,omitempty"`
	STATODELTT                 string    `json:"statott,omitempty"`
	SISTEMAEMITTENTE           string    `json:"sistemaemittente,omitempty"`
	TTGRUPPORESP               string    `json:"ttgrupporesp,omitempty"`
	TTGRUPPOCOMP               string    `json:"ttgruppocomp,omitempty"`
	NOMEGRUPPO                 string    `json:"nomegruppo,omitempty"`
	COMPETENTETT               string    `json:"competentett,omitempty"`
	RESPONSABILEDELTT          string    `json:"responsabilett,omitempty"`
	SUBMITTER                  string    `json:"submitter,omitempty"`
	NOTECNA                    string    `json:"notecna,omitempty"`
	NOTEDICHIUSURA             string    `json:"notechiusura,omitempty"`
	NOTELONG                   string    `json:"notelong,omitempty"`
	TT                         string    `json:"tt,omitempty"`
	RECAPITOTELEFONICO1        string    `json:"telefono1,omitempty"`
	RECAPITOTELEFONICO2        string    `json:"telefono2,omitempty"`
	RECAPITOTELEFONICO3        string    `json:"telefono3,omitempty"`
	CELLS                      []string  `json:"cells,omitempty"`
	MAILS                      []string  `json:"mails,omitempty"`
	TOKENS                     []string  `json:"tokens,omitempty"`
}

type Parola struct {
	Lemma      string
	Occorrenze int
}

type Parole []Parola

type Diagnosi uint8

const (
	Other Diagnosi = iota
	CuboVisionContenuti
	SunriseEvoluti
	CuboMusicaContenuti
	TimVisionSmartTVContenuti
	TimVisionSmartTVAcquisti
)
