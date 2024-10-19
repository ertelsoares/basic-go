package main

import (
	"context"
	"fmt"
	"grpc/client"
	"grpc/server"
	"html/template"
	"net/http"
	"sync"

	"google.golang.org/grpc"
	"grpc/pb"
)

// Função para servir o HTML
func loginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("login.html"))
	tmpl.Execute(w, nil)
}

// Função para processar o login via POST
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Pegando os dados do formulário
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Conectando ao servidor gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	// Fazendo a chamada para o servidor gRPC
	req := &pb.AddUserRequest{
		Id:       "1",
		Username: username,
		Password: password,
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		http.Error(w, "Login failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Exibindo o resultado
	response := fmt.Sprintf(`{"message": "Login successful!", "user_id": "%s", "username": "%s"}`, res.Id, res.Username)
	fmt.Fprintf(w, response)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2) // Espera duas goroutines

	go func() {
		defer wg.Done() // Marca esta goroutine como concluída
		server.Run()
	}()

	go func() {
		defer wg.Done() // Marca esta goroutine como concluída
		client.Run()
	}()

	wg.Wait()
}
