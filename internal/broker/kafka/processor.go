package kafka

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/fajarabdillahfn/go-wallet/internal/model"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

func AboveThresholdProcessor() {
	log.Println("start aboveThreshold processor...")
	cb := func(ctx goka.Context, msg interface{}) {

		walletId, _ := strconv.Atoi(ctx.Key())

		amount, _ := strconv.ParseFloat(msg.(string), 64)

		_, ok := model.ThresholdList[int32(walletId)]
		if !ok {
			model.ThresholdList[int32(walletId)] = &model.ThresholdCheck{
				Amounts:   amount,
				StartTime: time.Now(),
			}
		} else {
			if time.Now().Minute()-model.ThresholdList[int32(walletId)].StartTime.Minute() < 2 {
				model.ThresholdList[int32(walletId)].Amounts += amount

				if model.ThresholdList[int32(walletId)].Amounts > 10001 {
					model.ThresholdList[int32(walletId)].AboveThreshold = true
				}
			} else {
				model.ThresholdList[int32(walletId)].Amounts = amount
				model.ThresholdList[int32(walletId)].AboveThreshold = false
			}
		}
	}

	g := goka.DefineGroup(aboveThresholdGroup,
		goka.Input(topic, new(codec.String), cb),
		goka.Persist(new(codec.Int64)),
	)

	p, err := goka.NewProcessor(brokers, g)
	if err != nil {
		log.Fatalf("error creating processor: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		if err = p.Run(ctx); err != nil {
			log.Fatalf("error running processor: %v", err)
		}
	}
}

func BalanceProcessor() {
	log.Println("start balance processor...")
	cb := func(ctx goka.Context, msg interface{}) {

		walletId, _ := strconv.Atoi(ctx.Key())

		amount, _ := strconv.ParseFloat(msg.(string), 64)

		_, ok := model.WalletsBalance[int32(walletId)]
		if !ok {
			model.WalletsBalance[int32(walletId)] = amount
		} else {
			model.WalletsBalance[int32(walletId)] += amount
		}
	}

	g := goka.DefineGroup(balanceGroup,
		goka.Input(topic, new(codec.String), cb),
		goka.Persist(new(codec.Int64)),
	)

	p, err := goka.NewProcessor(brokers, g)
	if err != nil {
		log.Fatalf("error creating processor: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		if err = p.Run(ctx); err != nil {
			log.Fatalf("error running processor: %v", err)
		}
	}
}
