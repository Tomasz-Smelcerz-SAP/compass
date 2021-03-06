package graphql

import (
	"io"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kyma-incubator/compass/components/director/pkg/scalar"
)

type Timestamp time.Time

func (y *Timestamp) UnmarshalGQL(v interface{}) error {
	tmpStr, err := scalar.ConvertToString(v)
	if err != nil {
		return err
	}

	t, err := time.Parse(time.RFC3339, tmpStr)
	if err != nil {
		return err
	}

	*y = Timestamp(t)

	return nil
}

func (y Timestamp) MarshalGQL(w io.Writer) {
	_, err := w.Write([]byte(strconv.Quote(time.Time(y).Format(time.RFC3339))))
	if err != nil {
		log.Errorf("while writing %T: %s", y, err)
	}
}

// UnmarshalJSON is used only by integration tests
func (y *Timestamp) UnmarshalJSON(in []byte) error {
	instr, err := strconv.Unquote(string(in))
	if err != nil {
		return err
	}

	t, err := time.Parse(time.RFC3339, instr)
	if err != nil {
		return err
	}

	*y = Timestamp(t)

	return nil
}
