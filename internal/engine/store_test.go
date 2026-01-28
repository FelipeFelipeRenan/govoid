package engine

import (
	"fmt"
	"sync"
	"testing"
)

// TestHashDeterminism garante que a matemática não é aleatória
func TestHashDeterminism(t *testing.T) {
	key := "analepsy"
	expected := HashFNV32(key)

	// Roda 100 vezes pra ter certeza absoluta
	for i := 0; i < 100; i++ {
		if result := HashFNV32(key); result != expected {
			t.Fatalf("O Hash mudou! Esperado %d, Recebido %d", expected, result)
		}
	}
}

// TestShardDistribution verifica se estamos respeitando o limite do array
func TestShardIndexBounds(t *testing.T) {
	store := NewStringStore() // Certifique-se que você criou essa função no store.go
	
	inputs := []string{"java", "go", "rust", "cpp", "python", "javascript"}

	for _, key := range inputs {
		shard := store.getShard(key)
		if shard == nil {
			t.Fatalf("Shard nil retornado para a chave %s", key)
		}

	}
}

func TestConcurrentAccess(t *testing.T) {
	store := NewStringStore()
	
	// Vamos usar WaitGroup para esperar todas as goroutines terminarem
	var wg sync.WaitGroup

	// Cenário: 100 goroutines escrevendo e lendo ao mesmo tempo
	// Se o seu Sharding ou Mutex estiverem errados, isso vai explodir (panic ou race).
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			key := fmt.Sprintf("key-%d", id) // Cada um escreve uma chave diferente
			store.Set(key, "value")
			
			// E tenta ler uma chave "quente" que todo mundo acessa
			// Isso testa se o RLock está funcionando (vários lendo ao mesmo tempo)
			_, _ = store.Get("key-0") 
		}(i)
	}

	wg.Wait()
}

func BenchmarkHashFNV32(b *testing.B) {
	key := "uma_chave_longa_para_testar_o_algoritmo_de_hash_do_voidkv"
	
	// b.N é controlado pelo Go. Ele vai rodar 1 vez, depois 100, depois 1 milhão...
	// até conseguir uma média estatística confiável.
	for i := 0; i < b.N; i++ {
		HashFNV32(key)
	}
}