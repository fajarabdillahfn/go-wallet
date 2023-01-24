package kafka

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/fajarabdillahfn/go-wallet/internal/model"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

func Emit(ctx context.Context, data *model.Deposit) error {
	_, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	emitter, err := goka.NewEmitter(brokers, topic, new(codec.String))
	if err != nil {
		return err
	}
	defer emitter.Finish()

	wallet_id := strconv.Itoa(int(data.Id))

	err = emitter.EmitSync(wallet_id, fmt.Sprintf("%f", data.Amount))
	if err != nil {
		return err
	}

	return nil
}
