package handlers

import (
    "context"
    "gitlab.com/tokend/subgroup/project/internal/data"
    "net/http"

    "gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
    logCtxKey ctxKey = iota
    personCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
    return func(ctx context.Context) context.Context {
        return context.WithValue(ctx, logCtxKey, entry)
    }
}

func Log(r *http.Request) *logan.Entry {
    return r.Context().Value(logCtxKey).(*logan.Entry)
}

func Person(r *http.Request) data.PersonQ {
    return r.Context().Value(personCtxKey).(data.PersonQ).New()
}

func CtxPerson(q data.PersonQ) func(context.Context) context.Context {
    return func(ctx context.Context) context.Context {
        return context.WithValue(ctx, personCtxKey, q)
    }
}