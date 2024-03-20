package ecomorph

import (
	"testing"
	"time"
)

//func TestGetEcomorphById(t *testing.T) {
//	StartServerGRPC()
//
//	panic("impl")
//	//conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
//	//if err != nil {
//	//	log.Error("Could not connect: %v", err)
//	//}
//	//defer conn.Close()
//	//
//	//ctx := context.Background()
//	//
//	//client := ecomorph.NewEcomorphServiceClient("localhost:8080")
//
//}

func TestFirst(t *testing.T) {
	// Канал для синхронизации завершения теста
	done := make(chan struct{})

	// Горутина для выполнения первого теста
	go func() {
		// Ваш код первого теста
		time.Sleep(3 * time.Second) // Предположим, что это ваш длительный тест
		t.Log("First test done")
		close(done) // Оповещаем о завершении первого теста
	}()

	// Дожидаемся завершения первого теста
	<-done
}

func TestSecond(t *testing.T) {
	// Канал для синхронизации завершения первого теста
	done := make(chan struct{})

	// Горутина для выполнения второго теста
	go func() {
		// Дожидаемся завершения первого теста
		<-done
		// Ваш код второго теста
		time.Sleep(2 * time.Second) // Предположим, что это ваш длительный тест
		t.Log("Second test done")
	}()

	// Необходимая работа для второго теста
	// ...

	// Завершаем второй тест и даем время для выполнения
	done <- struct{}{}
	time.Sleep(4 * time.Second)
}
