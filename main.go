package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Structure for vulnerability patterns
type Vulnerability struct {
	Search   []string
	Severity string
	Message  string
}

// Structure for findings
type Finding struct {
	Vulnerability string
	Severity      string
	Message       string
	Line          int
	Content       string
}

// Create sample config file
func createSampleConfig() string {
	config := `DB_PASSWORD=admin123
DB_USER=root
API_KEY=sl-1234567890abcdef
DEBUG_MODE=True
SSL_ENABLED=False
ADMIN_PASSWORD=password
FIREWALL_DISABLED=yes`

	file, err := os.Create("server_config.txt")
	if err != nil {
		fmt.Println("[-] Error creating file:", err)
		return ""
	}
	defer file.Close()

	file.WriteString(config)
	fmt.Println("[+] Created server_config.txt")
	return "server_config.txt"
}

// Define patterns
func getPatterns() map[string]Vulnerability {
	return map[string]Vulnerability{
		"weak_password": {
			Search:   []string{"password=admin", "password=password"},
			Severity: "HIGH",
			Message:  "Secret found in config file - should use environment variables",
		},
		"exposed_secret": {
			Search:   []string{"API_KEY=", "SECRET_TOKEN="},
			Severity: "CRITICAL",
			Message:  "Secret found in config file - should use environment variables",
		},
		"insecure_setting": {
			Search:   []string{"DEBUG_MODE=True", "SSL_ENABLED=False", "FIREWALL_DISABLED=yes"},
			Severity: "HIGH",
			Message:  "Insecure setting detected - security feature disabled",
		},
		"default_credentials": {
			Search:   []string{"DB_USER=root", "DB_PASSWORD=admin123"},
			Severity: "HIGH",
			Message:  "Default credentials found - change immediately",
		},
	}
}

// Scan file
func scanFile(filename string, patterns map[string]Vulnerability) []Finding {
	var findings []Finding

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("[-] File does not exist:", filename)
		return findings
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		lineUpper := strings.ToUpper(strings.TrimSpace(line))

		for vulnName, vulnInfo := range patterns {
			for _, search := range vulnInfo.Search {
				if strings.Contains(lineUpper, strings.ToUpper(search)) {
					findings = append(findings, Finding{
						Vulnerability: vulnName,
						Severity:      vulnInfo.Severity,
						Message:       vulnInfo.Message,
						Line:          lineNum,
						Content:       line,
					})
					break
				}
			}
		}
	}

	return findings
}

// Print report
func printReport(findings []Finding) {
	fmt.Println("\n============================================================")
	fmt.Println("VULNERABILITY SCAN REPORT")
	fmt.Println("============================================================")

	if len(findings) == 0 {
		fmt.Println("\n[+] No vulnerabilities found!\n")
		return
	}

	critical := 0
	high := 0

	for _, f := range findings {
		if f.Severity == "CRITICAL" {
			critical++
		}
		if f.Severity == "HIGH" {
			high++
		}
	}

	fmt.Printf("\nTotal: %d | CRITICAL: %d | HIGH: %d\n\n", len(findings), critical, high)

	for _, severity := range []string{"CRITICAL", "HIGH"} {
		fmt.Printf("[!] %s ISSUES:\n", severity)
		idx := 1
		for _, f := range findings {
			if f.Severity == severity {
				fmt.Printf("  %d. Line %d (%s)\n", idx, f.Line, f.Content)
				fmt.Printf("      %s\n", f.Message)
				idx++
			}
		}
	}

	fmt.Println("============================================================\n")
}

// Main function
func main() {
	fmt.Println("============================================================")
	fmt.Println("AUTO VULNERABILITY SCANNER")
	fmt.Println("============================================================")

	filename := createSampleConfig()
	patterns := getPatterns()

	fmt.Printf("\n[+] Loaded %d patterns\n", len(patterns))

	findings := scanFile(filename, patterns)
	printReport(findings)

	fmt.Println("Scan Complete!")
}