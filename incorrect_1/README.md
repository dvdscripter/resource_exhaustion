This code has a type of DoS (denial of service) called Resource Exhaustion.
Resource Exhaustion occurs when developers allocate some type of resource and 
forget to clean up after job is done.

http.Request.ParseMultipartForm(maxMemory int64) will create files at temporary 
directory if maxMemory is exceeded. After running this code with input csv you 
can check your temporary directory is can be clean of any file. Increasing 
maxMemory parameter can help to avoid creating temporary files but your service 
can start consuming more memory than you can provide. Even with increased 
maxMemory clients can still send files exceeding your maxMemory and without 
proper clean up resource exhaustion can occur.