# Usando a biblioteca

**Como utilizar:**

- [Como Usar](#como-usar)
- [Exemplos de uso](#exemplos-de-uso)

## Como usar

Para configuração inicial, a biblioteca irá precisar de 2 informações:
- Service name
- Level

Para alterar a configuração, basta instânciar uma nova configuração com o parâmetros configurados para a função `MakeLogger` do pacote de `factory`.

```Golang
import (
	"github.com/madeiramadeirabr/lib-go-logger/factory"
	f "github.com/madeiramadeirabr/lib-go-logger/formatter"
	"github.com/madeiramadeirabr/lib-go-logger/log_level"
	"github.com/madeiramadeirabr/lib-go-logger/logger"
)

    ...
	config := logger.Config{
		ServiceName: "Teste",
		Level:       log_level.LogLevelTrace,
	}

	logger := factory.MakeLogger(config)

    logger.Info("Mensagem descritiva", f.LogMessageOptions{
		GlobalEventName: "A_SAMPLE_EVENT",
		Context: map[string]string{
			"foo": "bar",
		},
		TraceId:   "abcdefghijkl",
		SessionId: "mnopqrstuv",
	})

	logger.Debug("This is a debug message", f.LogMessageOptions{})

	logger.Info("This is a info message", f.LogMessageOptions{})

	logger.Warning("This is a warning message", f.LogMessageOptions{})

	logger.Error("This is a error message", f.LogMessageOptions{})

	logger.Emergency("This is a emergency message", f.LogMessageOptions{})
```

## Exemplos de uso

Golang
```Golang
	logger.Info("This is a info message", f.LogMessageOptions{
		GlobalEventName: "A_SAMPLE_EVENT",
		Context: map[string]string{
			"foo": "bar",
		},
		TraceId:   "abcdefghijkl",
		SessionId: "mnopqrstuv",
	})

	logger.Error("This is a error message", f.LogMessageOptions{})
```

Resposta

```json
{
  "global_event_timestamp": "2023-08-21T16:50:22-03:00",
  "global_event_name": "A_SAMPLE_EVENT",
  "level": "INFO",
  "context": { "foo": "bar" },
  "message": "This is a info message",
  "service_name": "Teste",
  "trace_id": "abcdefghijkl",
  "session_id": "mnopqrstuv",
}

{
    "global_event_timestamp": "2023-08-21T16:50:22-03:00",
    "level": "ERROR",
    "message": "This is a error message",
    "service_name": "Teste"
}
```