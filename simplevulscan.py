import os

def create_sample_config():
         config = """DB_PASSWORD=admin123
    DB_USER=root
    API_KEY=sl-1234567890abcdef
    DEBUG_MODE=True
    SSL_ENABLED=False
    ADMIN_PASSWORD=password
    FIREWALL_DISABLED=yes"""
         with open("server_config.txt", "w") as f:
            f.write(config)
         print("[+] Created server_config.txt")
         return "server_config.txt"

def get_patterns():
    return {
         "weak_password": {
              "search": ["password=admin", "password=password"],
              "severity": "High",
              "message": "Secret found in config file - should use enevironment variables"
         },
         "exposed_secret": {
                "search": ["API_KEY=", "SECRET_TOKEN="],
                "severity": "CRITICAL",
                "message": "Secret found in config file - should use environment variables"
             },"insecure_setting": {
                "search": ["DEBUG_MODE=True", "SSL_ENABLED=False", "FIREWALL_DISABLED=yes"],
                "severity": "HIGH",
                "message": "Insecure setting detected - security feature disabled"

             },
             "default_credentials": {
                "search": ["DB_USER=root", "DB_PASSWORD=admin123"],
                "severity": "HIGH",
                "message": "Default credentials found - change immediately"
             }
              
         }


def scan_file(filename,patterns):
     findings = []
     if not os.path.exists(filename):
            print(f"[-] File {filename} does not exist")
            return findings
     with open(filename, "r") as f:
           lines = f.readlines()

     for line_num, line in enumerate(lines, start=1):
        line_upper = line.upper().strip()
        for vuln_name, vuln_info in patterns.items():
                for search in vuln_info["search"]:
                    if search.upper() in line_upper:
                        findings.append({
                            "vulnerability": vuln_name,
                            "severity": vuln_info["severity"],
                            "message": vuln_info["message"],
                            "line": line_num,
                            "content": line.strip()
                     })
                        break
                return findings 
                
                
def print_report(findings):
                             print("\n" + "=" * 60)
                             print("VULNERABILITY SCAN REPORT")
                             print("=" * 60)
                             if not findings:
                                    print("\n[+] No vulnerabilities found!\n")
                                    return 
                             critical = sum(1 for f in findings if f["severity"] == "CRITICAL")
                             high = sum(1 for f in findings if f["severity"] == "HIGH")

                             print(f"\ntotal: {len(findings)} | CRITICAL: {critical} | HIGH: {high}\n")
                             
                             for severity in ["CRITICAL", "HIGH"]:
                                    severity_findings = [f for f in findings if f["severity"] == severity]
                                    if findings["severity"] == severity:
                                           print(f"[!] {severity} ISSUES:")
                                    for idx, f in enumerate (severity_findings, 1):
                                                    print(f"  {idx}. Line{f['line']} (line {f['content']})")
                                                    print(f"      {f['message']}")
print("=" * 60 + "\n")
def main():
       print("=" * 60)
print("AUTO VULNERABILITY SCANNER")
print("=" * 60)

filename = create_sample_config()
patterns = get_patterns()
print(f"\n[+] Loaded {len(patterns)} patterns")
findings = scan_file(filename, patterns)
print_report(findings)
print("Scan Complete!\n")
if __name__ == "__main__":
    main()




                                           