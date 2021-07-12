HashiCorp Vault

  Vault is a tool which is for securely accessing secrets. Secrets is anything that you want a tightly controlled access to or one which you want to secure such as API, keys, passwords, or certificates.

  Write

  The write command is used to write the data in vault at the specific path. The data can be credentials, secrets, configuration or arbitrary data. 
       
       go run vault-write.go
       
       Output: 
       2021/07/12 13:12:28 Success: Data stored in secret/hello 
       
  Now try to read the data from vault. To verify, whether the data is been stored into vault or not.
  
      vault read secret/hello
      
      Output:
      Key                 Value
      ---                 -----
      refresh_interval    768h
      name                Priyanshi
      password            admin1234
      username            priya12
  
  Read

  The read  command reads data from Vault at the given path. It will reads the secrets and generate dynamic credentials and get configuration details.
  
      go run vault-read.go
      
      Output:
      secret &{a9b3bc66-d3e1-7fc0-0082-660de84f189f  2764800 false map[name:Priyanshi password:admin1234 username:priya12] [] <nil> <nil>}
  
      
