package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type CommandOutput struct {
	Command string `json:"command"`
	Output  string `json:"output"`
	Date    string `json:"date"`
}

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test [string to print]",
	Short: "Executa o comando npm test do projeto e envia o output para uma API",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		userCommand := strings.Join(args, " ")
		fullCommand := "npm test " + userCommand
		fmt.Println("🎹 Comando: " + fullCommand)

		// Executa o comando no terminal do estudante
		output, err := exec.Command("npm", "test", userCommand).CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%s\n", output)

		// Monta o payload para enviar para a API
		cmdOutput := CommandOutput{
			Command: fullCommand,
			// Output:  string(output), caso no futuro o retorno do npm seja util
			Date: time.Now().Format(time.RFC3339),
		}
		payload, err := json.Marshal(cmdOutput)
		if err != nil {
			fmt.Println("🔥 Error:", err)
			return
		}
		fmt.Println("📦 Payload:", cmdOutput)

		// Faz a requisição POST para a API
		resp, err := http.Post("http://example.com/api/command-output", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("🔥 Error:", err)
			return
		}
		fmt.Println("📬 Comando Enviado:" + resp.Status)

		defer resp.Body.Close()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
