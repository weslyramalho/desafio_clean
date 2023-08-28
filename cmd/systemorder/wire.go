//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/weslyramalho/GO/tree/main/clean_arch/internal/entity"
	"github.com/weslyramalho/GO/tree/main/clean_arch/internal/event"
	"github.com/weslyramalho/GO/tree/main/clean_arch/internal/infra/database"
	"github.com/weslyramalho/GO/tree/main/clean_arch/internal/infra/web"
	"github.com/weslyramalho/GO/tree/main/clean_arch/internal/usecase"
	"github.com/weslyramalho/GO/tree/main/clean_arch/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUserCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
